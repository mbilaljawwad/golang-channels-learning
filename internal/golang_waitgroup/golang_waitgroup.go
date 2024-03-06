package golang_waitgroup

import (
	"fmt"
	"sync"
)

func ExecuteExample() {
	evilNinjas := []string{"Tommy", "Bobby", "Johnny"}
	var beeper sync.WaitGroup
	for _, evilNinja := range evilNinjas {
		beeper.Add(1)
		go attack(evilNinja, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission completed")
}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja:", evilNinja)
	defer beeper.Done()
}
