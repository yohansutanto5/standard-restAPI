package test

import (
	"app/pkg/util"
	"fmt"
	"log"
	"testing"
)

type SourceStruct struct {
	SourceField1 int          `convert:"TargetField1"`
	SourceField2 string       `convert:"TargetField2"`
	NestedStruct NestedStruct `convert:"NestedTargetStruct"`
}

type TargetStruct struct {
	TargetField1 int
	TargetField2 string
	NestedStruct NestedTargetStruct
}

type NestedStruct struct {
	NestedField1 int `convert:"NestedTargetField1"`
	NestedField2 string
}

type NestedTargetStruct struct {
	NestedTargetField1 int
	NestedTargetField2 string
}

func TestStructConverter(t *testing.T) {
	sourceInstance := SourceStruct{
		SourceField1: 42,
		SourceField2: "Hello, World!",
		NestedStruct: NestedStruct{
			NestedField1: 24,
			NestedField2: "Nested Hello!",
		},
	}

	// Convert the source struct to the target struct
	var targetInstance TargetStruct
	util.ConvertStruct(sourceInstance, &targetInstance)
	fmt.Printf("Source Struct: %+v\n", sourceInstance)
	fmt.Println(fmt.Printf("Target Struct: %+v\n", targetInstance))
	log.Fatal(targetInstance)
}


func TestStructConverterToMap(t *testing.T) {
	sourceInstance := SourceStruct{
		SourceField1: 42,
		SourceField2: "Hello, World!",
		NestedStruct: NestedStruct{
			NestedField1: 24,
			NestedField2: "Nested Hello!",
		},
	}
	fieldsToInclude := []string {"TargetField1","TargetField2","NestedTargetStruct","NestedTargetField1"}
	mapped := util.ConvertStructToMap(sourceInstance,fieldsToInclude)
	log.Fatal(fmt.Printf("Target Struct: %+v\n", mapped))
}