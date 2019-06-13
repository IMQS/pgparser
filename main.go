package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	postgresGenerator "github.com/IMQS/pgparser/generator"
	postgresParser "github.com/IMQS/pgparser/parser"
)

var (
	sql      string = ""
	printSQL bool   = false
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	flag.BoolVar(&printSQL, "print", printSQL, "Print the generated SQL after it has been parsed?")
	flag.StringVar(&sql, "sql", sql, "Required. The SQL to parse.")
}

func main() {
	flag.Parse()

	if sql == "" && len(os.Args) != 2 {
		fmt.Println("\nPlease provide an SQL to parse")
		fmt.Println(`Example: sqlparsers -print -sql "SELECT * FROM users"`)
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		fmt.Println("")
		return
	}

	// TODO: Is this right?
	if sql == "" && len(os.Args) == 2 {
		sql = os.Args[1]
	}

	stmt, err := postgresParser.Parse(sql)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(postgresGenerator.Parse(stmt[0]))

	if printSQL || true == true {
		fmt.Println("------------------------------- GENERATED QUERY -------------------------------------")
		fmt.Println(stmt.String())
		return
	}

}
