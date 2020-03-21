package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("The stack is already empty \n")
	}

	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}

func pushValues(stack *Stack) {

	fmt.Printf("pushing the value camila to the stack \n")
	stack.Push("camila")

	fmt.Printf("pushing the value daniel to the stack \n")
	stack.Push("daniel")

	fmt.Printf("pushing the value 33 to the stack \n")
	stack.Push(33)

	fmt.Printf("pushing the value 3.1 to the stack \n")
	stack.Push(3.1)

}

func popValues(stack *Stack) {

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println("Poping value ", value)
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Is stack empty ? ", stack.IsEmpty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func main() {

	stack := Stack{}
	fmt.Println("Initializing your stack")
	fmt.Println("Size? ", stack.Size())

	fmt.Println("Pushing values into the stack")
	pushValues(&stack)
	fmt.Printf("The size of the stack is %d \n", stack.Size())
	fmt.Println("Poping values out of the stack")
	popValues(&stack)
	fmt.Printf("The size of the stack is %d \n", stack.Size())
}
