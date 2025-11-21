package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"go-expenses/manager"
	"go-expenses/model"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	expenseManager := manager.ExpensesManager{}
	expenseManager.LoadFile("data/expenses.csv")

	for {
		fmt.Println("\nPersonal Expense Tracker")
		fmt.Println("========================")
		fmt.Println("1. Expenses")
		fmt.Println("2. Save and Exit")
		fmt.Println("3. Exit without Saving")
		fmt.Print("\nEnter your choice: ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		choice, _ := strconv.Atoi(line)

		switch choice {

		case 1:
			fmt.Println("\nExpenses Menu")
			fmt.Println("1. Add Expenses")
			fmt.Println("2. View Expenses")
			fmt.Println("3. Update Expenses")
			fmt.Println("4. Delete Expense")
			fmt.Println("5. Back")
			fmt.Print("\nEnter your choice: ")

			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			choice, _ = strconv.Atoi(line)

			switch choice {

			case 1:
				var date string
				var amount float64
				var categoryChoice int

				fmt.Print("Enter Date (MM-DD-YYYY): ")
				date, _ = reader.ReadString('\n')
				date = strings.TrimSpace(date)

				fmt.Print("Select Category (1. Needs | 2. Wants): ")

				catLine, _ := reader.ReadString('\n')
				catLine = strings.TrimSpace(catLine)
				categoryChoice, _ = strconv.Atoi(catLine)

				category := model.Needs
				if categoryChoice == 2 {
					category = model.Wants
				}

				fmt.Print("Enter Amount: ")

				amtLine, _ := reader.ReadString('\n')
				amtLine = strings.TrimSpace(amtLine)
				val, _ := strconv.ParseFloat(amtLine, 64)
				amount = math.Round(val*100) / 100 // round to 2 decimals

				fmt.Print("Enter Description (optional): ")
				desc, _ := reader.ReadString('\n')
				desc = strings.TrimSpace(desc)

				var descPtr *string
				if desc != "" {
					descPtr = &desc
				}

				fmt.Println("\nAdding Expense...")

				if err := expenseManager.CreateExpense(date, category, amount, descPtr); err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Expense added successfully!")
				}

			case 2:
				fmt.Println("\nAll Expenses:")
				expenseManager.ReadExpenses()

			case 3:
				if err := expenseManager.UpdateExpense(); err != nil {
					fmt.Println("Error:", err)
				}

			case 4:
				if err := expenseManager.DeleteExpense(); err != nil {
					fmt.Println("Error:", err)
				}

			case 5:
				continue

			default:
				fmt.Println("Invalid choice, try again.")
			}

		// save & exit
		case 2:
			expenseManager.SaveFile("data/expenses.csv")
			fmt.Println("\nFile saved. Exiting...")
			return

		// exit w/o save
		case 3:
			fmt.Println("\nExiting without saving.")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
