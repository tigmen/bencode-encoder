package main

import "fmt"

const (
	STATE_D = iota
	STATE_L
	STATE_I
)

func main() {
	test := []byte("d1:ci20ee")
	state_stack := make([]int, 0, len(test)/2)

	for _, value := range test {
		switch value {
		case 'd':
			state_stack = append(state_stack, STATE_D)

		case 'i':
			state_stack = append(state_stack, STATE_L)

		case 'e':
			state_stack = state_stack[:len(state_stack)-1]

		default:
		}
		fmt.Println(string(value), state_stack)
	}
}
