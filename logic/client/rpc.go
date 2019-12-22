package client

import (
	"log"
	"net/rpc"
	"notify/logic/server"
)

func AddNotify(title,content string) error {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		return err
	}
	notifyItem := server.Notify{
		Message:server.Message{
			Title:title,
			Content:content,
		},
		Type:server.NotifyTypeInterval,
		NotifyInterval:2,
		NotifyTimePoint:0,
		LastNotifyTimePoint:0,
		IsClosed:false,
	}

	var reply server.NotifyAddResponse
	err = client.Call("RpcNotify.AddNotify", notifyItem, &reply)
	if err != nil {
		return err
	}
	log.Println(reply)
	return nil
}