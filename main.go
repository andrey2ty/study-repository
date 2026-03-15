package main

import (
	"fmt"
	"study/feature"
	"study/feature2"
	"study/postgres"
)

func main() {
	fmt.Println("hello git")
	feature.Feature()
	feature2.Feature2()
	postgres.ConnetctionToPostgres()
}
