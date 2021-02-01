package structs

import (
	"../cache"
	"fmt"
	"os"
	"strconv"
)

var inventoryNameToStruct = make(map[string]InventoryStruct)
var inventoryIntToStruct = make(map[int]InventoryStruct)
var usersNameToStruct = make(map[string]UserStruct)
var usersIdToStruct = make(map[int]UserStruct)

func ShowInventory() {
	for key, value := range inventoryNameToStruct {
		fmt.Println(key + ": " + value.descriptionOfInventory +
			" - id: " + strconv.Itoa(value.id) +
			" @" + value.assigned.userName)
	}
}

func ShowUsers() {
	for _, value := range usersNameToStruct {
		fmt.Println(value.userName + " - id: " + strconv.Itoa(value.userId))
	}
}

func AddInventoryToMap() {
	var nameInventory string
	var descriptionInventory string

	fmt.Println("name of inventory: ")
	nameInventory = ReadLine()

	fmt.Println("description of inventory: ")
	descriptionInventory = ReadLine()

	newInventory := InventoryStruct{
		nameOfInventory:        nameInventory,
		descriptionOfInventory: descriptionInventory,
		id:                     cache.Cache.GetLastInventoryId(),
	}

	inventoryNameToStruct[newInventory.nameOfInventory] = newInventory
	inventoryIntToStruct[newInventory.id] = newInventory
}

func AddUserToMap() {
	var userName string

	fmt.Println("user name: ")
	userName = ReadLine()

	newUser := UserStruct{
		userName: userName,
		userId:   cache.Cache.GetLastUserId(),
	}

	usersNameToStruct[newUser.userName] = newUser
	usersIdToStruct[newUser.userId] = newUser
}

func DeleteUser() {
	var userName string

	fmt.Println("user name to delete: ")
	userName = ReadLine()

	delete(usersIdToStruct, usersNameToStruct[userName].userId)
	delete(usersNameToStruct, userName)
}

func DeleteInventory() {
	var nameInventory string

	fmt.Println("name of inventory to delete: ")
	fmt.Fscan(os.Stdin, &nameInventory)

	delete(inventoryIntToStruct, inventoryNameToStruct[nameInventory].id)
	delete(inventoryNameToStruct, nameInventory)
}

func CreateAssign_() {
	var inventoryStr string
	var userStr string

	fmt.Println("input assign inventoryId/inventoryName to userId/username")
	fmt.Println("var of inventory: ")
	inventoryStr = ReadLine()

	fmt.Println("var of user: ")
	userStr = ReadLine()

	if err, value := stringToInt(inventoryStr); err == nil {
		if err2, value2 := stringToInt(userStr); err2 == nil {
			createAssign(value, value2)
		} else {
			createAssign(value, userStr)
		}
		return
	} else {
		if err2, value2 := stringToInt(userStr); err2 == nil {
			createAssign(inventoryStr, value2)
		} else {
			createAssign(inventoryStr, userStr)
		}
	}
}

func createAssign(inventory interface{}, user interface{}) {
	switch valueFir := inventory.(type) {
	case int:
		InvStruct := inventoryIntToStruct[valueFir]
		var UsStruct UserStruct

		switch valueSec := user.(type) {
		case int:
			UsStruct = usersIdToStruct[valueSec]
		case string:
			UsStruct = usersNameToStruct[valueSec]
		}
		InvStruct.assigned = UsStruct
		inventoryIntToStruct[valueFir] = InvStruct
		inventoryNameToStruct[InvStruct.nameOfInventory] = InvStruct
	case string:
		InvStruct := inventoryNameToStruct[valueFir]
		var UsStruct UserStruct

		switch valueSec := user.(type) {
		case int:
			UsStruct = usersIdToStruct[valueSec]
		case string:
			UsStruct = usersNameToStruct[valueSec]
		}
		InvStruct.assigned = UsStruct
		inventoryIntToStruct[InvStruct.id] = InvStruct
		inventoryNameToStruct[valueFir] = InvStruct
	}
}

func GetInfAboutCommands() {
	fmt.Println("--add user //adds user to map\n" +
		"--add inventory //adds inventory to map \n" +
		"--show __ //show existing __\n" +
		"--delete __ //deletes __ from map\n" +
		"--crete assign //added new information about existing obj")
}

type InventoryStruct struct {
	nameOfInventory        string
	descriptionOfInventory string
	id                     int
	assigned               UserStruct
}

type UserStruct struct {
	userName string
	userId   int
}
