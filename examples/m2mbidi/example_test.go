package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xlzhangkeke/entdemo/examples/m2mbidi/ent"
	"github.com/xlzhangkeke/entdemo/examples/m2mbidi/ent/user"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func Example_M2MBidi() {
	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("mysql", "root:pandora@tcp(10.95.84.100:33306)/ent_examples?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	client = client.Debug()
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// [User(id=1, age=30, name=a8m)]
	// [User(id=2, age=28, name=nati)]
	// [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
}

func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SaveX(ctx)
	nati := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		AddFriends(a8m).
		SaveX(ctx)

	// Query friends. Unlike `All`, `AllX` panics if an error occurs.
	friends := nati.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m)]

	friends = a8m.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=2, age=28, name=nati)]

	// Query the graph:
	friends = client.User.
		Query().
		Where(user.HasFriends()).
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
