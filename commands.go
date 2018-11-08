package main

import "fmt"

type Command struct {
	Letter      string
	Name        string
	Description string
	Action      func()
}

var commands = []Command{
	{"a", "assemble", "assemble montage", assemble},
	{"l", "list", "list montages", list},
	{"n", "count", "count characters, bytes and words", count},
	{"c", "clean", "cleanup", clean},
	{"d", "dump", "dump configuration", dump},
}

func (command Command) String() string {
	return fmt.Sprintf("  %s  %-10.10s %s", command.Letter, command.Name, command.Description)
}

func dump() {
	fmt.Printf("%+v\n", configuration)
}
