package main

import (
	"fmt"

	"github.com/cosmos-gg/meq/proto"
	meq "github.com/cosmos-gg/meq/sdks/go-meq"
)

func sub(conn *meq.Connection) {
	conn.OnMessage(func(msg *proto.PubMsg) {
		fmt.Println("recv msg:", string(msg.ID), string(msg.Topic), string(msg.Payload))
	})
	conn.OnUnread(func(topic []byte, count int) {
		fmt.Println("未读消息数量：", string(topic), count)

		conn.ReduceCount([]byte(topic), proto.MAX_PULL_COUNT)
	})

	err := conn.Subscribe([]byte(topic))

	if err != nil {
		panic(err)
	}

	// 先拉取x条消息
	err = conn.PullMsgs([]byte(topic), proto.MAX_PULL_COUNT, proto.MSG_NEWEST_OFFSET)
	if err != nil {
		fmt.Println(err)
	}
	select {}

	// fmt.Println("累积消费未ACK消息数：", n1)
}