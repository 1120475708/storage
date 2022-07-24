package objects

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func get(c *gin.Context) {
	path, _ := os.Getwd()
	f, err := os.Open(path + storageRoot + "/objects/" + c.Param("name"))
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Writer, http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(c.Writer, f)

}
