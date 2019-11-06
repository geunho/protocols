package echo

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestMarshal(t *testing.T) {
	req := &EchoRequest{Name: "Geunho"}
	
	data, err := proto.Marshal(req)

	if err != nil {
		t.Error(err)
	}

	//[10 6 71 101 117 110 104 111]
	t.Log(data)
}

func TestUnmarshal(t *testing.T) {
	res := &EchoRequest{}

	//&echo.EchoRequest{Name: "Geunho"}
	data := []byte{10, 6, 71, 101, 117, 110, 104, 111}

	err := proto.Unmarshal(data, res)

	if err != nil {
		t.Error(err)
	}

	//name:"Geunho"
	t.Log(res)
}