package main

import (
	"./structs"
	"fmt"
	"os"
)

func cli_start() {
	for {
		var choisen string
		fmt.Println("command: ")
		choisen = structs.ReadLine()

		switch choisen {
		default:
			fallthrough
		case "exit":
			return
		case "add user":
			structs.AddUserToMap()
		case "add inventory":
			structs.AddInventoryToMap()
		case "show users":
			structs.ShowUsers()
		case "show inventory":
			structs.ShowInventory()
		case "delete user":
			structs.DeleteUser()
		case "delete inventory":
			structs.DeleteInventory()
		case "create assign":
			structs.CreateAssign_()
		case "help":
			structs.GetInfAboutCommands()
		}

	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	structs.InicializeMaps()

	switch len(argsWithoutProg) {
	case 1:
		if argsWithoutProg[0] == "help" {
			structs.GetInfAboutCommands()
		}
	case 2:
		if argsWithoutProg[1] == "user" {
			structs.AddUserToMap()
		} else if argsWithoutProg[1] == "inventory" {
			structs.AddInventoryToMap()
		}
	}

	cli_start()
}
