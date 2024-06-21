package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := &pb.Person{
		Name:   "me",
		Age:    18,
		Emails: []string{"111"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "189",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "188",
				Type:   pb.PhoneType_MOBILE,
			},
		},
	}
	data, err := proto.Marshal(person)
	if err != nil {
		return
	}
	newData := &pb.Person{}
	err = proto.Unmarshal(data, newData)
	if err != nil {
		return
	}
	fmt.Println(newData)
}
