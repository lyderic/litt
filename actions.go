package main

import "fmt"

type Action struct {
	Letter      string
	Name        string
	Description string
	Function    func()
}

var actions = []Action{
	{"a", "assemble", "assemble montage", assemble},
	{"l", "list", "list montages", list},
	{"n", "count", "count characters, bytes and words", count},
	{"c", "clean", "cleanup", clean},
	{"d", "dump", "dump configuration", dump},
}

func (action Action) String() string {
	return fmt.Sprintf("  %s  %-10.10s %s", action.Letter, action.Name, action.Description)
}

func dump() {
	fmt.Printf("%+v\n", configuration)
}
