package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/volvofixthis/pow-server/internal/adapters/handlers/dtos"
	"github.com/volvofixthis/pow-server/internal/core/utils"
	"github.com/volvofixthis/pow-server/internal/infra/config"
)

func main() {
	cfg := config.NewCfg()

	conn, err := net.Dial("tcp", cfg.TCPAddress)
	if err != nil {
		log.Println("Error connecting to TCP server:", err)
		return
	}
	defer conn.Close()

	// Request challenge
	powHelloReq := &dtos.PowHelloReq{State: dtos.RequestState}
	body, err := json.Marshal(powHelloReq)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := conn.Write(body); err != nil {
		log.Println(err)
		return
	}

	// Receive and decode task
	task := &dtos.PowTaskResp{}
	err = json.NewDecoder(conn).Decode(task)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Received text \"%s\" and salt %x\n", task.Text, task.Salt)

	// Generate hash
	hash := utils.GenerateProofOfWork(task.Text, task.Salt, task.Iteration, task.Memory)
	// Request passage
	passageReq := &dtos.PassageReq{Hash: hash}
	body, err = json.Marshal(passageReq)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := conn.Write(body); err != nil {
		log.Println(err)
		return
	}

	// Receive and decode passage
	passage := &dtos.PassageResp{}
	err = json.NewDecoder(conn).Decode(passage)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Listen pal:", passage.Text)
}
