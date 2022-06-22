package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type ChopStick struct { sync.Mutex }

type Philosopher struct {
	philosopherID int
	csLeft, csRight *ChopStick
}

func (philosopher Philosopher) Eat(dishes int, host *Host) {
	for i := 0; i < dishes; i++ {
		host.GivePermission(&philosopher)

		rand.Seed(time.Now().UnixNano())
		order := rand.Intn(2)
		if order == 0 {
			philosopher.csLeft.Lock()
			philosopher.csRight.Lock()
		} else {
			philosopher.csRight.Lock()
			philosopher.csLeft.Lock()
		}
		
		fmt.Printf("Philosopher %d started eating for the %d time\n", philosopher.philosopherID, i + 1)

		fmt.Printf("Philosopher %d finished eating for the %d time\n", philosopher.philosopherID, i + 1)

		philosopher.csLeft.Unlock()
		philosopher.csRight.Unlock()

		if i == dishes - 1 {
			fmt.Printf("Philosopher %d has eaten %d dishes and will not eat anymore\n", philosopher.philosopherID, dishes)
		} 

		host.ReleaseCS()
	}
}

type Host struct { 
	channel chan *Philosopher 
	dishesEaten, maxDishes int
}

func (host *Host) ControlDishes(wg *sync.WaitGroup) {
	for {
		if host.dishesEaten >= host.maxDishes {
			wg.Done()
			return
		}
	}
}

func (host *Host) GivePermission(philosopher *Philosopher) {
	host.channel <- philosopher
}

func (host *Host) ReleaseCS() {
	<- host.channel
	host.dishesEaten++
}

func main() {
	const pCount = 5
	const csCount = pCount
	const pDishes = 3
	const syncPhilosophers = 2

	host := Host{ make(chan *Philosopher, syncPhilosophers), 0, pDishes * pCount }

	chopsticks := make([]*ChopStick, csCount)
	for i := 0; i < csCount; i++ {
		chopsticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, pCount)
	for i := 0; i < pCount; i++ {
		philosophers[i] = &Philosopher{ i + 1, chopsticks[i], chopsticks[(i + 1) % 5] }
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go host.ControlDishes(&wg)

	for _, philosopher := range philosophers {
		go philosopher.Eat(pDishes, &host)
	}

	wg.Wait()
	fmt.Println("All philosophers have stopped eating")
}