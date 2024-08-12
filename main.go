package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

var jsonToParse = ``

var getValueBetween = regexp.MustCompile(`(?:\")(.*?)(?:\")`)

type StructBuild struct {
	Name   string
	Fields []StructValues
}

type StructValues struct {
	Field string
	Type  string
	Json  string
}

func main() {

	file, _ := os.ReadFile("jsonToParse.json")

	jsonMap := make(map[string]interface{})
	json.Unmarshal(file, &jsonMap)

	var structBuilders StructBuild
	var strs []StructBuild
	DefineValuesType(jsonMap, structBuilders, &strs)

	finalResult := BuildStruct(strs)
	fmt.Println(finalResult)

}

func DefineValuesType(jsonMap map[string]interface{}, structBuilder StructBuild, strs *[]StructBuild) {
	if structBuilder.Name == "" {
		mainStruct := StructBuild{
			Name: "Root",
		}

		structBuilder = mainStruct
	}

	for key, value := range jsonMap {
		if value != nil && reflect.TypeOf(value).Kind() == reflect.Slice {
			v := value.([]interface{})[0]

			if reflect.TypeOf(v).Kind() == reflect.Map {
				childStruct := StructBuild{
					Name: strings.Title(key),
				}
				m := v.(map[string]interface{})
				DefineValuesType(m, childStruct, strs)
			}

		} else if value != nil && reflect.TypeOf(value).Kind() == reflect.Map {
			childStruct := StructBuild{
				Name: strings.Title(key),
			}
			m := value.(map[string]interface{})
			DefineValuesType(m, childStruct, strs)
		}

		var structValues StructValues

		if value == nil {
			structValues.Type = "interface{}"
		} else {
			switch reflect.TypeOf(value).Kind().String() {
			case "map":
				structValues.Type = strings.Title(key)
			default:
				structValues.Type = reflect.TypeOf(value).Kind().String()
			}
		}
		structValues.Field = strings.Title(key)
		structValues.Json = key

		structBuilder.Fields = append(structBuilder.Fields, structValues)
	}

	*strs = append(*strs, structBuilder)

}

func BuildStruct(structs []StructBuild) string {

	builder := strings.Builder{}

	for _, str := range structs {
		builder.WriteString(fmt.Sprintf("type %s struct { \n", str.Name))
		for _, field := range str.Fields {
			if field.Type == "slice" {
				builder.WriteString(fmt.Sprintf("%s []%s `json:\"%s\"` \n", field.Field, field.Field, field.Json))
				continue
			}
			builder.WriteString(fmt.Sprintf("%s %s `json:\"%s\"` \n", field.Field, field.Type, field.Json))
		}
		builder.WriteString("} \n")
		builder.WriteString("\n")
		builder.WriteString("\n")
	}

	return builder.String()
}
