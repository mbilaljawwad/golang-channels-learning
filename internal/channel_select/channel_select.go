package channel_select

import "fmt"

func ExecuteExample() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}
