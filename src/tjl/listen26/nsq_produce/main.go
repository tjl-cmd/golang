package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

var producer *nsq.Producer

func main() {
	//nsq的地址
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Println("init producer failed ", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string failed", err)
			continue
		}
		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}
		err = producer.Publish("order_queue", []byte(data))
		if err != nil {
			fmt.Println("publish message failed err,", err)
			continue
		}
		fmt.Println("publish data ", data)
	}
}

//初始化生成者
func initProducer(str string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		return err
	}
	return nil
}
