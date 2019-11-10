package copystruct

import (
	"errors"
	"reflect"
)

//CopyStruct copies everything from source to target by looking up for fields with same name and type
func CopyStruct(sourceStruct, targetStruct interface{}) error {
	//Make sure the incoming interface are ptr of struct
	sourceT := reflect.TypeOf(sourceStruct)
	targetT := reflect.TypeOf(targetStruct)

	if sourceT.Kind() != reflect.Ptr || targetT.Kind() != reflect.Ptr {
		return errors.New("Source and Target must be pointer of a struct")
	}

	//Get struct
	sourceT = sourceT.Elem()
	targetT = targetT.Elem()

	if sourceT.Kind() != reflect.Struct || targetT.Kind() != reflect.Struct {
		return errors.New("Source and Target must be pointer of a struct")
	}

	sourceV := reflect.ValueOf(sourceStruct).Elem()
	targetV := reflect.ValueOf(targetStruct).Elem()

	for i := 0; i < sourceT.NumField(); i++ {
		sField := sourceT.Field(i)
		//Trying to get the filed from target struct
		if tField, exists := targetT.FieldByName(sField.Name); exists && tField.Type == sField.Type {
			if tValue := targetV.FieldByName(tField.Name); tValue.CanSet() {
				tValue.Set(sourceV.Field(i))
			}
		}
	}

	return nil
}
