package copystruct

import (
	"fmt"
	"log"
	"testing"
	"time"
)

//CopyStruct from source to target with that shares the same field name and field type
func TestCopyStruct(t *testing.T) {
	type SourceStruct struct {
		ID       int
		Name     string
		Value    int64
		Extra    int
		DontCopy time.Duration
	}

	type TargetStruct struct {
		ID      int
		Name    string
		Value   int64
		Extra   string
		PlzCopy time.Duration
	}

	sourceStruct := SourceStruct{ID: 1, Name: "Testing", Value: 999999999999999999, Extra: 3, DontCopy: time.Second}
	targetStruct := TargetStruct{ID: 2, Name: "Control", Value: 9, Extra: "extra", PlzCopy: time.Minute}

	if err := CopyStruct(&sourceStruct, &targetStruct); err != nil {
		log.Fatal(err)
	}

	//compare them
	if targetStruct.ID != sourceStruct.ID || targetStruct.Name != sourceStruct.Name || targetStruct.Value != sourceStruct.Value ||
		targetStruct.Extra != "extra" || targetStruct.PlzCopy != time.Minute {
		log.Fatal("Copy failed")
	}

	fmt.Println(targetStruct)
}
