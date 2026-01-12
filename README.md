# Personal Expense Tracker – GoLang

## Overview

This project is a terminal-based Personal Expense Tracker written in Go.  
It allows users to manage expenses using basic CRUD operations. The application runs entirely in the terminal and stores data in memory while the program is running.

The project is intentionally kept simple to demonstrate Go fundamentals such as structs, interfaces, slices, and basic input/output handling.

* * *

## How to Run the Application

Make sure Go is installed on your system.

To run the application, navigate to the project directory and execute:

`go run main.go`

Once started, the application displays a menu and waits for user input.

* * *

## Application Flow

The program runs inside an infinite loop.  
Each iteration displays a menu and waits for the user to enter a choice.

`for {     // show menu     // read choice     // perform action }`

Based on the user’s choice, the corresponding CRUD operation is executed.  
The loop continues until the user selects the exit option.

* * *

## Data Structure

Each expense is represented using an `Expense` struct.

`type Expense struct {     ID     int     Title  string     Amount float64 }`

*   `ID` uniquely identifies each expense.
    
*   `Title` stores a short description.
    
*   `Amount` stores the expense value.
    

* * *

## Manager Structure

All expenses are managed using the `Manager` struct.

`type Manager struct {     expenses []Expense     nextID   int }`

*   `expenses` holds all expense records in a slice.
    
*   `nextID` is used to generate unique IDs for new expenses.
    

* * *

## Use of Interface

An interface is used to define the expected behavior of an expense manager.

`type ExpenseManager interface {     AddExpense(title string, amount float64)     ViewExpenses()     UpdateExpense(id int)     DeleteExpense(id int) }`

The `Manager` struct implements this interface.  
This helps separate definition from implementation and keeps the design clean.

* * *

## AddExpense Function

This function adds a new expense to the slice.

`func (m *Manager) AddExpense(title string, amount float64) {     expense := Expense{         ID:     m.nextID,         Title:  title,         Amount: amount,     }     m.expenses = append(m.expenses, expense)     m.nextID++ }`

The function creates a new `Expense`, assigns a unique ID, and appends it to the slice.  
The ID counter is incremented after every insert.

* * *

## ViewExpenses Function

This function prints all expenses in a formatted table.

`fmt.Printf("%-4d | %-19s | %.2f\n", e.ID, e.Title, e.Amount)`

Fixed-width formatting ensures proper alignment even when titles have different lengths.  
If no expenses exist, the user is informed.

* * *

## UpdateExpense Function

This function updates an existing expense based on ID.

`func (m *Manager) UpdateExpense(id int) {     // find expense     // ask user what to update     // modify selected fields }`

The user chooses whether to update:

*   Only the title
    
*   Only the amount
    
*   Both title and amount
    

Only the selected fields are changed.  
If the ID does not exist, an error message is displayed.

* * *

## DeleteExpense Function

This function removes an expense using its ID.

`m.expenses = append(m.expenses[:i], m.expenses[i+1:]...)`

The slice is restructured to exclude the selected expense.  
If the expense is not found, the user is notified.

* * *

## User Input Handling

User input is read using a buffered reader.

`reader := bufio.NewReader(os.Stdin) input, _ := reader.ReadString('\n')`

All input is trimmed and converted to the required data type using:

`strings.TrimSpace(input) strconv.Atoi(input) strconv.ParseFloat(input, 64)`

This ensures clean and predictable input handling.

* * *

## Data Flow Explanation

User input starts as a string entered in the terminal.  
The input is read, cleaned, and converted into appropriate data types.

This data is then passed to the manager methods, which update the in-memory slice.  
Whenever data is viewed or modified, the application directly reads from this slice.

No database or file system is involved.

* * *

## Error Handling Approach

The application handles basic errors such as:

*   Invalid menu choices
    
*   Non-existent expense IDs
    

Advanced validation is intentionally avoided to keep the code simple and readable.

* * *

## Limitations

All data is stored in memory.  
Once the application exits, all expenses are lost.

This design is intentional and suitable for a beginner-level assignment.

* * *

## Conclusion

This project demonstrates how a basic CRUD application can be built in Go using a clean and minimal design. It focuses on core language concepts and terminal interaction, making it a good learning exercise for understanding backend fundamentals.
