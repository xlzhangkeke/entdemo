package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xlzhangkeke/entdemo/examples/o2mrecur/ent"
	"github.com/xlzhangkeke/entdemo/examples/o2mrecur/ent/node"
)

func Example_O2MRecur() {
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
	// Tree leafs [1 3 5]
	// [1 3 5]
	// Node(id=1, value=2, parent_id=0)
}

func Do(ctx context.Context, client *ent.Client) error {
	root, err := client.Node.
		Create().
		SetValue(2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the root: %w", err)
	}

	// Add additional nodes to the tree:
	//
	//       2
	//     /   \
	//    1     4
	//        /   \
	//       3     5
	//

	// Unlike `Save`, `SaveX` panics if an error occurs.
	n1 := client.Node.
		Create().
		SetValue(1).
		SetParent(root).
		SaveX(ctx)
	n4 := client.Node.
		Create().
		SetValue(4).
		SetParent(root).
		SaveX(ctx)
	n3 := client.Node.
		Create().
		SetValue(3).
		SetParent(n4).
		SaveX(ctx)
	n5 := client.Node.
		Create().
		SetValue(5).
		SetParent(n4).
		SaveX(ctx)

	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
	// Output: Tree leafs [1 3 5]

	// Get all leafs (nodes without children).
	// Unlike `Int`, `IntX` panics if an error occurs.
	ints := client.Node.
		Query(). // All nodes.
		Where(node.Not(node.HasChildren())). // Only leafs.
		Order(ent.Asc(node.FieldValue)). // Order by their `value` field.
		GroupBy(node.FieldValue). // Extract only the `value` field.
		IntsX(ctx)
	fmt.Println(ints)
	// Output: [1 3 5]

	// Get orphan nodes (nodes without parent).
	// Unlike `Only`, `OnlyX` panics if an error occurs.
	orphan := client.Node.
		Query().
		Where(node.Not(node.HasParent())).
		OnlyX(ctx)
	fmt.Println(orphan)
	// Output: Node(id=1, value=2, parent_id=0)

	return nil
}
