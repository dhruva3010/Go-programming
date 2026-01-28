# Go Programming Tutorial Demos

This repository contains a series of Go programming tutorial demos designed to teach fundamental and intermediate Go concepts through practical examples.

## Glossary of Tutorial Demos

### 1. Starter (`1_starter`)
**Topic**: Hello World  
**Description**: The most basic Go program that prints "Hello, World" to the console. Introduces the basic structure of a Go program including package declaration, imports, and the main function.  
**Key Concepts**: 
- Package declaration
- Import statements
- `main()` function
- `fmt.Print()` for output

---

### 2. Investment Project (`2_investment_project`)
**Topic**: Variables, User Input, Functions, and Math Operations  
**Description**: A financial calculator that computes future investment value with inflation adjustment. Takes investment amount, expected return rate, and duration as inputs.  
**Key Concepts**:
- Variable declaration (`var`)
- User input with `fmt.Scan()`
- Constants (`const`)
- Math operations and `math.Pow()`
- Multiple return values
- Named return values
- String formatting with `fmt.Sprintf()`

---

### 3. Profit Calculator (`3_profit_calculator`)
**Topic**: Error Handling and File I/O  
**Description**: Calculates business profitability metrics (EBT, net profit, and ratio) based on revenue, expenses, and tax rate. Saves results to a file.  
**Key Concepts**:
- Error handling and custom errors
- `error` type and `errors.New()`
- Input validation
- File operations with `os.WriteFile()`
- Multiple return values with errors
- Early returns for error cases

---

### 4. Bank Project (`4_bank_project`)
**Topic**: Loops, Switch Statements, Packages, and External Dependencies  
**Description**: An interactive bank management system with deposit, withdrawal, and balance inquiry features. Persists account balance to file.  
**Key Concepts**:
- Infinite loops with `for`
- `switch` statements
- Custom packages and imports
- External dependencies (`github.com/Pallinder/go-randomdata`)
- File operations for persistence
- Menu-driven program flow

---

### 5. Pointers Starter (`5_pointers_starter`)
**Topic**: Pointers and Memory Management  
**Description**: Demonstrates pointer basics by calculating adult years (age minus 18) using pointer manipulation.  
**Key Concepts**:
- Pointer declaration with `*`
- Address-of operator `&`
- Dereferencing pointers
- Mutating values via pointers
- Passing pointers to functions

---

### 6. Structs (`6_structs`)
**Topic**: Structs, Methods, and Object-Oriented Patterns  
**Description**: User management system that creates regular users and admin users with validation. Demonstrates struct usage and methods.  
**Key Concepts**:
- Struct definitions
- Constructor functions (`New()` pattern)
- Methods on structs
- Pointer receivers
- Struct embedding/composition
- Field validation
- Date/time handling

---

### 7. Custom Types (`7_custom_types`)
**Topic**: Custom Type Definitions and Methods  
**Description**: Shows how to create custom types based on built-in types and attach methods to them.  
**Key Concepts**:
- Type definitions (`type` keyword)
- Methods on custom types
- Type aliasing
- Extending built-in types

---

### 8. Notes Taking Project (`8_notes_taking_project`)
**Topic**: Advanced I/O, JSON, and Package Organization  
**Description**: A note-taking application that captures title and content from user input, displays the note, and saves it as JSON.  
**Key Concepts**:
- `bufio.Reader` for advanced input
- String manipulation (`strings` package)
- JSON encoding/marshaling
- File I/O with custom packages
- Struct tags for JSON
- Cross-platform newline handling

---

## Running the Demos

Each demo is a standalone Go module. To run any demo:

```bash
cd <demo-folder>
go run .
```

For example:
```bash
cd 1_starter
go run .
```

## Building Executables

To build an executable:

```bash
cd <demo-folder>
go build
```

This will create an executable file in the same directory.

## Prerequisites

- Go 1.x or higher installed on your system
- Basic understanding of command-line interface

## Learning Path

The demos are numbered in a suggested learning order:
1. Start with basic syntax and output
2. Progress to user input and calculations
3. Learn error handling and file operations
4. Understand control flow and packages
5. Master pointers and memory concepts
6. Explore structs and OOP patterns
7. Create custom types
8. Apply concepts in a complete project

---

*Happy Learning! 🚀*
