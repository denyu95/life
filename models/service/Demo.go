package service

import (
	"fmt"

	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/qq/event"
)

func Demo(msg event.PrivateMsg) {
	demo := dao.Demo{
		Name: msg.Message,
	}
	demo.AddDemo()
}

func Demo1(msg event.PrivateMsg) {
	fmt.Println("大写" + msg.Message)
}
