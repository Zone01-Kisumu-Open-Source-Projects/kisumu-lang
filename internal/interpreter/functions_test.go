package interpreter

import (
	"testing"
)

func TestFunctionLiterals(t *testing.T) {
	input := "const fn = func(x) { x + 2; }; fn"

	evaluated := testEval(input)
	fn, ok := evaluated.(*Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0].String())
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"const identity = func(x) { x; }; identity(5);", 5},
		{"const double = func(x) { x * 2; }; double(5);", 10},
		{"const add = func(x, y) { x + y; }; add(5, 5);", 10},
		{"const add = func(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"const fn = func(x) { x; }; fn(5);", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
	const newAdder = func(x) {
		return func(y) { x + y; };
	};
	const addTwo = newAdder(2);
	addTwo(2);
	`

	testIntegerObject(t, testEval(input), 4)
}

func TestRecursiveFunctions(t *testing.T) {
	input := `
	const countDown = func(x) {
		if (x > 0) {
			return countDown(x - 1);
		} else {
			return x;
		}
	};
	countDown(1);
	`

	testIntegerObject(t, testEval(input), 0)
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		if (10 > 1) {
			if (10 > 1) {
				return 10;
			}
			return 1;
		}
		`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionWithReturn(t *testing.T) {
	input := `
	const add = func(x, y) {
		return x + y;
	};
	add(5, 5);
	`

	testIntegerObject(t, testEval(input), 10)
}

func TestFunctionWithoutReturn(t *testing.T) {
	input := `
	const add = func(x, y) {
		x + y;
	};
	add(5, 5);
	`

	testIntegerObject(t, testEval(input), 10)
}

func TestFunctionParameters(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"func() { 5; }();", 5},
		{"func(x) { x; }(5);", 5},
		{"func(x, y) { x + y; }(5, 5);", 10},
		{"func(x, y, z) { x + y + z; }(5, 5, 5);", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionScope(t *testing.T) {
	input := `
	const x = 10;
	const add = func(y) {
		x + y;
	};
	add(5);
	`

	testIntegerObject(t, testEval(input), 15)
}

func TestFunctionErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5(5);",
			"not a function: INTEGER",
		},
		{
			"const add = func(x, y) { x + y; }; add(5);",
			"wrong number of arguments. got=1, want=2",
		},
		{
			"const add = func(x, y) { x + y; }; add(5, 5, 5);",
			"wrong number of arguments. got=3, want=2",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		err, ok := evaluated.(*Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if err.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, err.Message)
		}
	}
}

func TestNestedFunctions(t *testing.T) {
	input := `
	const outer = func(x) {
		func(y) {
			x + y;
		};
	};
	const inner = outer(10);
	inner(5);
	`

	testIntegerObject(t, testEval(input), 15)
}

func TestFunctionAsValue(t *testing.T) {
	input := `
	const add = func(x, y) { x + y; };
	const sub = func(x, y) { x - y; };
	const apply = func(fn, x, y) { fn(x, y); };
	apply(add, 10, 5);
	`

	testIntegerObject(t, testEval(input), 15)
}
