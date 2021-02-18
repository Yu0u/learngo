package main

import (
	"fmt"
	"learngo/ini"
)

type test struct {
	port string
}

func newTest(port string) *test {
	return &test{port: port}
}

func main() {
	cfg := ini.SetConfig("./ini/test.ini")
	value := cfg.GetValue("server", "port")
	t := newTest(value)
	fmt.Println(t.port)
}
