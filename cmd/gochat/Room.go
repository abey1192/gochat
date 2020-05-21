package main

import (
	"fmt"
	"strings"
)

type Room struct {
	id string
	channel chan string
	users map[string] *User
}

func NewRoom(id string) *Room {
	return &Room{id : id, channel : make(chan string), users: make(map[string] *User) }
}

func (room *Room) Subscribe(user *User) {
	room.users[user.id] = user
	user.rooms[room.id] = room
}

func (room *Room) Unsubscribe(user *User) {
	delete(room.users, user.id)
	delete(user.rooms, room.id)
}

func (room *Room) Broadcast(userId string, msg string) {
	room.channel <- userId + ":" + msg
}

func (room *Room) Run() {

	for {
		msg := <- room.channel

		fragments := strings.Split(msg, ":")
		userId := fragments[0]
		message := fragments[1]

		for id, _ := range(room.users) {
			fmt.Println(id + ":" +  userId + " says " + message)
		}
	}
}