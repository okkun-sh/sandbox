package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type User struct {
	ID   int64
	Name string
}

func main() {
	ctx := context.Background()
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Addr:   fmt.Sprintf("%s:%v", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		Net:    "tcp",
	}

	m, err := sql.Open("mysql", cfg.FormatDSN())
	defer m.Close()

	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(m, mysqldialect.New())
	name, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	u := &User{
		Name: name.String(),
	}
	_, err = db.NewInsert().Model(u).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
