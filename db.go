package achieve_db_go

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"time"
)

var client *tarantool.Connection

func Init() {
	opts := tarantool.Opts{
		Timeout:       500 * time.Second,
		Reconnect:     1 * time.Second,
		MaxReconnects: 3,
		User:          tarantoolUser,
		Pass:          tarantoolPassword,
	}
	tempClient, err := tarantool.Connect(tarantoolHost, opts)
	client = tempClient
	if err != nil {
		fmt.Println("Connection refused:", err)
	}

}

func Select(space string, index string, query string) *tarantool.Response {
	if client == nil {
		Init()
	}
	resp, err := client.Select(space, index, 0, 1, tarantool.IterEq, []interface{}{query})
	if err != nil {
		log.Println("Error on query:", err)
	}
	return resp

}

func Insert(space string, obj []interface{}) *tarantool.Response {
	if client == nil {
		Init()
	}
	resp, err := client.Insert(space, obj)
	if err != nil {
		log.Println("Error on query:", err)
	}
	return resp
}

func Delete(space string, index string, query string) *tarantool.Response {
	if client == nil {
		Init()
	}
	resp, err := client.Delete(space, index, []interface{}{query})
	if err != nil {
		log.Println("Error on query:", err)
	}
	return resp
}

func SelectUsers(space string, index string, query string) *[]User {
	if client == nil {
		Init()
	}
	var res []User
	err := client.SelectTyped(space, index, 0, 1, tarantool.IterEq, []interface{}{query}, &res)
	if err != nil {
		log.Println("Error on query:", err)
	}
	return &res

}

func SelectConnections(space string, index string, query string) *[]User {
	if client == nil {
		Init()
	}
	var res []User
	err := client.SelectTyped(space, index, 0, 1, tarantool.IterEq, []interface{}{query}, &res)
	if err != nil {
		log.Println("Error on query:", err)
	}
	return &res

}