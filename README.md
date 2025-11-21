# Go Expense Tracker (Console-Based App)

A simple command-line expense tracker written in Go.

The app supports:
- Adding expenses
- Viewing expenses
- Updating expenses
- Deleting expenses
- Storing data in a CSV file

## How to Build

go build -o expense_tracker_go.exe ./cmd

This will create the `expense_tracker_go.exe` binary in the project root.

## How to Run

.\expense_tracker_go.exe

## Data Storage

Expenses are saved in:

data/expenses.csv

This CSV file is ignored by Git, but the folder stays in the repo (via `.gitkeep`) so everything works right after cloning.

## Requirements

- Go 1.21+
- Windows for `.exe` builds, or remove `.exe` when building on Linux/Mac

## Notes

This is a small project to refresh my skills and get more comfortable with Go coming from a C++ background.
