package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type ChopS struct{ sync.Mutex }

type Philo struct {
	name int
	leftCS, rightCS *ChopS
}
func (a *ChopS) Pickup() {
	a.Lock()
}
func (a *ChopS) Putdown() {
	a.Unlock()
}
func (a *Philo) Eat() {
	fmt.Println(fmt.Sprintf("starting to eat %d", a.name))
}
func (a *Philo) Finish() {
	fmt.Println(fmt.Sprintf("finishing eating %d", a.name))
}

 
func main() {
	rand.Seed(time.Now().UnixNano())
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
	   CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i+1, CSticks[i], CSticks[((i+1)%5)]}
	}
	 channel := make(chan struct{}, 2)
	 var wg sync.WaitGroup
	 for i := 0; i< 5; i++ {
		 wg.Add(1)
		 go dine(&wg, channel, philos[i])
	 }
	 wg.Wait()
}

func dine(wg *sync.WaitGroup, channel chan struct{}, philosopher *Philo) {
	for i := 0; i < 3; i++ {
		channel <- struct{}{}
		// randomize picking chopsticks
		randBinaryValue := rand.Intn(2)
		// fmt.Println(randBinaryValue)
		if randBinaryValue == 0 {
			philosopher.leftCS.Pickup()
			philosopher.rightCS.Pickup()
		} else {
			philosopher.rightCS.Pickup()
			philosopher.leftCS.Pickup()
		}

		philosopher.Eat()
		philosopher.Finish()
		if randBinaryValue == 0 {
			philosopher.leftCS.Putdown()
			philosopher.rightCS.Putdown()
		} else {
			philosopher.rightCS.Putdown()
			philosopher.leftCS.Putdown()
		}
		<- channel
	}
	wg.Done()
}