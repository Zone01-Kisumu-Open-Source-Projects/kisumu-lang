package interpreter

import (
	"fmt"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/ast"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	STRING_OBJ       = "STRING"
	ERROR_OBJ        = "ERROR"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

// Object represents a value in the language
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represents an integer value
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Error represents an error value
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// ReturnValue represents a return value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Environment represents the variable environment
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment with an outer scope
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a value from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set stores a value in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// Interpreter represents the interpreter
type Interpreter struct {
	env *Environment
}

// New creates a new interpreter
func New() *Interpreter {
	return &Interpreter{
		env: NewEnvironment(),
	}
}

// Eval evaluates an AST node
func (i *Interpreter) Eval(node ast.Node) Object {
	switch node := node.(type) {
	case *ast.Program:
		return i.evalProgram(node)
	case *ast.ExpressionStatement:
		return i.Eval(node.Expression)
	case *ast.IntNode:
		return &Integer{Value: int64(node.Value)}
	case *ast.StringNode:
		return &String{Value: node.Value}
	case *ast.BoolNode:
		if node.Value {
			return TRUE
		}
		return FALSE
	case *ast.Identifier:
		return i.evalIdentifier(node)
	case *ast.PrefixExpression:
		right := i.Eval(node.Right)
		if isError(right) {
			return right
		}
		return i.evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := i.Eval(node.Left)
		if isError(left) {
			return left
		}
		right := i.Eval(node.Right)
		if isError(right) {
			return right
		}
		return i.evalInfixExpression(node.Operator, left, right)
	case *ast.LetStatement:
		val := i.Eval(node.Value)
		if isError(val) {
			return val
		}
		i.env.Set(node.Name.Value, val)
		return val
	case *ast.ReturnStatement:
		val := i.Eval(node.ReturnValue)
		if isError(val) {
			return val
		}
		return &ReturnValue{Value: val}
	}
	return &Error{Message: fmt.Sprintf("unknown node type: %T", node)}
}

func (i *Interpreter) evalProgram(program *ast.Program) Object {
	var result Object

	for _, statement := range program.Statements {
		result = i.Eval(statement)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

func (i *Interpreter) evalIdentifier(node *ast.Identifier) Object {
	val, ok := i.env.Get(node.Value)
	if !ok {
		return &Error{Message: "identifier not found: " + node.Value}
	}
	return val
}

func (i *Interpreter) evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return i.evalBangOperatorExpression(right)
	case "-":
		return i.evalMinusPrefixOperatorExpression(right)
	default:
		return &Error{Message: "unknown operator: " + operator + string(right.Type())}
	}
}

func (i *Interpreter) evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func (i *Interpreter) evalMinusPrefixOperatorExpression(right Object) Object {
	if right.Type() != INTEGER_OBJ {
		return &Error{Message: "unknown operator: -" + string(right.Type())}
	}

	value := right.(*Integer).Value
	return &Integer{Value: -value}
}

func (i *Interpreter) evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return i.evalIntegerInfixExpression(operator, left, right)
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		return i.evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return i.evalEqualityExpression(left, right)
	case operator == "!=":
		return i.evalNotEqualExpression(left, right)
	case left.Type() != right.Type():
		return &Error{Message: "type mismatch: " + string(left.Type()) + " " + operator + " " + string(right.Type())}
	default:
		return &Error{Message: "unknown operator: " + string(left.Type()) + " " + operator + " " + string(right.Type())}
	}
}

func (i *Interpreter) evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return &Boolean{Value: leftVal < rightVal}
	case ">":
		return &Boolean{Value: leftVal > rightVal}
	case "==":
		return &Boolean{Value: leftVal == rightVal}
	case "!=":
		return &Boolean{Value: leftVal != rightVal}
	default:
		return &Error{Message: "unknown operator: " + string(left.Type()) + " " + operator + " " + string(right.Type())}
	}
}

func (i *Interpreter) evalStringInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return &Boolean{Value: leftVal == rightVal}
	case "!=":
		return &Boolean{Value: leftVal != rightVal}
	default:
		return &Error{Message: "unknown operator: " + string(left.Type()) + " " + operator + " " + string(right.Type())}
	}
}

func (i *Interpreter) evalEqualityExpression(left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		leftVal := left.(*Integer).Value
		rightVal := right.(*Integer).Value
		return &Boolean{Value: leftVal == rightVal}
	case left.Type() == BOOLEAN_OBJ && right.Type() == BOOLEAN_OBJ:
		leftVal := left.(*Boolean).Value
		rightVal := right.(*Boolean).Value
		return &Boolean{Value: leftVal == rightVal}
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		leftVal := left.(*String).Value
		rightVal := right.(*String).Value
		return &Boolean{Value: leftVal == rightVal}
	case left.Type() != right.Type():
		return FALSE
	default:
		// For predefined values (TRUE, FALSE, NULL), use pointer comparison
		if left == TRUE || left == FALSE || left == NULL {
			return &Boolean{Value: left == right}
		}
		return FALSE
	}
}

func (i *Interpreter) evalNotEqualExpression(left, right Object) Object {
	equality := i.evalEqualityExpression(left, right)
	if equality.Type() == ERROR_OBJ {
		return equality
	}
	return &Boolean{Value: !equality.(*Boolean).Value}
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}

// Predefined values
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	NULL  = &Null{}
)
