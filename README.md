# Personal Expense Tracker – GoLang

## Overview

This project is a terminal-based Personal Expense Tracker written in Go.  
It allows users to manage expenses using basic CRUD operations.  
The application runs entirely in the terminal and stores data in memory while the program is running.

The project is intentionally kept simple.  
It demonstrates Go fundamentals such as structs, interfaces, slices, and basic input/output handling.

* * *

## How to Run the Application

Make sure Go is installed on your system.

Navigate to the project directory and run:

`go run main.go`

Once started, the application displays a menu and waits for user input.

* * *

## Application Flow

The program runs inside an infinite loop.  
Each iteration shows a menu and waits for a choice.

```go
for {
    // show menu
    // read choice
    // perform action
}
```

Based on the user’s choice, the corresponding CRUD operation is executed.  
The loop continues until the exit option is selected.

* * *

## Data Structure

Each expense is represented using an `Expense` struct.

```go
type Expense struct {
    ID     int
    Title  string
    Amount float64
}
```

*   `ID` uniquely identifies each expense
*   `Title` stores a short description
*   `Amount` stores the expense value

* * *

## Manager Structure

All expenses are managed using the `Manager` struct.

```go
type Manager struct {
    expenses []Expense
    nextID   int
}
```

*   `expenses` stores all expense records
*   `nextID` generates unique IDs for new expenses

* * *

## Use of Interface

An interface defines the expected behavior of an expense manager.

```go
type ExpenseManager interface {
    AddExpense(title string, amount float64)
    ViewExpenses()
    UpdateExpense(id int)
    DeleteExpense(id int)
}
```

The `Manager` struct implements this interface.  
This separates definition from implementation.

* * *

## AddExpense Function

This function adds a new expense.

A new `Expense` is created and appended to the slice.  
The ID counter increases after every insert.

* * *

### ViewExpenses Function

This function prints all expenses in a formatted table.

Fixed-width formatting keeps columns aligned.  
If no expenses exist, the user is informed.

* * *

### UpdateExpense Function

This function updates an existing expense by ID.

The user can update:

*   Only the title
*   Only the amount
*   Both title and amount

If the ID does not exist, an error message is shown.

* * *

### DeleteExpense Function

This function removes an expense using its ID.

The slice is rebuilt without the deleted expense.  
If the ID is not found, the user is notified.

* * *

## User Input Handling

Input is read using a buffered reader.

```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
```

Input is cleaned and converted using:

```go
strings.TrimSpace(input)
strconv.Atoi(input)
strconv.ParseFloat(input, 64)
```

This ensures predictable input handling.

* * *

## Data Flow Explanation

User input starts as a string.  
It is read, trimmed, and converted to required types.

The data is passed to manager methods.  
All operations work directly on the in-memory slice.

No database or file system is used.

* * *

### Error Handling Approach

The application handles basic errors:

*   Invalid menu choices
*   Non-existent expense IDs

Advanced validation is avoided for simplicity.

* * *

## Limitations

All data is stored in memory.  
Data is lost when the application exits.

This is intentional for a beginner-level project.

* * *

## Conclusion

This project demonstrates a simple CRUD application in Go.  
It focuses on clean structure, core concepts, and terminal interaction.  
It is well suited for learning backend fundamentals.
