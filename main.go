package main

import (
	"log"
	"path/filepath"

	"github.com/hauntarl/task-manager/cmd"
	"github.com/hauntarl/task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	// This allows us to create our database at user's home location
	// irrespective of their OS, which unables the user to access the same db
	// regardless of the location of their executable.
	//
	// If we do not generate a home directory path and simply pass the db name
	// as path then the db is created at the same location as executable, then
	// if we go on and move our executable, we lose all the data: after running
	// the script, the code will create another db at it's current location.
	home, _ := homedir.Dir()
	path := filepath.Join(home, "todos.db")

	// initialize database
	if err := db.Init(path); err != nil {
		log.Fatal(err)
	}
	cmd.Execute() // start app
}
