package main

import (
	"fmt"
	"math/rand"
	"time"
)

type raceResult struct {
	finishTime time.Time
	car        string
}

func main() {
	fmt.Println("I am in the main method")
	prepareRace()
	time.Sleep(12 * time.Second)
	fmt.Println("Race finished")
}

func prepareRace() {
	finishTime, cancel := startTheRace()
	// https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open
	// No need to close the finishTime channel since it is shared channel
	// Garbage collection will automatically reclaim it when it is no longer used and needed.
	// defer close(finishTime)
	r := <-finishTime
	cancel()

	fmt.Printf("Winner is %s", r.car)
	fmt.Println()
}

func startTheRace() (chan raceResult, func()) {
	finishTime := make(chan raceResult)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}

	fmt.Println("Race started")
	go car1(finishTime, done)
	go car2(finishTime, done)
	go car3(finishTime, done)
	return finishTime, cancel
}

func race(car string) raceResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(10)
	time.Sleep(time.Second * time.Duration(randomNumber))
	fmt.Printf("%s took %d seconds to complete the race", car, randomNumber)
	fmt.Println()
	return raceResult{time.Now(), car}
}

func car1(finishTime chan<- raceResult, done <-chan struct{}) {
	select {
	case finishTime <- race("Car 1"):
		fmt.Println("Car 1 has finished the race")
	case <-done:
		fmt.Println("Car 1 lost, aborting the race.")
		return
	}
}

func car2(finishTime chan<- raceResult, done <-chan struct{}) {
	select {
	case finishTime <- race("Car 2"):
		fmt.Println("Car 2 has finished the race")
	case <-done:
		fmt.Println("Car 2 lost, aborting the race.")
		return
	}
}

func car3(finishTime chan<- raceResult, done <-chan struct{}) {
	select {
	case finishTime <- race("Car 3"):
		fmt.Println("Car 3 has finished the race")
	case <-done:
		fmt.Println("Car 3 lost, aborting the race.")
		return
	}
}
