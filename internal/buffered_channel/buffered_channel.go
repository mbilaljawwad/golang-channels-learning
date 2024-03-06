package buffered_channel

import (
	"fmt"
	"math/rand"
	"time"
)

func ExecuteExample() {
	channel := make(chan string)
	go throwingNinjaStar(channel)

	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored:", score)
	}
	close(channel)
}
