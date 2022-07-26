package main

import (
	"dataServer/heartbeat"
	"dataServer/locate"
	"dataServer/objects"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	r := gin.Default()
	objects.Handlers(r)
	r.Run(os.Getenv("LISTEN_ADDRESS"))

}
