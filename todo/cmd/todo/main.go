package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sivsivsree/grpc/todo"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		_, _ = fmt.Fprintln(os.Stderr, "Missing subcomands list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("Connection Failed %v", err)
		os.Exit(0)
	}

	client := todo.NewTasksClient(conn)

	switch flag.Arg(0) {

	case "list":

		err = list(context.Background(), client)

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

func add(text string) error {

	return fmt.Errorf("Not implemented")
}

func list(ctx context.Context, client todo.TasksClient) error {

	l, err := client.List(ctx, &todo.Void{})

	if err != nil {
		fmt.Errorf("Coudnt fetch list %v", err)
		return err
	}

	for _, t := range l.Tasks {
		if t.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("👎")
		}

		fmt.Printf(" %s\n", t.Text)
	}

	return nil
}
