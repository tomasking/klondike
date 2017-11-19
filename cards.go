package main

import "fmt"

func main() {

	output := takeTurn()
	for a := 0; a < 9; a++ {
		fmt.Println(output[a])
	}
}

func takeTurn() [9]string {

	var a [9]string
	a[0] = "[?][ ]   [ ][ ][ ][ ]"
	a[1] = ""
	a[2] = "[?][?][?][?][?][?][?]"
	a[3] = "   [?][?][?][?][?][?]"
	a[4] = "      [?][?][?][?][?]"
	a[5] = "         [?][?][?][?]"
	a[6] = "            [?][?][?]"
	a[7] = "               [?][?]"
	a[8] = "                  [?]"
	return a
}
