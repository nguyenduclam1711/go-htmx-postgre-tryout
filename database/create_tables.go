package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"
)

func createTable(v interface{}) {
	var fields []string
	s := reflect.ValueOf(v)
	typeOfV := s.Type()

	for i := 0; i < s.NumField(); i++ {
		field := typeOfV.Field(i)
		fieldName := field.Tag.Get("db")
		fieldType := mapTypeToPostgres(field.Type.String())
		parsedConfig := parseFieldConfig(field.Tag.Get("config"))

		if parsedConfig.serial {
			fieldType = "serial"
		}
		if parsedConfig.unique {
			fieldType += " unique"
		}
		if parsedConfig.notNull {
			fieldType += " not null"
		}
		if parsedConfig.primaryKey {
			fieldType += " primary key"
		}
		if parsedConfig.isDefault {
			fieldType += " default " + parsedConfig.defaultVal
		}
		fields = append(fields, fmt.Sprintf("%s %s", fieldName, fieldType))
	}

	tableName := strings.ToLower(typeOfV.Name())
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS "%s" (
			%s
		)
	`, tableName, strings.Join(fields, ", "))

	_, err := Db.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error creating %s table: %v\nQuery: %s", tableName, err, query)
	}
	fmt.Printf("Create %s table successfully\n", tableName)
}

func generateUserTable() {
	createTable(models.User{})
}

func CreateTables() {
	fmt.Println("Create tables")
	generateUserTable()
}
