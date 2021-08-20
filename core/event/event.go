package event

import (
	"fmt"

	"wms.com/core/queue"
)

type Subscriber func(*queue.Conn)

func Subscribe(subscribers ...Subscriber) {
	conn, err := queue.GetConn()
	if err != nil {
		fmt.Println(err)
	}
	for _, subscriber := range subscribers {
		subscriber(&conn)
	}
}
