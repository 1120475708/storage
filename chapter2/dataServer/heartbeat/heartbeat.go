package heartbeat

import (
	"dataServer/constant"
	"github.com/1120475708/common/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv(constant.RabbitMQServer))
	defer q.Close()
	for {
		q.Publish(constant.ApiServer, os.Getenv(constant.ListenAddress))
		time.Sleep(5 * time.Second)
	}
}
