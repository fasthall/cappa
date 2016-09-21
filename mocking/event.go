package main

import (
	"fmt"
	"github.com/fasthall/cappa/mq"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("go run event.go ID PAYLOAD")
		return
	}
	mqc, err := mq.NewMQ()
	if err != nil {
		fmt.Println(err)
		return
	}
	data := mq.Data{
		Id:      os.Args[1],
		Payload: []byte(os.Args[2]),
	}
	mqc.Send(data)
}
