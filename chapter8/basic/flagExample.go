package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your wonderful name")
var age = flag.Int("age", 0, "our graceful age")
var surname string

func main() {
	flag.Parse()
	log.Printf("Hello %s %s (%d years), Welcome to the command line world", *name, surname, *age)
}

func init() {
	flag.StringVar(&surname, "surname", "none", "Your wonderful surname")
}
