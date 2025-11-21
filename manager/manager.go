package manager

import (
	"bufio"
	"errors"
	"fmt"
	"go-expenses/model"
	"go-expenses/utils"
	"os"
	"strconv"
	"strings"
)

type ExpensesManager struct {
	Expenses []model.Expense
}

func (m *ExpensesManager) CreateExpense(
	date string,
	category model.CategoryType,
	amount float64,
	description *string,
) error {

	if amount <= 0 {
		return errors.New("amount cannot be 0 or negative")
	}

	if date == "" {
		return errors.New("date cannot be empty")
	}

	if !utils.IsValidDate(date) {
		return errors.New("invalid date format. expected MM-DD-YYYY")
	}

	newExpense := model.Expense{
		Date:        date,
		Category:    category,
		Amount:      amount,
		Description: description,
	}

	m.Expenses = append(m.Expenses, newExpense)
	return nil
}

func (m *ExpensesManager) ReadExpenses() {
	for _, expense := range m.Expenses {
		fmt.Printf("Date: %s, Category: %s, Amount: %.2f",
			expense.Date,
			expense.Category.String(),
			expense.Amount,
		)

		if expense.Description != nil {
			fmt.Printf(", Description: %s", *expense.Description)
		}

		fmt.Println()
	}
}

func (m *ExpensesManager) DeleteExpense() error {
	reader := bufio.NewReader(os.Stdin)

	// print list
	fmt.Println("=== All Expenses ===")
	if len(m.Expenses) == 0 {
		fmt.Println("No expenses available.")
		return nil
	}

	for i, exp := range m.Expenses {
		fmt.Printf("%d. Date: %s, Category: %s, Amount: %.2f",
			i+1,
			exp.Date,
			exp.Category.String(),
			exp.Amount,
		)
		if exp.Description != nil {
			fmt.Printf(", Description: %s", *exp.Description)
		}
		fmt.Println()
	}
	fmt.Println()

	// get user input
	fmt.Print("Enter Date of Expense to Delete (MM-DD-YYYY): ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	if !utils.IsValidDate(date) {
		return errors.New("invalid date format. expected MM-DD-YYYY")
	}

	// filter list by date
	filteredIndexes := []int{}
	for i, exp := range m.Expenses {
		if exp.Date == date {
			filteredIndexes = append(filteredIndexes, i)
		}
	}

	if len(filteredIndexes) == 0 {
		return errors.New("no expenses found for the given date")
	}

	// print filtered list
	fmt.Printf("\nExpenses on %s:\n", date)
	for i, idx := range filteredIndexes {
		exp := m.Expenses[idx]
		fmt.Printf("%d. Date: %s, Category: %s, Amount: %.2f",
			i+1,
			exp.Date,
			exp.Category.String(),
			exp.Amount,
		)
		if exp.Description != nil {
			fmt.Printf(", Description: %s", *exp.Description)
		}
		fmt.Println()
	}

	var choice int
	for {
		fmt.Print("Select expense number to delete: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err == nil && num >= 1 && num <= len(filteredIndexes) {
			choice = num
			break
		}
		fmt.Println("Invalid number. Try again.")
	}

	deleteIndex := filteredIndexes[choice-1]
	m.Expenses = append(m.Expenses[:deleteIndex], m.Expenses[deleteIndex+1:]...)

	fmt.Println("Expense deleted successfully.")
	return nil
}

func (m *ExpensesManager) UpdateExpense() error {
	reader := bufio.NewReader(os.Stdin)

	// print list expenses
	fmt.Println("=== All Expenses ===")
	if len(m.Expenses) == 0 {
		fmt.Println("No expenses available.")
		return nil
	}

	for i, exp := range m.Expenses {
		fmt.Printf("%d. Date: %s, Category: %s, Amount: %.2f",
			i+1,
			exp.Date,
			exp.Category.String(),
			exp.Amount,
		)
		if exp.Description != nil {
			fmt.Printf(", Description: %s", *exp.Description)
		}
		fmt.Println()
	}
	fmt.Println()

	// date user input
	fmt.Print("Enter Date of Expense to Update (MM-DD-YYYY): ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	if !utils.IsValidDate(date) {
		return errors.New("invalid date format. expected MM-DD-YYYY")
	}

	// filter list by date
	filteredIndexes := []int{}
	for i, exp := range m.Expenses {
		if exp.Date == date {
			filteredIndexes = append(filteredIndexes, i)
		}
	}

	if len(filteredIndexes) == 0 {
		return errors.New("no expenses found for the given date")
	}

	// print filtered list
	fmt.Printf("Expenses on %s:\n", date)
	for i, idx := range filteredIndexes {
		exp := m.Expenses[idx]
		fmt.Printf("%d. Date: %s, Category: %s, Amount: %.2f",
			i+1,
			exp.Date,
			exp.Category.String(),
			exp.Amount,
		)
		if exp.Description != nil {
			fmt.Printf(", Description: %s", *exp.Description)
		}
		fmt.Println()
	}

	var choice int
	for {
		fmt.Print("Select expense number to update: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err == nil && num >= 1 && num <= len(filteredIndexes) {
			choice = num
			break
		}
		fmt.Println("Invalid number. Try again.")
	}

	idx := filteredIndexes[choice-1]
	exp := &m.Expenses[idx]

	// update values

	fmt.Printf("New Date (current: %s, press Enter to keep): ", exp.Date)
	newDate, _ := reader.ReadString('\n')
	newDate = strings.TrimSpace(newDate)
	if newDate != "" {
		if !utils.IsValidDate(newDate) {
			return errors.New("invalid new date format")
		}
		exp.Date = newDate
	}

	fmt.Printf("New Category (current: %s, 1=Needs, 2=Wants, Enter to keep): ",
		exp.Category.String())
	categoryLine, _ := reader.ReadString('\n')
	categoryLine = strings.TrimSpace(categoryLine)
	if categoryLine != "" {
		switch categoryLine {
		case "1":
			exp.Category = model.Needs
		case "2":
			exp.Category = model.Wants
		default:
			return errors.New("invalid category input")
		}
	}

	fmt.Printf("New Amount (current: %.2f, press Enter to keep): ", exp.Amount)
	amountLine, _ := reader.ReadString('\n')
	amountLine = strings.TrimSpace(amountLine)
	if amountLine != "" {
		newAmount, err := strconv.ParseFloat(amountLine, 64)
		if err != nil || newAmount <= 0 {
			return errors.New("invalid new amount")
		}
		exp.Amount = newAmount
	}

	fmt.Printf("New Description (current: %v, press Enter to keep): ",
		func() string {
			if exp.Description == nil {
				return ""
			}
			return *exp.Description
		}())
	descLine, _ := reader.ReadString('\n')
	descLine = strings.TrimSpace(descLine)
	if descLine != "" {
		exp.Description = &descLine
	}

	fmt.Println("Expense updated successfully.")
	return nil
}

func (m *ExpensesManager) SaveFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	for _, expense := range m.Expenses {
		_, err := file.WriteString(expense.ToCSV() + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ExpensesManager) LoadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close() //

	m.Expenses = []model.Expense{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		expense, err := model.ExpenseFromCSV(line)
		if err != nil {
			fmt.Printf("Error parsing line: %s (%v)\n", line, err)
			continue
		}

		m.Expenses = append(m.Expenses, expense)
	}

	return nil
}
