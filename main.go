package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Expense struct
type Expense struct {
	ID     int
	Title  string
	Amount float64
}

// Interface
type ExpenseManager interface {
	AddExpense(title string, amount float64)
	ViewExpenses()
	UpdateExpense(id int)
	DeleteExpense(id int)
}

// Implementation
type Manager struct {
	expenses []Expense
	nextID   int
}

func (m *Manager) AddExpense(title string, amount float64) {
	expense := Expense{
		ID:     m.nextID,
		Title:  title,
		Amount: amount,
	}
	m.expenses = append(m.expenses, expense)
	m.nextID++
	fmt.Println("Expense added successfully")
}

func (m *Manager) ViewExpenses() {
	if len(m.expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	fmt.Println("\nID   | Title               | Amount")
	fmt.Println("-------------------------------------")
	for _, e := range m.expenses {
		fmt.Printf("%-4d | %-19s | %.2f\n", e.ID, e.Title, e.Amount)
	}
}

func (m *Manager) UpdateExpense(id int) {
	for i, e := range m.expenses {
		if e.ID == id {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("What do you want to update?")
			fmt.Println("1. Title")
			fmt.Println("2. Amount")
			fmt.Println("3. Both")
			fmt.Print("Enter choice: ")

			choiceInput, _ := reader.ReadString('\n')
			choice, _ := strconv.Atoi(strings.TrimSpace(choiceInput))

			switch choice {
			case 1:
				fmt.Print("Enter new title: ")
				title, _ := reader.ReadString('\n')
				m.expenses[i].Title = strings.TrimSpace(title)

			case 2:
				fmt.Print("Enter new amount: ")
				amountInput, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
				m.expenses[i].Amount = amount

			case 3:
				fmt.Print("Enter new title: ")
				title, _ := reader.ReadString('\n')
				m.expenses[i].Title = strings.TrimSpace(title)

				fmt.Print("Enter new amount: ")
				amountInput, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
				m.expenses[i].Amount = amount

			default:
				fmt.Println("Invalid choice")
				return
			}

			fmt.Println("Expense updated successfully")
			return
		}
	}

	fmt.Println("Expense not found")
}

func (m *Manager) DeleteExpense(id int) {
	for i, e := range m.expenses {
		if e.ID == id {
			m.expenses = append(m.expenses[:i], m.expenses[i+1:]...)
			fmt.Println("Expense deleted successfully")
			return
		}
	}
	fmt.Println("Expense not found")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	manager := &Manager{nextID: 1}

	for {
		fmt.Println("\n--- Personal Expense Tracker ---")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Update Expense")
		fmt.Println("4. Delete Expense")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choiceInput, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(choiceInput))

		switch choice {
		case 1:
			fmt.Print("Enter title: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Enter amount: ")
			amountInput, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)

			manager.AddExpense(strings.TrimSpace(title), amount)

		case 2:
			manager.ViewExpenses()

		case 3:
			fmt.Print("Enter expense ID: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			manager.UpdateExpense(id)

		case 4:
			fmt.Print("Enter expense ID: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			manager.DeleteExpense(id)

		case 5:
			fmt.Println("Exiting application")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

