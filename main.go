package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/IMQS/pgparser/generator"
	"github.com/IMQS/pgparser/parser"
)

func main() {
	s := flag.String("sql", "", "SQL Statement")
	f := flag.String("file", "", "File Containing SQL")
	p := flag.Bool("print", true, "Print SQL")
	g := flag.Bool("generate", false, "Run Generator")

	flag.Parse()

	if *s == "" && *f == "" {
		fmt.Println("\nPlease provide SQL to parse")
		fmt.Println("Either:\n pgparser -sql \"SELECT * FROM users\"")
		fmt.Println("Or:\n pgparser -file <filename containing sql>")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		fmt.Println("")
		return
	}

	var sql string
	if *s != "" {
		sql = *s
	} else {
		content, err := ioutil.ReadFile(*f)
		if err != nil {
			log.Fatal(err)
		}
		sql = string(content)
	}

	stmt, err := parser.Parse(sql)
	if err != nil {
		log.Println(err)
		return
	}
	if *g {
		fmt.Println(generator.Parse(stmt[0]))
	}

	if *p {
		fmt.Println("------------------------------- GENERATED QUERY -------------------------------------")
		fmt.Println(stmt.String())
		return
	}

}
