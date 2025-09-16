package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5 // Number of philosophers and forks

type Fork struct {
	id      int
	request chan chan bool // channel to receive requests for the fork
	release chan bool      // channel to receive release signals
}

func forkProcess(fork *Fork) {
	for {
		// Wait for a philosopher to request the fork
		reply := <-fork.request
		// Grant the fork
		reply <- true
		// Wait for the fork to be released
		<-fork.release
	}
}

type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork
	eatCount  int
	maxEats   int
	wg        *sync.WaitGroup
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
}

func (p *Philosopher) eat() {
	fmt.Printf("Philosopher %d is eating (%d/%d)\n", p.id, p.eatCount+1, p.maxEats)
	time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
}

func (p *Philosopher) dine() {
	defer p.wg.Done()
	for p.eatCount < p.maxEats {
		p.think()

		// To avoid deadlock, the last philosopher picks up right fork first
		var firstFork, secondFork *Fork
		if p.id == N-1 {
			firstFork, secondFork = p.rightFork, p.leftFork
		} else {
			firstFork, secondFork = p.leftFork, p.rightFork
		}

		// Request first fork
		firstReply := make(chan bool)
		firstFork.request <- firstReply
		<-firstReply

		// Request second fork
		secondReply := make(chan bool)
		secondFork.request <- secondReply
		<-secondReply

		// Eat
		p.eat()
		p.eatCount++

		// Release forks
		firstFork.release <- true
		secondFork.release <- true
	}
	fmt.Printf("Philosopher %d is done eating\n", p.id)
}

func main() {
	var wg sync.WaitGroup

	// Create forks
	forks := make([]*Fork, N)
	for i := 0; i < N; i++ {
		forks[i] = &Fork{
			id:      i,
			request: make(chan chan bool),
			release: make(chan bool),
		}
		go forkProcess(forks[i])
	}

	// Create philosophers
	philosophers := make([]*Philosopher, N)
	for i := 0; i < N; i++ {
		philosophers[i] = &Philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%N],
			maxEats:   3,
			wg:        &wg,
		}
	}

	// Start philosophers
	for i := 0; i < N; i++ {
		wg.Add(1)
		go philosophers[i].dine()
	}

	wg.Wait()
	fmt.Println("All philosophers are done eating.")
}

/*
Why this does not deadlock:
- If all philosophers tried to pick up their left fork first, a deadlock could occur if all pick up their left fork and wait for the right.
- Here, the last philosopher (id == N-1) picks up the right fork first, breaking the cycle and ensuring at least one philosopher can eat at any time.
- This prevents circular wait, a necessary condition for deadlock.
*/
