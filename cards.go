package main

import "fmt"
import "./game"

func main() {

	output := game.TakeTurn()
	fmt.Println(output)

	//fmt.Printf("[?][ ]   [ ][ ][ ][ ]\n")
	//fmt.Printf("\n")
	//	fmt.Printf("[?][?][?][?][?][?][?]\n")
	//	fmt.Printf("   [?][?][?][?][?][?]\n")
	//	fmt.Printf("      [?][?][?][?][?]\n")
	//	fmt.Printf("         [?][?][?][?]\n")
	//	fmt.Printf("            [?][?][?]\n")
	//	fmt.Printf("               [?][?]\n")
	//	fmt.Printf("                  [?]\n")
}
