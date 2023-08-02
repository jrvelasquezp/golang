//simple program
//there are two resources and the arbitrator tries to assign the resource
//arbitrator cannot assign if the resource is busy

package main

import (
	"fmt"
	"sync"
	"time"
)

type Seat struct {
	sync.Mutex
	user int
	busy bool
	id int
}

func Host (seats []*Seat, user int, wg *sync.WaitGroup) {
	for i:= range seats {
		if !seats[i].busy {
			seats[i].Lock()
			seats[i].user=user
			seats[i].busy=true
			fmt.Printf("Seat %d is assigned to user %d\n",i,user)
			time.Sleep(time.Second)
			go Free(seats[i])
			//eat part should be here
			break;
		}
	}
	wg.Done()
}

func Free (seat *Seat) {
	if (seat.busy) {
		seat.busy=false
		fmt.Printf("The seat %d is released by %d\n", seat.id,seat.user)
		seat.user=0
		seat.Unlock()
	} else {
		fmt.Printf("Seat %d is already free\n", seat.id)
	}
	//wg.Done()
}

func main() {
	fmt.Println("Test")
	//initialize two seats
	seats:=make([]*Seat,2)
	for i:= range seats {
		seats[i]=new(Seat)
		seats[i].user=0
		seats[i].busy=false
		seats[i].id=i
	}
	//setting a waitgroup
	var wg sync.WaitGroup
	for i:=0; i<10; i++ {
		wg.Add(5)
		go Host(seats,1,&wg)
		go Host(seats,4,&wg)
		go Host(seats,2,&wg)
		go Host(seats,3,&wg)
		go Host(seats,5,&wg)
		wg.Wait()
	}
	time.Sleep(10*time.Second)
}