package main

import (
"context"	
 "fmt"
 "os"

 "github.com/jackc/pgx/v5"
)

 type Responder interface {
 	Response(conn *pgx.Conn) error
 }

type Fields struct {
	fname, lname string
	id, age int
}

type Option struct {
 	qry string
 	optionID int
 	Fields
}

func createOption (qry string, id int) Option {
	return Option{qry, id, Fields{"","",0,0}}
}

func verifyForEachRowError(err error) error {
     if err != nil {
		fmt.Printf("ForEachRow error: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func (opt Option) Response(conn *pgx.Conn) error {
	qry := opt.qry
	rows, _ := conn.Query(context.Background(), qry)
	var err error 
	
	switch opt.optionID {
		case 1:
			_, err = pgx.ForEachRow(rows, []any{&opt.fname, &opt.lname}, func() error {
			  fmt.Printf("fname: %v, lname: %v\n", opt.fname, opt.lname)
			  return nil
		    })
		case 2:
		    _, err = pgx.ForEachRow(rows,[]any{
			  	&opt.fname, 
			  	&opt.lname, 
			  	&opt.age,}, func() error {
				fmt.Printf("fname: %v, lname: %v, age:%v\n", opt.fname,opt.lname,opt.age)
				return nil
			})
		default:
			_, err = pgx.ForEachRow(rows, []any{&opt.fname, &opt.lname}, func() error {
			  fmt.Printf("fname: %v, lname: %v\n", opt.fname, opt.lname)
			  return nil
		    })
	}
	
	return verifyForEachRowError(err)
}

func main() {
	
	uri:= "postgres://andyperalta@localhost/PeopleDB"
	os.Setenv("DATABASE_URL", uri)
  
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to a database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())


    var argv int
	fmt.Println("Select a query to ask the People's database given the following Options")
	fmt.Println(
		"Type '1' to get the each user's name\n" +
		"Type '2' to get each user's name and age\n" +
		"Type '3' Compare the age between 2 users\n" +
		"Type '4' to get the each user's id\n" +
		"Type '5' to get the each user's last name and age\n",
	)
	fmt.Scanf("%d\n", &argv)
	
	var rqst Responder
	switch argv {
	case 1:
		sqlCMD := "SELECT firstname,lastname FROM People"
		rqst = createOption(sqlCMD, 1)
    case 2:
    	sqlCMD := "SELECT firstname,lastname, age FROM People"
    	rqst = createOption(sqlCMD, 2)
	default:
		sqlCMD := "SELECT firstname,lastname FROM People"
		rqst = createOption(sqlCMD, 1)
	}
	rqst.Response(conn)
}