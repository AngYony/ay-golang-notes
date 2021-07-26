package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {

	err := rpc.RegisterName("HelloService", &HelloService{})

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		err = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			log.Fatal(err)
		}

	})

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}

}
