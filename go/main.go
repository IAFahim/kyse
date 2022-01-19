package main

import (
	"fmt"
	"github.com/daspoet/gowinkey"
	"time"
)

func main() {
	events, stopFn := gowinkey.Listen()

	time.AfterFunc(time.Second, func() {
		stopFn()
	})

	for e := range events {
		switch e.State {
		case gowinkey.KeyDown:
			fmt.Println("pressed", e)
		case gowinkey.KeyUp:
			fmt.Println("released", e)
		}
	}
}
