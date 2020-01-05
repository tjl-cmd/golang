package log

import "fmt"

type ConsoleLog struct {
}

func (c *ConsoleLog) LogDebug(msg string) {
	fmt.Println("console", msg)
}

func (c *ConsoleLog) LogWarn(msg string) {
	fmt.Println("console", msg)
}

func NewConsoleLog(file string) LogInterface {
	return &ConsoleLog{}
}
