package database

import (
	"context"
	"fmt"
)

func generateUserTable() {
	ctx := context.Background()
	var exists bool
	existsErr := Db.QueryRow(ctx, "select exists (select from pg_tables where tablename = 'users')").Scan(&exists)

	if existsErr != nil {
		fmt.Println("generateUserDb err", existsErr)
		return
	}
	if exists {
		return
	}
	_, createUserTableErr := Db.Exec(ctx, `create table users (
		id serial primary key,
		username varchar(50) unique not null,
		create_at timestamp default current_timestamp,
		full_name varchar(100)
	)`)

	if createUserTableErr != nil {
		fmt.Println("create users table err", createUserTableErr)
		return
	}
	fmt.Println("create users table successfully")
}

func GenerateTables() {
	fmt.Println("Generate tables")
	generateUserTable()
}
