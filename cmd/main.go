package main

import (
	"context"
	"fmt"

	"github.com/yurchenkovr/chat-stress-test/internal/request"
)

func main() {
	cl := request.NewClient()

	ctx := context.Background()

	// 1. register new users
	_, err := cl.Register(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. login users
	_, err = cl.Login(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// todo 1. open ws connection
	// 2. send messages concurrently for all connections
	// 3. output some logs (time..)

	fmt.Println("Success")
}
