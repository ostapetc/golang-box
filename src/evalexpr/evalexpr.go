package evalexpr

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
)

func EvalExpression(str string) int32 {
	operandStack  := stack.New()
	operatorStack := stack.New()

	for _, char := range str {
		schar := fmt.Sprintf("%c", char)

		if schar == "(" {
			continue
		}

		if schar == "+" || schar == "*" {
			operatorStack.Push(schar)
			continue
		}

		if schar == ")" {
			operandStack.Push(calculate(operandStack, operatorStack))
			continue
		}

		operand, err := strconv.Atoi(schar)
		if err != nil {
			panic(err)
		}

		operandStack.Push(int32(operand))
	}

	fmt.Println("Stack Length", operandStack.Len())

	print(*operandStack, *operatorStack)
	fmt.Println("top1", operatorStack.Pop())
	fmt.Println("top2", operatorStack.Pop())
	//operandStack.Push(calculate(operandStack, operatorStack))
	//print(*operandStack, *operatorStack)

	return operandStack.Pop().(int32)
}

func calculate(operandStack *stack.Stack, operatorStack *stack.Stack) int32 {
	operand1 := operandStack.Pop().(int32)
	operand2 := operandStack.Pop().(int32)
	operator := operatorStack.Pop()

	if operator == "+" {
		return operand1 + operand2
	} else if operator == "*" {
		return operand1 * operand2
	} else {
		fmt.Println("unknown operator ", operator)
		return 0
	}
}

func print(operandStack stack.Stack, operatorStack stack.Stack) {
	fmt.Println("Operand stack")

	for operandStack.Len() > 0 {
		fmt.Println(operandStack.Pop())
	}

	fmt.Println("Operator stack")

	for operatorStack.Len() > 0 {
		fmt.Println(operatorStack.Pop())
	}
}