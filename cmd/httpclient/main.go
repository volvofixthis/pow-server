package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/volvofixthis/pow-server/internal/adapters/handlers/dtos"
	"github.com/volvofixthis/pow-server/internal/core/utils"
	"github.com/volvofixthis/pow-server/internal/infra/config"
)

func main() {
	cfg := config.NewCfg()

	taskURL := fmt.Sprintf("http://%s/v1/pow", cfg.ApiAddress)
	passageURL := fmt.Sprintf("http://%s/v1/passage", cfg.ApiAddress)

	body := []byte{}
	r, err := http.NewRequest("POST", taskURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	task := &dtos.PowTaskResp{}
	err = json.NewDecoder(res.Body).Decode(task)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Received text %s and salt %x\n", task.Text, task.Salt)

	hash := utils.GenerateProofOfWork(task.Text, task.Salt, task.Iteration, task.Memory)
	passageReq := &dtos.PassageReq{Hash: hash}
	body, err = json.Marshal(passageReq)
	if err != nil {
		log.Println(err)
		return
	}
	r, err = http.NewRequest("POST", passageURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	res, err = client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	passage := &dtos.PassageResp{}
	err = json.NewDecoder(res.Body).Decode(passage)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Listen pal:", passage.Text)
}
