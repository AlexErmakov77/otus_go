package main

import (
	"flag"
	"log"
	"strings"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	// Place your code here.
	if len(strings.TrimSpace(from)) == 0 {
		log.Fatal("empty argument  -from")
	}
	if len(strings.TrimSpace(to)) == 0 {
		log.Fatal("empty argument  -to")
	}
	if limit < 0 {
		log.Fatal("enter limit pozitive !")
	}
	if offset < 0 {
		log.Fatal("enter offset pozitive !")
	}
	if err := Copy(from, to, offset, limit); err != nil {
		println(err)
	}
}
