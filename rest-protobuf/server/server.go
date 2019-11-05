package main

import (
	"fmt"
	echo "github.com/geunho/protocols/rest-protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Echo(rw http.ResponseWriter, req *http.Request) {
	request := &echo.EchoRequest{}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	err = proto.Unmarshal(data, request)
	if err != nil {
		log.Fatalf("Unable to unmarshal message : %v", err)
	}

	name := request.GetName()
	result := &echo.EchoResponse{Message: "Hello " + name}

	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	rw.Write(response)
}

func main() {
	fmt.Println("Start Echo API server...")

	r := mux.NewRouter()
	r.HandleFunc("/echo", Echo).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr: "0.0.0.0:8888",
		WriteTimeout: time.Second,
		ReadTimeout: time.Second,
	}

	log.Fatal(server.ListenAndServe())
}