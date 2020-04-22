package main

import "easy-account/utils"

func main() {
	details := utils.NewDetails()

	// Define a variable to control when the for-loop is terminated
	loop := true

	for loop {
		key := utils.ShowOpening()
		switch key {
		case "1":
			details.Show()
		case "2":
			details.AddDetail(true)
		case "3":
			details.AddDetail(false)
		case "4":
			loop = false
		}
	}
}
