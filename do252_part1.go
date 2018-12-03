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
	lect := make(chan bool) // creates a synchronous channel
	wait := make(chan chan string, n) // creates an asynchronous channel of size n
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

			// meeting duration
			s := rand.Intn(5000)+1000
			time.Sleep(time.Duration(s) * time.Millisecond)

			// tell std meeting is over
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
		select {
		case <-lect:
			fmt.Println(name, "is in a meeting!")
			time.Sleep(time.Millisecond * 5000) //meeting duration
			lect <- true //let the lecturer know that the meeting is over
			fmt.Println("Meeting with ", name, "has finished")
		default:
			wait <- std
			fmt.Println(name, "is waiting")
			<-std //For the lecturer to get back to the student via handshake.
			fmt.Println(name, "is in a meeting!")
			// in a meeting
			<-std
			fmt.Println(name, "finished meeting.")
		}

}

//func student(wait chan chan string, lect chan bool, name string){
//	std := make(chan string)
//	std <- name
//	select{
//	case <-lect:
//		fmt.Println(name, "is in a meeting!")
//		time.Sleep(time.Millisecond * 5000)
//		fmt.Println("Meeting with ", name, "has finished")
//	default:
//		wait <- std
//		fmt.Println(name, "is waiting")
//	}
//}

/*func lecturer(wait chan chan string, lect chan bool){
	select{
		case stud := <- wait:
			fmt.Println("New student is called")
			//fmt.Println(<-stud, " is in a meeting!")
			<-stud
			s := rand.Intn(5000)
			time.Sleep(time.Duration(s) * time.Millisecond)
		default:
			time.Sleep(time.Second)
	}
}*/

/*func student(wait chan chan string, lect chan bool, name string){
	select {
	case <- lect:
		fmt.Println(name, "is in a meeting!")
		<- wait
		fmt.Println("Meeting with", name, "has ended")
	default:
		st := make(chan string)
		st <- name
		wait <- st
		fmt.Println(name, " is waiting")
		<- lect
	}
}*/