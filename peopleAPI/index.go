package main

import (
"context"	
 "fmt"
 "os"

"github.com/jackc/pgx/v5/pgxpool"
"github.com/georgysavva/scany/v2/pgxscan"
)

type People struct {
	Fname string `db:"firstname"`
	Lname string
	ID int
	Age int
}

func main() {
	
	uri:= "postgres://username:secret@ipAddr:5432/PeopleDB"
	os.Setenv("DATABASE_URL", uri)
	ctx:= context.Background()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to a database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()


    // var choice int
// 	fmt.Println("Select a query to ask the People's database given the following Options")
// 	fmt.Println(
// 		"Type '1' to get each user's name\n" +
// 		"Type '2' to get each user's name and age\n" +
// 		"Type '3' Compare the age between 2 users\n" +
// 		"Type '4' to get the name of each user and their id\n" +
// 		"Type '5' to get each user's last name and age",
// 	)
// 	fmt.Scanf("%d\n", &choice)

	
	var users []*People
	if err := pgxscan.Select(
		ctx, db, &users, "SELECT firstname FROM People",)
	err != nil {
		fmt.Fprintf(os.Stderr, "Error Querying Row %v\n", err)
		os.Exit(1)
	}

	for _, person := range users {
		fmt.Printf("%v\n", person.Fname)
	}
}