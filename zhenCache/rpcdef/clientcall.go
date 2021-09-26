package rpcdef

import (
	"basic/zhenCache/innerDB/keystruct"
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func call(rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}

//get some value by cli
func Get(key keystruct.KeyStruct) (interface{}, error) {
	args := StoreArgs{
		Command: GET,
		Key:     key,
		Value:   nil,
	}
	reply := StoreReply{}
	call("Coordinator.Get", &args, &reply)
	if reply.Reply == SUCCESS {
		return reply.Value, nil
	}
	return nil, errors.New("key not found")
}

//set some value by cli
func Set(key keystruct.KeyStruct, value interface{}, expire time.Duration) error {
	args := StoreArgs{
		Command: SET,
		Key:     key,
		Value:   value,
		Expire:  expire,
	}
	reply := StoreReply{}
	call("Coordinator.Set", &args, &reply)
	if reply.Reply == SUCCESS {
		return nil
	}
	return errors.New("set fail")
}