package server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	Success = iota + 1000
	ErrMarshal
	ErrUnMarshal
)

type NotifyAddRequest Notify
type NotifyAddResponse struct {
	Errno int
	ErrMsg string
}

type RpcNotify int

func (rn *RpcNotify) AddNotify(req *NotifyAddRequest,resp *NotifyAddResponse) error {
	bReq,err := json.Marshal(req)
	if err != nil {
		resp.Errno = ErrMarshal
		resp.ErrMsg = GetErrMsg(resp.Errno)
		return errors.New(resp.ErrMsg)
	}

	log.Println("Fitz：",string(bReq))
	var notifyItem = &Notify{}
	err = json.Unmarshal(bReq,notifyItem)
	if err != nil {
		log.Println("Fitz：",err)
		resp.Errno = ErrUnMarshal
		resp.ErrMsg = GetErrMsg(resp.Errno)
		return errors.New(resp.ErrMsg)
	}
	NotifyList = append(NotifyList,notifyItem)
	resp.Errno = Success
	resp.ErrMsg = GetErrMsg(resp.Errno)
	return nil
}

func GetErrMsg(errCode int) string {
	ErrMsg :=  map[int]string{
		Success: "Fine",
		ErrMarshal: "Marshal Fail",
		ErrUnMarshal: "UnMarshal Fail",
	}
	return ErrMsg[errCode]
}

func RpcHandler() error{
	var listener net.Listener
	rpcHandler := new(RpcNotify)
	err := rpc.Register(rpcHandler)
	if err != nil {
		return err
	}
	rpc.HandleHTTP()
	listener, err = net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	log.Println("RpcService Start Listen")
	err = http.Serve(listener, nil)
	if err != nil {
		return err
	}
	log.Println("RpcService Started")
	return nil
}