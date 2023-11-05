package main

import (
"context"	
 "fmt"
 "os"

 "github.com/jackc/pgx/v5"
)

 type Responder interface {
 	Response() error
 }

// sqlCMD is only option I provide
// The rest I ask for the result back
 type Option struct {
 	sqlCMD string
 	id, age int
 	fname, lname string
 	optionId int // number of fields I expect back
 }



func (opt Option) Response(conn *pgx.Conn) error {
	qry := opt.sqlCMD
	fmt.Printf("SQLCMD:\n%s\n", qry)
	rows, _ := conn.Query(context.Background(), qry)
	var err error 
	
	switch opt.optionId {
	 case 1:
	  fmt.Printf("OptionId 1 was given\n")	
	  _, err = pgx.ForEachRow(rows, []any{&opt.fname, &opt.lname}, func() error {
		  fmt.Printf("fname: %v, lname: %v\n", opt.fname,opt.lname)
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
	  case 3:
	  _, err = pgx.ForEachRow(rows, []any{&opt.id, &opt.fname, &opt.lname, &opt.age}, func() error {
		  fmt.Printf("id: %v, fname: %v, lname: %v, age:%v\n", opt.id,opt.fname,opt.lname,opt.age)
		  return nil
	  })
	  case 4:
	  _, err = pgx.ForEachRow(rows, []any{&opt.id, &opt.fname, &opt.lname, &opt.age}, func() error {
		  fmt.Printf("id: %v, fname: %v, lname: %v, age:%v\n", opt.id,opt.fname,opt.lname,opt.age)
		  return nil
	  })
	  case 5:
	  _, err = pgx.ForEachRow(rows, []any{&opt.id, &opt.fname, &opt.lname, &opt.age}, func() error {
		  fmt.Printf("id: %v, fname: %v, lname: %v, age:%v\n", opt.id,opt.fname,opt.lname,opt.age)
		  return nil
	  })					
	}

	if err != nil {
		fmt.Printf("ForEachRow error: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func createOption(qry string) Option {
	 return Option {
		qry,0, 0,"","",0,
	}
}

func main() {
	
    uri := "postgres://user:secret@ipaddr:5432/databasename"
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
	

	
	switch argv {
	case 1:
		sqlCMD := "SELECT firstname,lastname FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 1 
		opt.Response(conn)
    case 2:
    	sqlCMD := "SELECT firstname,lastname, age FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 2 
		opt.Response(conn)
	case 3:
		sqlCMD := "SELECT * FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 3 
		opt.Response(conn)
	case 4:
		sqlCMD := "SELECT * FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 4 
		opt.Response(conn)	
	case 5:
		sqlCMD := "SELECT * FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 5 
		opt.Response(conn)
	default:
		sqlCMD := "SELECT * FROM People"
		opt := createOption(sqlCMD)
		opt.optionId = 1 
		opt.Response(conn)	
	}
}