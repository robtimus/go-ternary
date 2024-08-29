# go-ternary
[![Build Status](https://github.com/robtimus/go-ternary/actions/workflows/build.yml/badge.svg)](https://github.com/robtimus/go-ternary/actions/workflows/build.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=robtimus%3Ago-ternary&metric=alert_status)](https://sonarcloud.io/summary/overall?id=robtimus%3Ago-ternary)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=robtimus%3Ago-ternary&metric=coverage)](https://sonarcloud.io/summary/overall?id=robtimus%3Ago-ternary)
[![Known Vulnerabilities](https://snyk.io/test/github/robtimus/go-ternary/badge.svg)](https://snyk.io/test/github/robtimus/go-ternary)

A simple implementation of ternary expressions in Go.

The Go language does not support ternary expressions. That means that you have to write quite some boilerplate code to achieve the same:

```go
var result TYPE
if condition {
    result = trueResult
} else {
    result = falseResult
}
```

This module allows you to do the same with just a single line. To allow Go to infer the generic type, ternary expressions need to be written as in Python: `trueResult if condition else falseResult`:

```go
result := Return(trueResult).When(condition).Else(falseResult)
```

## Lazy evaluation

The `Return` and `Else` above both require the values to be evaluated eagerly. For constants, pre-existing variables and simple expressions this is fine. However, for more complex expressions it makes more sense to use lazy evaluation. That can be achieved using `Call` and `ElseCall`:

```go
result := Call(func() TYPE { ... }).When(condition).ElseCall(func() TYPE { ... })
```

It's of course also possible to mix eager and lazy evaluation:

```go
result1 := Return(trueResult).When(condition).ElseCall(func() TYPE { ... })
result2 := Call(func() TYPE { ... }).When(condition).Else(falseResult)
```
