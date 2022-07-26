package locate

import (
	"dataServer/constant"
	"github.com/1120475708/common/rabbitmq"
	"github.com/1120475708/common/utils"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv(constant.ApiServer))
	defer q.Close()
	q.Bind(constant.DataServer)
	c := q.Consume()
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(utils.GetPrefixPath() + object) {
			q.Send(msg.ReplyTo, os.Getenv(constant.ListenAddress))
		}
	}
}
