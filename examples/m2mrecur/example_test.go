package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xlzhangkeke/entdemo/examples/m2mrecur/ent"
	"github.com/xlzhangkeke/entdemo/examples/m2mrecur/ent/user"
)

func Example_M2MRecur() {
	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("mysql", "root:pandora@tcp(10.95.84.100:33306)/ent_examples?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	client = client.Debug()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// [User(id=2, age=28, name=nati)]
	// []
	// []
	// [User(id=1, age=30, name=a8m)]
	// [28]
	// [a8m]
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
		AddFollowers(a8m).
		SaveX(ctx)

	// Query following/followers:

	flw := a8m.QueryFollowing().AllX(ctx)
	fmt.Println(flw)
	// Output: [User(id=2, age=28, name=nati)]

	flr := a8m.QueryFollowers().AllX(ctx)
	fmt.Println(flr)
	// Output: []

	flw = nati.QueryFollowing().AllX(ctx)
	fmt.Println(flw)
	// Output: []

	flr = nati.QueryFollowers().AllX(ctx)
	fmt.Println(flr)
	// Output: [User(id=1, age=30, name=a8m)]

	// Traverse the graph:

	ages := nati.
		QueryFollowers(). // [a8m]
		QueryFollowing(). // [nati]
		GroupBy(user.FieldAge). // [28]
		IntsX(ctx)
	fmt.Println(ages)
	// Output: [28]

	names := client.User.
		Query().
		Where(user.Not(user.HasFollowers())).
		GroupBy(user.FieldName).
		StringsX(ctx)
	fmt.Println(names)
	// Output: [a8m]
	return nil
}
