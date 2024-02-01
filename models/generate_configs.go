package models

import (
	"fmt"
	"reflect"
	"strings"
)

type ModelConfig struct {
	CreateTableConfig string
	TableName         string
	Fields            []string
}

var AllModels = []interface{}{
	User{},
}

var MapModelConfigs = map[string]ModelConfig{}

func GetTableName(v interface{}) string {
	s := reflect.ValueOf(v)
	typeOfV := s.Type()
	return strings.ToLower(typeOfV.Name())
}

func GetModelConfig(v interface{}) ModelConfig {
	var fields []string
	s := reflect.ValueOf(v)
	typeOfV := s.Type()

	for i := 0; i < s.NumField(); i++ {
		field := typeOfV.Field(i)
		fieldName := field.Tag.Get("db")
		fieldType := MapTypeToPostgres(field.Type.String())
		parsedConfig := ParseFieldConfig(field.Tag.Get("config"))

		if parsedConfig.Serial {
			fieldType = "serial"
		}
		if parsedConfig.Unique {
			fieldType += " unique"
		}
		if parsedConfig.NotNull {
			fieldType += " not null"
		}
		if parsedConfig.PrimaryKey {
			fieldType += " primary key"
		}
		if parsedConfig.IsDefault {
			fieldType += " default " + parsedConfig.DefaultVal
		}
		fields = append(fields, fmt.Sprintf("%s %s", fieldName, fieldType))
	}
	return ModelConfig{
		TableName:         GetTableName(v),
		CreateTableConfig: strings.Join(fields, ", "),
		Fields:            fields,
	}
}

func GenerateModelConfigs() []ModelConfig {
	res := []ModelConfig{}

	for _, v := range AllModels {
		modelConfig := GetModelConfig(v)
		res = append(res, modelConfig)
		MapModelConfigs[modelConfig.TableName] = modelConfig
	}
	return res
}
