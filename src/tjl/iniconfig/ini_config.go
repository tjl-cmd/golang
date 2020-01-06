package iniconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MarshalFile(filename string, data interface{}) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, result, 0755)

}
func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	var conf []string
	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		sectionField := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)
		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionField.Name
		}
		section := fmt.Sprintf("\n[%s]\n\n", tagVal)
		conf = append(conf, section)
		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			filedTagVal := keyField.Tag.Get("ini")
			if len(filedTagVal) == 0 {
				filedTagVal = keyField.Name
			}
			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", filedTagVal, valField)
			conf = append(conf, item)
		}
	}
	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}
func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	return UnMarshal(data, result)
}
func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	//var m map[string]string
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return
	}
	var lastFiledName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//如果是注释，直接忽略掉
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			lastFiledName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
			}
			continue
		}
		err = parseItem(lastFiledName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno %d", err, index+1)
		}
	}
	return
}
func parseItem(lastFileName, line string, result interface{}) (err error) {
	Index := strings.Index(line, "=")
	if Index == -1 {
		err = fmt.Errorf("syntax error ,lines:%s", line)
		return
	}
	key := strings.TrimSpace(line[0:Index])
	val := strings.TrimSpace(line[Index+1:])
	if len(key) == 0 {
		err = fmt.Errorf("syntax error ,lines:%s", line)
		return
	}
	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFileName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field:%s must be struct", lastFileName)
		return
	}
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		intVla, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVla)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		intVla, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVla)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetFloat(floatVal)
	default:
		err = fmt.Errorf("unsupport type:%d", fieldValue.Type().Kind())
	}
	return
}
func parseSection(line string, typeInfo reflect.Type) (fieldName string, err error) {

	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error ,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error ,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error ,invalid section:%s", line)
			return
		}
		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			tagVal := field.Tag.Get("ini")
			if tagVal == sectionName {
				fieldName = field.Name

				break
			}
		}
	}
	return
}
