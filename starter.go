package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/xlzhangkeke/entdemo/ent"
	"github.com/xlzhangkeke/entdemo/ent/car"
	"github.com/xlzhangkeke/entdemo/ent/group"
	"github.com/xlzhangkeke/entdemo/ent/user"
)

func main() {
	client, err := ent.Open("mysql", "root:pandora@tcp(10.95.84.100:33306)/entdemo?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	client = client.Debug()
	ctx := context.Background()
	//user, err := QueryUser(ctx, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(user)
	//err = QueryCarUsers(ctx, user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = CreateGraph(ctx, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = QueryGithub(ctx, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = QueryArielCars(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(30).SetName("a8m").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// `Only` fails if no user found, or more than 1 user returned.
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.Create().SetAge(30).SetName("a8m").AddCars(tesla, ford).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().Where(car.Model("Ford")).Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// Query the inverse edge.
	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.Create().SetAge(30).SetName("Ariel").Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users created above.
	// Attach this car to Ariel.

	err = client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).SetOwner(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	// Attach this car to Ariel.
	err = client.Car.Create().SetModel("Mazda").SetRegisteredAt(time.Now()).SetOwner(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	// Attach this graph to Neta.
	err = client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).SetOwner(neta).Exec(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	err = client.Group.Create().SetName("GitLab").AddUsers(neta, a8m).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.Create().SetName("GitHub").AddUsers(a8m).Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// Get "Ariel" from previous steps.
	a8m := client.User.Query().Where(
		user.HasCars(),
		user.Name("Ariel"),
	).OnlyX(ctx)
	// Get the groups, that a8m is connected to:
	cars, err := a8m.QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
					QueryUsers(). // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
					QueryCars().Where(
		car.Not( //  Get Neta and Ariel cars, but filter out
			car.Model("Mazda"), //  those who named "Mazda"
		),
	).All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}
