/*
TO OPEN ME PLEASE EXECUTE THE FOLLOWING:

$ go run paraEdu.go
*/


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var alwaysHere = true
	scanner := bufio.NewScanner(os.Stdin)

	var initial = "\nMe pongo canciones para escribirte.\n"
	fmt.Println(initial)
	fit := "A veces las cosas no encajan ni a golpes,\notras nacen rimando pa' siempre."
	fall := "Todo somos humanos\nque ante los problemas\nsolemos fallar.\nNo tengas miedo de caer al suelo\nque yo pienso estar al final."

	needMore := "Que todo sea saltar contigo al desastre,\nsólo para ver que hay detrás.\n"

	if(alwaysHere){
		fmt.Println(fit + "\n\n" + fall)

	}
	fmt.Println("\n\n******************************\nDo you want an extra of love?\nUsage:\n\n[y]es , [n]o or exit\n******************************")

	for scanner.Scan() {
		line := scanner.Text()
		if(line == "yes" || line == "y") {
			fmt.Println(needMore)
			os.Exit(0)
		} else if (line == "no" || line == "n") {
			fmt.Println("LOVE YOU, BYE :)\n")
			os.Exit(0)
		} else if line == "exit" {
			os.Exit(0)
		} else {
			fmt.Println("C'mon you can do it better...Try again!\n")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}