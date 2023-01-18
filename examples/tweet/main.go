package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent"
	"github.com/xlzhangkeke/entdemo/internal/config"
)

func main() {
	config.Init()
	client, err := ent.Open("mysql", config.MySqlDataSource)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Debug().Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
