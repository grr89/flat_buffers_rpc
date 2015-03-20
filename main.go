package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/google/flatbuffers/go"
	test "github.com/grr89/fb_rpc/Test"
)

func main() {

	builder := flatbuffers.NewBuilder(0)
	name := builder.CreateString("Vitalik") //preparing strings
	surname := builder.CreateString("Makarevich")
	test.PersonStart(builder)
	test.PersonAddName(builder, name) //strings must be passed from variable
	test.PersonAddSurname(builder, surname)
	person := test.PersonEnd(builder)
	builder.Finish(person)
	buf := builder.Bytes[builder.Head():] //get clean data without header

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234") //connect to rpc server
	if err != nil {
		log.Fatal(err)
	}

	var s string

	now := time.Now() //time calculation start

	for i := 1; i <= 1000; i++ {
		err := client.Call("Some.ConcatInfo", buf, &s) //call remote procedure *Some.ConcatInfo with 2 params out buffer+container
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s) //output container
	}

	fmt.Println(time.Since(now)) //time calculation end
}
