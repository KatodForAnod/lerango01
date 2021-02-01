package structs

import (
	"bufio"
	"os"
	"strconv"
)

func stringToInt(str string) (error, int) {
	var i int
	i, err := strconv.Atoi(str)

	return err, i
}
func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func InicializeMaps() {
	invstruct1 := InventoryStruct{
		nameOfInventory:        "MacBook",
		descriptionOfInventory: "MacBook Pro 15, RAM 8GB, SDD 256GB",
		id:                     1,
	}

	invStruct2 := InventoryStruct{
		nameOfInventory:        "Book",
		descriptionOfInventory: "Go PROGRAMMING",
		id:                     2,
	}
	inventoryNameToStruct["MacBook"] = invstruct1
	inventoryNameToStruct["Book"] = invStruct2

	inventoryIntToStruct[1] = invstruct1
	inventoryIntToStruct[2] = invStruct2

	userStruct1 := UserStruct{
		userName: "Pete Anderson",
		userId:   1,
	}
	userStruct2 := UserStruct{
		userName: "Alisa Isterina",
		userId:   2,
	}
	usersNameToStruct["Pete Anderson"] = userStruct1
	usersNameToStruct["Alisa Isterina"] = userStruct2

	usersIdToStruct[1] = userStruct1
	usersIdToStruct[2] = userStruct2

}
