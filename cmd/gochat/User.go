package main

import "strings"

type User struct {
	id string
	channel chan string
	rooms map[string] *Room
}

func NewUser(id string) *User {
	return &User{id: id, channel: make(chan string), rooms: make(map[string] * Room) }
}

func (user *User) Send(id string, msg string) {
	user.channel <- id + " : " + msg
}

func (user *User) runChatLoop() {

	for {
		msg := <- user.channel

		fragments := strings.Split(msg, ":")
		roomId := fragments[0]
		message := fragments[1]

		user.rooms[roomId].Broadcast(user.id, message)
	}
}


func (user *User) Run() {
	go user.runChatLoop()
}

func (user *User) Says(roomId string, message string) {
	user.channel <- roomId + ":" + message
}
