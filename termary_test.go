package ternary

import (
	"fmt"
	"testing"
)

func TestTernaryWithEagerEvaluation(t *testing.T) {
	parameters := []struct {
		condition   bool
		trueResult  int
		falseResult int
		expected    int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for i := range parameters {
		condition := parameters[i].condition
		trueResult := parameters[i].trueResult
		falseResult := parameters[i].falseResult
		expected := parameters[i].expected

		t.Run(fmt.Sprintf("%d if %t else %d", trueResult, condition, falseResult), func(t *testing.T) {
			actual := Return(trueResult).When(condition).Else(falseResult)
			if actual != expected {
				t.Errorf("expected: %d, actual: %d", expected, actual)
			}
		})
	}
}

func TestTernaryWithLazyEvaluation(t *testing.T) {
	parameters := []struct {
		condition   bool
		trueResult  int
		falseResult int
		expected    int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for i := range parameters {
		condition := parameters[i].condition
		trueResult := parameters[i].trueResult
		falseResult := parameters[i].falseResult
		expected := parameters[i].expected

		t.Run(fmt.Sprintf("() => %d if %t else () => %d", trueResult, condition, falseResult), func(t *testing.T) {
			trueFunc := &function[int]{result: trueResult}
			falseFunc := &function[int]{result: falseResult}

			actual := Call(trueFunc.Invoke).When(condition).ElseCall(falseFunc.Invoke)
			if actual != expected {
				t.Errorf("expected: %d, actual: %d", expected, actual)
			}
			if condition {
				if trueFunc.invoked != 1 {
					t.Errorf("expected true func to be invoked once, actual: %d", trueFunc.invoked)
				}
				if falseFunc.invoked != 0 {
					t.Errorf("expected false func to not be invoked, actual: %d", falseFunc.invoked)
				}
			} else {
				if trueFunc.invoked != 0 {
					t.Errorf("expected true func to not be invoked, actual: %d", trueFunc.invoked)
				}
				if falseFunc.invoked != 1 {
					t.Errorf("expected false func to be invoked once, actual: %d", falseFunc.invoked)
				}
			}
		})
	}
}

func TestTernaryWithLazyTrueEvaluation(t *testing.T) {
	parameters := []struct {
		condition   bool
		trueResult  int
		falseResult int
		expected    int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for i := range parameters {
		condition := parameters[i].condition
		trueResult := parameters[i].trueResult
		falseResult := parameters[i].falseResult
		expected := parameters[i].expected

		t.Run(fmt.Sprintf("() => %d if %t else %d", trueResult, condition, falseResult), func(t *testing.T) {
			trueFunc := &function[int]{result: trueResult}

			actual := Call(trueFunc.Invoke).When(condition).Else(falseResult)
			if actual != expected {
				t.Errorf("expected: %d, actual: %d", expected, actual)
			}
			if condition && trueFunc.invoked != 1 {
				t.Errorf("expected true func to be invoked once, actual: %d", trueFunc.invoked)
			}
			if !condition && trueFunc.invoked != 0 {
				t.Errorf("expected true func to not be invoked, actual: %d", trueFunc.invoked)
			}
		})
	}
}

func TestTernaryWithLazyFalseEvaluation(t *testing.T) {
	parameters := []struct {
		condition   bool
		trueResult  int
		falseResult int
		expected    int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for i := range parameters {
		condition := parameters[i].condition
		trueResult := parameters[i].trueResult
		falseResult := parameters[i].falseResult
		expected := parameters[i].expected

		t.Run(fmt.Sprintf("%d if %t else () => %d", trueResult, condition, falseResult), func(t *testing.T) {
			falseFunc := &function[int]{result: falseResult}

			actual := Return(trueResult).When(condition).ElseCall(falseFunc.Invoke)
			if actual != expected {
				t.Errorf("expected: %d, actual: %d", expected, actual)
			}
			if condition && falseFunc.invoked != 0 {
				t.Errorf("expected false func to not be invoked, actual: %d", falseFunc.invoked)
			}
			if !condition && falseFunc.invoked != 1 {
				t.Errorf("expected false func to be invoked once, actual: %d", falseFunc.invoked)
			}
		})
	}
}

type function[T any] struct {
	result  T
	invoked int
}

func (f *function[T]) Invoke() T {
	f.invoked++
	return f.result
}
