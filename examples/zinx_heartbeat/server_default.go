package main

import (
	"github.com/iaoizo/zinx/znet"
	"time"
)

func main() {
	s := znet.NewServer()

	// Start heartbeating detection.
	s.StartHeartBeat(5 * time.Second)

	s.Serve()
}
