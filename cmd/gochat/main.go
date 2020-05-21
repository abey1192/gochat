package main

import (
	"fmt"
	"strings"
	"time"
)

func dispatch_message(broadcast chan string, members map[string]chan string) {

	for {
		msg := <-broadcast

		ss := strings.Split(msg, "/")
		id := ss[0]
		m := ss[1]

		ms := id + " : " + m

		for pid, _ := range members {
//			ch <- ms

			fmt.Println(pid + " <- " + ms)
		}
	}
}

func receive_message(id string, recv chan string, broadcast chan string) {

	do_loop := true

	for do_loop {
		msg := <-recv
		do_loop = msg != "quit"
		broadcast <- id + "/" + msg
		fmt.Println("Broadcasting message " + msg + " by " + id)
	}
}

func main() {
	room := NewRoom("Global")
	go room.Run()

	time.Sleep(2000)

	u1 := NewUser("A")
	u1.Run()
	u2 := NewUser("B")
	u1.Run()
	u3 := NewUser("C")
	u2.Run()

	room.Subscribe(u1)
	room.Subscribe(u2)
	room.Subscribe(u3)

	u1.Says("Global", "Hello")
	u2.Says("Global", "Hi")
	u1.Says("Global", "yoyo")

	time.Sleep(time.Second * 3)
}
