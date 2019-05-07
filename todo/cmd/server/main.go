package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sivsivsree/grpc/todo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type taskServer struct {
}

func (ts taskServer) List(context.Context, *todo.Void) (*todo.TaskList, error) {
	return nil, fmt.Errorf("Not Implemented yet")
}

func main() {

	srv := grpc.NewServer()

	var tasks taskServer

	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Errorf("TCP server failed. %v", err)
	}

	log.Fatal(srv.Serve(l))
}

const dbPath = "mydb.pb"

func add(text string) error {

	task := &todo.Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)

	if err != nil {
		fmt.Errorf("Unable to encode protocol buffer")
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Errorf("Database cannot be found %s : %v", dbPath, err)
	}

	_, werr := f.Write(b)

	if werr != nil {
		fmt.Errorf("coudnt write to database %s : %v", dbPath, err)
	}

	f.Close()

	return nil
}

func list() error {
	return nil
}
