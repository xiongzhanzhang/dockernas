package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("test", errors.New("saas"))
}
