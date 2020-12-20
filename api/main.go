package main

import "log"

func main() {
	if err := migrateAllTables(); err != nil {
		log.Fatal(err)
	}
	setRouter()
	return
}
