package main

import (
	"fmt"
	"strconv"
	dm "github.com/JerryXie98/PathFynder/Data_Models"
	db "github.com/JerryXie98/PathFynder/Database"
)

func main() {
	fmt.Printf("test\n")
	player := dm.Game_Object{1,1,true}

	fmt.Printf("Player position is: (" +
				strconv.Itoa(player.PosX) + ", " +
				strconv.Itoa(player.PosY) + ")\n")

	db.GetHighScore(69)
}