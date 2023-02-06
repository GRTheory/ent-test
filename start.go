package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GRTheory/ent-test/ent"
	"github.com/GRTheory/ent-test/ent/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:pass_123@tcp(192.168.228.128:3306)/base?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating shema resources: %v", err)
	}

	// user, err := createUser(context.Background(), client)
	user, err := QueryUser(context.Background(), client)
	if err != nil {
		log.Fatalf("failed creating a new user table: %v", err)
	}
	log.Println("user", user)
}

func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
