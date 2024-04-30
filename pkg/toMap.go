package pkg

import "reflect"

func toLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]+32) + s[1:]
}

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	if objType.Kind() != reflect.Struct {
		if objType.Kind() == reflect.Ptr {
			if objValue.IsNil() {
				return nil
			}
			objValue = objValue.Elem()
			objType = objValue.Type()
		}
	}

	data := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)

		fieldValue := objValue.Field(i)
		if field.Name == "state" || field.Name == "sizeCache" || field.Name == "unknownFields" {
			continue
		}

		if field.Tag.Get("json") == "-" {
			continue
		}

		// Если тип поля является указателем, следует разыменовать его
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue
			}
			fieldValue = fieldValue.Elem()
		}

		fieldName := toLowerFirst(field.Name)
		// Если тип поля является структурой, рекурсивно вызываем функцию structToMap
		if fieldValue.Kind() == reflect.Struct {
			nestedData := StructToMap(fieldValue.Interface())
			data[fieldName] = nestedData
		} else if fieldValue.Kind() == reflect.Slice {
			var arrData []interface{}
			for j := 0; j < fieldValue.Len(); j++ {
				d := fieldValue.Index(j).Interface()
				nestedData := StructToMap(d)
				arrData = append(arrData, nestedData)
			}
			data[fieldName] = arrData
		} else {
			if fieldValue.Kind() == reflect.String && fieldValue.Interface().(string) == "" {
				continue
			}
			data[fieldName] = fieldValue.Interface()
		}
	}

	return data
}
