package model

import (
	"fmt"
	"strconv"
)

type Expense struct {
	Date        string
	Category    CategoryType
	Amount      float64
	Description *string
}

// convert object to csv
func (e Expense) ToCSV() string {

	descriptionVal := ""

	if e.Description != nil {
		descriptionVal = *e.Description
	}
	return fmt.Sprintf("%q, %q, %q, %q",
		e.Date,
		e.Category.String(),
		fmt.Sprintf("%.2f", e.Amount),
		descriptionVal)
}

// convert csv to object
func ExpenseFromCSV(line string) (Expense, error) {
	//get quote
	quotePositions := []int{}
	for i, c := range line {
		if c == '"' {
			quotePositions = append(quotePositions, i)
		}
	}

	if len(quotePositions) < 8 { // 4 fields
		return Expense{}, fmt.Errorf("invalid CSV format: %s", line)
	}

	// extract fields
	date := line[quotePositions[0]+1 : quotePositions[1]]
	categoryStr := line[quotePositions[2]+1 : quotePositions[3]]
	amountStr := line[quotePositions[4]+1 : quotePositions[5]]
	descriptionStr := line[quotePositions[6]+1 : quotePositions[7]]

	// conv string amt to float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return Expense{}, fmt.Errorf("invalid amount: %s", amountStr)
	}

	// convert str cat to CategoryType
	category, err := StringToCategory(categoryStr)
	if err != nil {
		return Expense{}, err
	}

	// conv optional desc
	var description *string
	if descriptionStr != "" {
		description = &descriptionStr
	}

	// ret converted obj
	return Expense{
		Date:        date,
		Category:    category,
		Amount:      amount,
		Description: description,
	}, nil
}
