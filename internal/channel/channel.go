package channel

import (
	"fmt"
	"time"
)

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second * 1)
	fmt.Println("Throwing ninja stars at:", target)
	attacked <- true
}

func ExecuteExample() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)

	fmt.Println("Smoke Signal:", <-smokeSignal)
}
