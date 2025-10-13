package interpreter

import (
	"fmt"
	"strings"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/ast"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	StringObj      = "STRING"
	ErrorObj       = "ERROR"
	ReturnValueObj = "RETURN_VALUE"
	FunctionObj    = "FUNCTION"
	BreakObj       = "BREAK"
	ContinueObj    = "CONTINUE"
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

func (i *Integer) Type() ObjectType { return IntegerObj }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BooleanObj }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value
type Null struct{}

func (n *Null) Type() ObjectType { return NullObj }
func (n *Null) Inspect() string  { return "null" }

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return StringObj }
func (s *String) Inspect() string  { return s.Value }

// Error represents an error value
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ErrorObj }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// ReturnValue represents a return value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }
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
	case *ast.IfExpression:
		return i.evalIfExpression(node)
	case *ast.WhileLoop:
		return i.evalWhileLoop(node)
	case *ast.ForLoop:
		return i.evalForLoop(node)
	case *ast.BreakStatement:
		return BREAK
	case *ast.ContinueStatement:
		return CONTINUE
	case *ast.BlockStatement:
		return i.evalBlockStatement(node)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &Function{Parameters: params, Body: body, Env: i.env}
	case *ast.CallExpression:
		function := i.Eval(node.Function)
		if isError(function) {
			return function
		}
		args := i.evalExpressions(node.Arguments)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return i.applyFunction(function, args)
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
	if right.Type() != IntegerObj {
		return &Error{Message: "unknown operator: -" + string(right.Type())}
	}

	value := right.(*Integer).Value
	return &Integer{Value: -value}
}

func (i *Interpreter) evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type() == IntegerObj && right.Type() == IntegerObj:
		return i.evalIntegerInfixExpression(operator, left, right)
	case left.Type() == StringObj && right.Type() == StringObj:
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
	case left.Type() == IntegerObj && right.Type() == IntegerObj:
		leftVal := left.(*Integer).Value
		rightVal := right.(*Integer).Value
		return &Boolean{Value: leftVal == rightVal}
	case left.Type() == BooleanObj && right.Type() == BooleanObj:
		leftVal := left.(*Boolean).Value
		rightVal := right.(*Boolean).Value
		return &Boolean{Value: leftVal == rightVal}
	case left.Type() == StringObj && right.Type() == StringObj:
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
	if equality.Type() == ErrorObj {
		return equality
	}
	return &Boolean{Value: !equality.(*Boolean).Value}
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ErrorObj
	}
	return false
}

// Function represents a function object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FunctionObj }
func (f *Function) Inspect() string {
	var out string
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out += "func(" + strings.Join(params, ", ") + ") " + f.Body.String()
	return out
}

// Break represents a break object
type Break struct{}

func (b *Break) Type() ObjectType { return BreakObj }
func (b *Break) Inspect() string  { return "break" }

// Continue represents a continue object
type Continue struct{}

func (c *Continue) Type() ObjectType { return ContinueObj }
func (c *Continue) Inspect() string  { return "continue" }

// Predefined values
var (
	TRUE     = &Boolean{Value: true}
	FALSE    = &Boolean{Value: false}
	NULL     = &Null{}
	BREAK    = &Break{}
	CONTINUE = &Continue{}
)

// evalIfExpression evaluates an if expression
func (i *Interpreter) evalIfExpression(ie *ast.IfExpression) Object {
	condition := i.Eval(ie.Condition)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return i.Eval(ie.Consequence)
	}
	if ie.Alternative != nil {
		return i.Eval(ie.Alternative)
	}
	return NULL
}

// evalWhileLoop evaluates a while loop
func (i *Interpreter) evalWhileLoop(wl *ast.WhileLoop) Object {
	var result Object = NULL

	for {
		condition := i.Eval(wl.Condition)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		result = i.Eval(wl.Body)
		if isError(result) {
			return result
		}

		// Handle break and continue
		if result.Type() == BreakObj {
			break
		}
		if result.Type() == ContinueObj {
			continue
		}
	}

	return result
}

// evalForLoop evaluates a for loop
func (i *Interpreter) evalForLoop(fl *ast.ForLoop) Object {
	var result Object = NULL

	// Execute initializer
	if fl.Initializer != nil {
		initResult := i.Eval(fl.Initializer)
		if isError(initResult) {
			return initResult
		}
	}

	for {
		// Check condition
		if fl.Condition != nil {
			condition := i.Eval(fl.Condition)
			if isError(condition) {
				return condition
			}

			if !isTruthy(condition) {
				break
			}
		}

		// Execute body
		result = i.Eval(fl.Body)
		if isError(result) {
			return result
		}

		// Handle break and continue
		if result.Type() == BreakObj {
			break
		}
		if result.Type() == ContinueObj {
			// Execute update before continuing
			if fl.Update != nil {
				updateResult := i.Eval(fl.Update)
				if isError(updateResult) {
					return updateResult
				}
			}
			continue
		}

		// Execute update
		if fl.Update != nil {
			updateResult := i.Eval(fl.Update)
			if isError(updateResult) {
				return updateResult
			}
		}
	}

	return result
}

// evalBlockStatement evaluates a block statement
func (i *Interpreter) evalBlockStatement(block *ast.BlockStatement) Object {
	var result Object = NULL

	// Create new environment for block scope
	env := NewEnclosedEnvironment(i.env)
	oldEnv := i.env
	i.env = env
	defer func() { i.env = oldEnv }()

	for _, statement := range block.Statements {
		result = i.Eval(statement)
		if isError(result) {
			return result
		}

		// Handle return, break, continue
		if result.Type() == ReturnValueObj || result.Type() == BreakObj || result.Type() == ContinueObj {
			return result
		}
	}

	return result
}

// evalExpressions evaluates a list of expressions
func (i *Interpreter) evalExpressions(exps []ast.Expression) []Object {
	result := []Object{}

	for _, e := range exps {
		evaluated := i.Eval(e)
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

// applyFunction applies a function to arguments
func (i *Interpreter) applyFunction(fn Object, args []Object) Object {
	function, ok := fn.(*Function)
	if !ok {
		return &Error{Message: "not a function: " + string(fn.Type())}
	}

	if len(function.Parameters) != len(args) {
		return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=%d", len(args), len(function.Parameters))}
	}

	// Create new environment for function scope
	env := NewEnclosedEnvironment(function.Env)

	// Bind parameters to arguments
	for paramIdx, param := range function.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	// Save old environment and set function environment
	oldEnv := i.env
	i.env = env
	defer func() { i.env = oldEnv }()

	// Evaluate function body
	result := i.Eval(function.Body)
	if isError(result) {
		return result
	}

	// Handle return value
	if result.Type() == ReturnValueObj {
		return result.(*ReturnValue).Value
	}

	return result
}

// isTruthy determines if an object is truthy
func isTruthy(obj Object) bool {
	switch obj := obj.(type) {
	case *Null:
		return false
	case *Boolean:
		return obj.Value
	default:
		return true
	}
}
