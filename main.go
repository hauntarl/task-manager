package main

import (
	"log"
	"path/filepath"

	"github.com/hauntarl/task-manager/cmd"
	"github.com/hauntarl/task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	path := filepath.Join(home, "todos.db")
	if err := db.Init(path); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
