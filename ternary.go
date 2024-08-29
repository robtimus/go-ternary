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
// Unlike Else the result is evaluated lazily.
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
// Unlike Return the result is evaluated lazily.
func Call[T any](fn func() T) TrueResult[T] {
	return TrueResult[T]{fn}
}
