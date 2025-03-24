package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"todo_cli/transport/param"
)

func main() {
	fmt.Println("Command", os.Args[0])
	message := "default message"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("can't dial the given address:", err)
	}
	defer conn.Close()

	fmt.Println("local address:", conn.LocalAddr())

	req := param.Request{
		Command: message,
	}
	if req.Command == "create-task" {
		req.CreateTaskRequest = param.CreateTaskRequest{
			Title:      "test",
			DueDate:    "test",
			CategoryID: 1,
		}
	}
	serializedData, err := json.Marshal(&req)
	if err != nil {
		log.Fatalln("can't serialize request:", err)
	}

	numberOfWriteBytes, wErr := conn.Write(serializedData)
	if wErr != nil {
		log.Fatalln("can't write data to connection", wErr)
	}
	fmt.Println("number of bytes written:", numberOfWriteBytes)

	var data = make([]byte, 1024)
	_, rErr := conn.Read(data)
	if rErr != nil {
		log.Fatalln("cannot read from connection:", rErr)
	}
	log.Println("read data:", string(data))
}
