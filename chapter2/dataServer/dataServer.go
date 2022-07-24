package main

import (
	"dataServer/objects"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	r := gin.Default()

	objects.Handlers(r)

	r.Run(os.Getenv("LISTEN_ADDRESS"))
	
}
