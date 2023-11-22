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
	Lname string `db:"lastname"`
	ID int       `db:"id"`
	Age int      `db:"age"`
}

type Option struct {
	ch int // numbered choice for query
	People 
}

//helper funct
func (op Option) String() string{
	choice:= op.ch
	
	var str string
	switch choice {
	  case 1:
	  	str = fmt.Sprintf("%v %v", op.Fname, op.Lname)
	  case 2:
	  	str = fmt.Sprintf("%v %v, %d", op.Fname, op.Lname, op.Age)	
	  case 3:
	  	str = fmt.Sprintf("%v %v, %d", op.Fname, op.Lname, op.Age)	
	  case 4:
	  	str = fmt.Sprintf("%v", op.Fname)	
	  case 5:
	  	str = fmt.Sprintf("%v", op.Fname)	
	}

	return str
}

func loopFields (ch int, users[]*People, err error) error {
  if err != nil {
	  fmt.Fprintf(os.Stderr, "Error Querying Row %v\n", err)
	  os.Exit(1)
  }
  
  for _, person := range users {
  	opt := Option{ch, *person}
	fmt.Println(opt)
  }
  return nil 
}

func query(choice int, ctx context.Context, 
	db  *pgxpool.Pool, users []*People,) error {

	var err error
	switch choice {
	 case 1:
	 	err = pgxscan.Select(
		ctx, db, &users, "SELECT firstname, lastname FROM People",)
	case 2:
	 	err = pgxscan.Select(
		ctx, db, &users, "SELECT firstname, lastname, age FROM People",)
	case 3:
	 	err = pgxscan.Select(
		ctx, db, &users, "SELECT firstname, lastname, age FROM People\n" +
		"WHERE firstname IN ('Dave', 'Yennifer')",)
	case 4:
	 	err = pgxscan.Select(
		ctx, db, &users, "SELECT firstname FROM People",)
	case 5:
	 	err = pgxscan.Select(
		ctx, db, &users, "SELECT firstname FROM People",)
	}
	
	return loopFields(choice, users, err)
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

    var choice int
	fmt.Println("Select a query to ask the People's database given the following Options")
	fmt.Println(
		"Type '1' to get each user's name\n" +
		"Type '2' to get each user's name and age\n" +
		"Type '3' Compare the age between 2 users\n" +
		"Type '4' to get the name of each user and their id\n" +
		"Type '5' to get each user's last name and age",
	)
	fmt.Scanf("%d\n", &choice)

	var users []*People
	err = query(choice, ctx, db, users)
	if err != nil {
	  fmt.Fprintf(os.Stderr, "Error Querying Row %v\n", err)
	  os.Exit(1)
	}
}