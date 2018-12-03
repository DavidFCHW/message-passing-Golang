package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"time"
)

func main(){
	m := 10
	n := 6
	lect := make(chan bool)
	wait := make(chan chan string, n)
	go lecturer(wait,lect)
	for i:=0 ; i < m ; i++ {
		go student(wait, lect, randomdata.SillyName())
	}
	time.Sleep(100 * time.Second)
}

func lecturer(wait chan chan string, lect chan bool) {
	for{
		select {
		case stdChan := <-wait:
			fmt.Println("New student called")
			stdChan <- "wake up"


			s := rand.Intn(5000)+1000
			time.Sleep(time.Duration(s) * time.Millisecond)

			stdChan <- "meeting is over"
		default:
			lect <- true
			<- lect
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func student(wait chan chan string, lect chan bool, name string){
	std := make(chan string)
	fmt.Println(name, "wants to meet")
	select {
	case <-lect:
		fmt.Println(name, "is in a meeting!")
		time.Sleep(time.Millisecond * 5000)
		lect <- true
		fmt.Println("Meeting with ", name, "has finished")
	default:
		select {
		case wait <- std:
			fmt.Println(name, "is waiting")
			<-std
			fmt.Println(name, "is in a meeting!")
			<-std
			fmt.Println(name, "finished meeting.")
		default:
			fmt.Println(name, "says bye!")
		}
	}

}