# Kisumu Language: Math Package

- This is the documentation for the Math Package in the Kisumu language.

## Contents

- [Kisumu Language: Math Package](#kisumu-language-math-package)
  - [Contents](#contents)
  - [Features](#features)
    - [Constants](#constants)
    - [Basic Math Operations](#basic-math-operations)
    - [Power and Root functions](#power-and-root-functions)
    - [Trigonometric Functions](#trigonometric-functions)
    - [Hyperbolic functions](#hyperbolic-functions)
    - [Rounding and Precision](#rounding-and-precision)
    - [Minimum, maximum and absolute values](#minimum-maximum-and-absolute-values)
    - [Random Number Generation](#random-number-generation)
    - [Factorial Combinatorics](#factorial-combinatorics)
    - [Number Properties](#number-properties)
    - [Degrees and Radian Conversion](#degrees-and-radian-conversion)
  - [About](#about)

## Features

In the math package, several functions and implementations are provided, including absolute value functions, precision functions, trigonometric functions, and basic floating-point operations.

### Constants

- First, we will start with the constant.
A constant variable is a buffer storage that does not allow manipulation of its initial value; it is commonly called an immutable variable.
- Some fundamental mathematical constants that will be featured are;

  - [X] The Pi constant: $\pi \approx 3.1415926535 $
  - [X] Euler's constant: $e \approx 2.7182818284 $
  - [X] Tau... : $\tau \approx 6.2831853072 $
  - [X] Phi... : $\Phi \approx 1.6180339887 $
  - [X] Square root of 2: $\sqrt{2} \approx 1.4142135623 $

### Basic Math Operations

The math package will handle some basic math operations, such as addition, subtraction, division, and multiplication.
- Examples;

  - add(a, b) → number
  - sub(a, b) → number
  - mul(a, b) → number
  - div(a, b) → number (Handle division by zero)
  - mod(a, b) → number.
- The math package for the Kisumu language will feature some basic math operations.

### Power and Root functions

- Some of the exponential functions that the Ksm language will feature in the coming days include:
  - _Power function_ pow(base, exponent) → number
  - _Square of a number_ sqrt(x) → number
  - _Cube root of a number_ cbrt(x) → number (Cube root)
  - _Exponent function_ exp(x) → number $e^x$
  - _Logarithm base x_ log(x) → number (Natural logarithm)
  - _Logarithm base 10_ log10(x) → number
  - _Logarithm base 2_ log2(x) → number

### Trigonometric Functions

- Some functions that will feature in the trigonometric package include:
  - sin(x) → number $\sin{x}$
  - cos(x) → number $\cos{x}$
  - tan(x) → number $\tan{x}$
  - asin(x) → number $\sinh{x}$
  - acos(x) → number $\cos-{x}$
  - atan(x) → number $\tan-{x}$
  - atan2(y, x) → number $\tan-{2(x, y)}$

### Hyperbolic functions

- Some of the functions under the trigonometric functions are hyperbolic functions:
  - sinh(x) → number $\sinh{x}$
  - cosh(x) → number $\cosh{x}$
  - tanh(x) → number $\tanh{x}$
  - asinh(x) → number $\sinh-{x}$
  - acosh(x) → number $\cosh-{x}$
  - atanh(x) → number $\tanh-{x}$

### Rounding and Precision

- In this section, the following will be laid out:
  - ceil(x) → number (Rounds up) $\lceil x \rceil$
  - floor(x) → number (Rounds down) $\lfloor x \rfloor$
  - round(x) → number (Nearest integer) $\R(x)$
  - trunc(x) → number (Removes decimal part)

### Minimum, maximum and absolute values

- min(a, b) → number
- max(a, b) → number
- abs(x) → number $|x|$

### Random Number Generation

- rand() → number (Returns a random float between 0 and 1)
- randInt(min, max) → int (Random integer in range)
- randFloat(min, max) → number
- seed(n) → void (Seeds the RNG)

### Factorial Combinatorics

- factorial(n) → number
- perm(n, r) → number (Permutations) $P(n)$
- comb(n, r) → number (Combinations)

### Number Properties

- isPrime(n) → bool
- gcd(a, b) → number (Greatest Common Divisor)
- lcm(a, b) → number (Least Common Multiple)

### Degrees and Radian Conversion

- toRadians(deg) → number $x\degree$ → $x^c$
- toDegrees(rad) → number $x^c$ → $x\degree$

## About

- The functions above were suggested by [Anxiel Raymond](https://martinraymondogwel.netlify.app)
- More features and ideas are welcome.
