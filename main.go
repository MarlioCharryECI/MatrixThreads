package main

import (
	"fmt"
	"time"

	"matrix_concurrency/config"
	"matrix_concurrency/game"
)

func main() {
	world := game.NewWorld(config.Default)
	done := make(chan string)

	go game.NeoRoutine(world, done)
	for i := range world.Agents {
		go game.AgentRoutine(world, i, done)
	}

	for {
		select {
		case msg := <-done:
			world.Print()
			fmt.Printf("\n%s\n", msg)
			return
		default:
			world.Print()
			time.Sleep(300 * time.Millisecond)
		}
	}
}
