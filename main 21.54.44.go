package main

import (
	"github.com/RoyJoel/TennisMomentBackEnd/cmd/grpc"
	"github.com/RoyJoel/TennisMomentBackEnd/cmd/web"
)

func main() {
	go grpc.Run()
	go web.Run()
	select {}
}
