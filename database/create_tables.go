package database

import (
	"context"
	"fmt"
	"log"

	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"
)

func createTable(modelConfig models.ModelConfig) {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS "%s" (
			%s
		)
	`, modelConfig.TableName, modelConfig.CreateTableConfig)
	_, err := Db.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error creating %s table: %v\nQuery: %s", modelConfig.TableName, err, query)
	}
	fmt.Printf("Create %s table successfully\n", modelConfig.TableName)
}

func CreateTables() {
	fmt.Println("Create tables")
	modelConfigs := models.GenerateModelConfigs()

	for _, c := range modelConfigs {
		createTable(c)
	}
}
