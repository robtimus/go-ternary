// The Go language does not support ternary expressions. That means that you have to write quite some boilerplate code to achieve the same:
//
//	var result TYPE
//	if condition {
//		result = trueResult
//	} else {
//		result = falseResult
//	}
//
// This package allows you to do the same with just a single line.
// To allow Go to infer the generic type, ternary expressions need to be written as in Python: "trueResult if condition else falseResult":
//
//	result := ternary.Return(trueResult).When(condition).Else(falseResult)
//
// The above requires the values to be eagerly evaluated. [Call] and [Condition.ElseCall] can be used instead to support lazy evaluation:
//
//	result := ternary.Return(trueResult).When(condition).ElseCall(func() TYPE { ... })
package ternary

// TrueResult represents the result of a ternary expression if the condition is true.
type TrueResult[T any] struct {
	result func() T
}

// When specifies the condition of a ternary expression.
func (r TrueResult[T]) When(condition bool) Condition[T] {
	return Condition[T]{condition, r.result}
}

// Condition represents the condition of a ternary expression.
type Condition[T any] struct {
	condition  bool
	trueResult func() T
}

// Else specifies the result of a ternary expression if the condition is false.
func (c Condition[T]) Else(value T) T {
	if c.condition {
		return c.trueResult()
	}
	return value
}

// ElseCall specifies the result of a ternary expression if the condition is false.
// Unlike [Condition.Else] the result is evaluated lazily.
func (c Condition[T]) ElseCall(fn func() T) T {
	if c.condition {
		return c.trueResult()
	}
	return fn()
}

// Return starts a ternary expression. It comes in the same format as Python, as "true if condition else false", to allow the result type to be inferred.
func Return[T any](value T) TrueResult[T] {
	return TrueResult[T]{func() T { return value }}
}

// Call starts a ternary expression. It comes in the same format as Python, as "true if condition else false", to allow the result type to be inferred.
// Unlike [Return] the result is evaluated lazily.
func Call[T any](fn func() T) TrueResult[T] {
	return TrueResult[T]{fn}
}
