package main

import "google.golang.org/grpc"

import "github.com/ronaldoveras/go-labs/grpc-test/01-Proto/echo"

import "context"

import "fmt"

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := echo.NewEchoServerClient(conn)
	resp, err := e.Echo(ctx, &echo.EchoRequest{
		Message: "Hello World!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Got from server: ", resp.Response)
}
