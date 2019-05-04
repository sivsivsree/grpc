package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sivsivsree/grpc/todo"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		_, _ = fmt.Fprintln(os.Stderr, "Missing subcomands list or add")
		os.Exit(1)
	}

	var err error
	switch flag.Arg(0) {

	case "list":
		err = list()

	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))

	default:
		err = fmt.Errorf("Unknown command")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
