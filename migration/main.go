package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xlzhangkeke/entdemo/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:pandora@tcp(10.95.84.100:33306)/entdemo?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	client = client.Debug()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
