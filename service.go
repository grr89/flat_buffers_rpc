package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	test "github.com/grr89/fb_rpc/Test"
)

type Some byte //any type for write remote methods

func (t *Some) ConcatInfo(buf []byte, s *string) error {
	person := test.GetRootAsPerson(buf, 0)                                                            //get entry point of class of flatbuffers
	*s = person.Name() + " " + person.Surname() + ", " + strconv.Itoa(int(person.Age())) + " age old" //concate struct fields and store to container
	return nil                                                                                        //no errors
}

func main() {
	s := new(Some)
	rpc.Register(s) //register type
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(l, nil) //start listening
}
