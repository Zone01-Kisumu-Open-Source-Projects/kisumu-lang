# Kisumu Lang — Syntax Reference

## Comments

```ksm
// Single-line comment

/* Multi-line
   comment */
```

---

## Variables

```ksm
const x = 10;
const name = "Kisumu";
const isReady = true;
```

---

## Data Types

```ksm
const integer = 42;
const float   = 3.14;
const str     = "hello";
const bool    = true;
```

---

## Arithmetic

```ksm
5 + 3   // 8
5 - 3   // 2
5 * 3   // 15
6 / 2   // 3
```

---

## Comparisons

```ksm
5 == 5  // true
5 != 3  // true
5 > 3   // true
3 < 5   // true
```

---

## String Concatenation

```ksm
"Hello, " + "Kisumu!"  // Hello, Kisumu!
```

---

## If / Else

```ksm
if (x > 5) {
    return true;
} else {
    return false;
}
```

---

## Functions

```ksm
const add = func(a, b) {
    return a + b;
};
add(5, 3)  // 8
```

---

## Recursion

```ksm
const factorial = func(n) {
    if (n == 0) { return 1; }
    return n * factorial(n - 1);
};
factorial(5)  // 120
```

---

## While Loop

```ksm
while (condition) {
    // body
}
```

---

## For Loop

```ksm
for (const i = 0; i < 5; i = i + 1) {
    // body
}
```

---

## Break / Continue

```ksm
break;
continue;
```

## Running scripts 
```
go run ./cmd/repl/main.go -file nameofyourfile.ks
```