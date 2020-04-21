package main

import (
	"fmt"
)

type accountDetail struct {
	inout  bool //income(true) or outcome(false)
	money  float32
	commit string //description of money
}

type Details []accountDetail

func (details *Details) sum () float32 {
	var result float32 = 0.0
	for _, obj := range *details{
		if obj.inout {
			result += obj.money
		} else {
			result -= obj.money
		}
	}
	return result
}

func main() {
	var details Details

	// Define a variable to control when the for-loop is terminated
	loop := true

	for loop {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("-----------------------------Account Software-----------------------------")
		fmt.Println("                             1. Show the details")
		fmt.Println("                             2. Add an income")
		fmt.Println("                             3. Add an outcome")
		fmt.Println("                             4. Exit")
		fmt.Println()
		fmt.Print("Select a num: ")
		var key string
		_, _ = fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("-----------------------------Account Details-----------------------------")
			fmt.Println("  \t money \t\t\t in/out \t\t description")
			if len(details) > 0 {
				for i, obj := range details {
					inout := "in"
					if !obj.inout {
						inout = "out"
					}
					fmt.Printf("#%v \t %v \t\t %v \t\t\t %v",i+1, obj.money, inout, obj.commit)
					fmt.Println()
				}
				fmt.Printf("Total: %v", details.sum())
			}
		case "2":
			var money float32
			var commit string
			fmt.Print("Add an income: ")
			_, _ = fmt.Scanln(&money)
			fmt.Print("Add a description: ")
			_, _ = fmt.Scanln(&commit)
			details = append(details, accountDetail{
				inout:  true,
				money:  money,
				commit: commit,
			})
		case "3":
			var money float32
			var commit string
			fmt.Print("Add an outcome: ")
			_, _ = fmt.Scanln(&money)
			fmt.Print("Add a description: ")
			_, _ = fmt.Scanln(&commit)
			details = append(details, accountDetail{
				inout:  false,
				money:  money,
				commit: commit,
			})
		case "4":
			loop = false
		}
	}
}
