package main

import (
	complexpb "complex"
	enumpb "enum_example"
	"fmt"
	"io/ioutil"
	"log"
	simplepb "simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "Robin",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Maria",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "David",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	ep := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}

	ep.DayOfTheWeek = enumpb.DayOfTheWeek_SUNDAY
	fmt.Println(ep)
}

func jsonDemo(sm proto.Message) {
	smAsJSON := toJSON(sm)
	fmt.Println("The protobuffer in JSON is", smAsJSON)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsJSON, sm2)

	fmt.Println("Read from JSON into the following protobuffer: ", sm2)
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Could not convert JSON into protobuffer", err)
	}
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Could not unmarshal the proto into JSON", err)
		return ""
	}

	return out
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content: ", sm2)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Could not put the butes into the protocol buffer struct", err)
		return err2
	}

	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Robin",
		SampleList: []int32{1, 4, 7, 8, 9},
	}

	sm.Name = "Maria"
	fmt.Println("The ID is", sm.GetId())
	fmt.Println("The name is", sm.GetName())

	return &sm
}
