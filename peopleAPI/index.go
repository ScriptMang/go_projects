package main

import (
"context"	
 "fmt"
 "os"

 "github.com/jackc/pgx/v5"
)

func main() {
	
	
    uri := "postgres://user:secret@ipaddr:5432/databasename"
	os.Setenv("DATABASE_URL", uri)
    
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to a database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	
	sqlCMD := "SELECT * FROM People"
	
	rows, _ := conn.Query(context.Background(), sqlCMD)
	

    var  fname, lname string
	var id, age int64
	_, err = pgx.ForEachRow(rows, []any{&id, &fname, &lname, &age},  func() error {
		fmt.Printf("id: %v, fname: %v, lname: %v, age:%v\n", id,fname,lname,age)
		return nil
	})
	
	if err != nil {
		fmt.Printf("ForEachRow error: %v\n", err)
		os.Exit(1)
	}

}