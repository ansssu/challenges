package main

import (
	"fmt"
	"reflect"

	pb "github.com/andersonlribeiro/challenges/go/protobufferDemo/internal/domain/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

func main() {

	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	fmt.Println(p.Name)
	fmt.Println(pb.Person_HOME.Enum())
	var phone pb.Person_PhoneType
	println(phone)
	if IsNil(phone) {
		println("nil")
	}

}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
