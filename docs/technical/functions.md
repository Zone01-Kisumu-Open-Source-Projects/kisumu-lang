# Kisumu Programming Language Specification

## 7. Functions

Functions in Kisumu are reusable blocks of code designed to perform specific tasks. They allow for modular, organized, and maintainable code.

### 7.1 Defining Functions

A function is defined using the `func` keyword, followed by the function name, parameters (if any), and the function body enclosed in curly braces `{}`.

#### Syntax

```k
func functionName(parameters) returnType {
    // Function body
    return value; // (Optional) Return statement
}
```

#### Example

```k
func greet(name string) string {
    return "Hello, " + name;
}
```

### 7.2 Invoking Functions

A function is called by using its name followed by parentheses `()`. If the function requires parameters, their values are passed inside the parentheses.

#### Syntax

```k
functionName(arguments);
```

#### Example

```k
string message = greet("Kisumu");
print(message); // Output: Hello, Kisumu
```

### 7.3 Parameters and Arguments

#### 7.3.1 Positional Parameters

Parameters are declared in the function definition and are passed in the same order during function invocation.

##### Example

```k
func add(a int, b int) int {
    return a + b;
}

int result = add(5, 10); // Output: 15
```

#### 7.3.2 Default Parameters

Kisumu supports default parameter values. If a value is not provided during invocation, the default is used.

##### Syntax

```k
func display(message string = "Hello, World!") {
    print(message);
}

display(); // Output: Hello, World!
```

#### 7.3.3 Variadic Functions

Functions can accept a variable number of arguments using the `...` syntax.

##### Example

```k
func sum(values ...int) int {
    int total = 0;
    for (int value : values) {
        total += value;
    }
    return total;
}

int total = sum(1, 2, 3, 4); // Output: 10
```

### 7.4 Return Values

Functions can return a single value or multiple values. Use the `return` keyword to specify the returned value(s).

#### 7.4.1 Single Return Value

```k
func square(num int) int {
    return num * num;
}
```

#### 7.4.2 Multiple Return Values

```k
func divide(a int, b int) (int, int) {
    return a / b, a % b; // Quotient and remainder
}

int quotient, remainder = divide(10, 3);
print(quotient, remainder); // Output: 3, 1
```

### 7.5 Scope

Variables declared within a function are local to that function and cannot be accessed outside its body.

#### Example

```k
func demo() {
    int localVar = 42;
    print(localVar); // Accessible here
}

// print(localVar); // Error: localVar is not defined
```

### 7.6 Anonymous Functions and Closures

Anonymous functions (functions without a name) and closures (functions capturing variables from their enclosing scope) are supported in Kisumu.

#### Anonymous Functions

```k
var square = func(x int) int {
    return x * x;
};

print(square(5)); // Output: 25
```

#### Closures

```k
func makeCounter() func() int {
    int count = 0;

    return func() int {
        count++;
        return count;
    };
}

var counter = makeCounter();
print(counter()); // Output: 1
print(counter()); // Output: 2
```

### 7.7 Recursion

A function can call itself to solve problems that can be broken into smaller sub-problems.

#### Example

```k
func factorial(n int) int {
    if (n == 0) {
        return 1;
    }
    return n * factorial(n - 1);
}

int result = factorial(5); // Output: 120
```

### 7.8 Error Handling in Functions

Kisumu functions will support structured error handling (e.g., `try-catch-finally`) in a future update.
