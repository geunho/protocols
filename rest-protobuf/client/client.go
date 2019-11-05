package main

import (
	"bytes"
	"fmt"
	echo "github.com/geunho/protocols/rest-protobuf/proto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func makeRequest(request *echo.EchoRequest) *echo.EchoResponse {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	res, err := http.Post("http://0.0.0.0:8888/echo", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	resObj := &echo.EchoResponse{}
	proto.Unmarshal(resBytes, resObj)

	return resObj
}

func main() {
	if len(os.Args) < 1 {
		log.Fatalln("Needs name parameter...")
	}

	name := os.Args[1]

	request := &echo.EchoRequest{Name: name}
	response := makeRequest(request)
	fmt.Printf("Response from API is : %v\n", response.GetMessage())
}