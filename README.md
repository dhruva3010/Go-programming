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

**Syntax**:
```go
package main

import "fmt"

func main() {
	fmt.Print("Hello, World")
}
```

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

**Syntax**:
```go
const inflationRate float64 = 2.5

var investmentAmount float64
fmt.Scan(&investmentAmount) // reads input into the pointer

// named return values (fv, rfv) are returned implicitly with a bare `return`
func CalculateFutureValue(amount, rate, years float64) (fv float64, rfv float64) {
	fv = amount * math.Pow(1+rate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}

fmt.Sprintf("Future Value: %.2f\n", futureValue) // format into a string
```

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

**Syntax**:
```go
func getUserInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("input must be greater than zero")
	}
	return input, nil
}

revenue, err := getUserInput("Enter Total Revenue: ")
if err != nil {
	fmt.Println("Error:", err)
	return // early return on error
}

os.WriteFile("financial_data.txt", []byte(data), 0644)
```

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

**Syntax**:
```go
import (
	"examples.com/bank_controls/fileops" // local module package
	"github.com/Pallinder/go-randomdata" // external dependency
)

for { // infinite loop, exits via `return`
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Balance:", accountBalance)
	case 2:
		// deposit logic, `continue` skips to next loop iteration
	default:
		fmt.Println("Invalid choice!")
	}
}
```

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

**Syntax**:
```go
age := 32
agePtr := &age // address-of: agePtr is *int

fmt.Println(*agePtr) // dereference to read the value

func EditAdultYears(age *int) {
	*age = *age - 18 // dereference to mutate the original variable
}
EditAdultYears(agePtr)
```

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

**Syntax**:
```go
type User struct {
	firstName string
	lastName  string
	createdAt time.Time
}

type admin struct {
	email    string
	User     // anonymous embedding -> admin inherits User's fields/methods
}

// pointer receiver: needed to mutate the struct's fields
func (user *User) ClearUserDetails() {
	user.firstName = ""
}

// constructor pattern, returns a pointer + error
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{firstName: firstName, lastName: lastName, createdAt: time.Now()}, nil
}
```

---

### 7. Custom Types (`7_custom_types`)
**Topic**: Custom Type Definitions and Methods
**Description**: Shows how to create custom types based on built-in types and attach methods to them.
**Key Concepts**:
- Type definitions (`type` keyword)
- Methods on custom types
- Type aliasing
- Extending built-in types

**Syntax**:
```go
type str string // custom type based on a built-in type

func (text str) log() { // method attached to the custom type
	fmt.Println(text)
}

var name str = "Custom Type Example"
name.log()
```

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

**Syntax**:
```go
type Note struct {
	Title     string    `json:"title"`      // struct tag controls the JSON key
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Save() error {
	data, err := json.Marshal(n) // only exported (capitalized) fields are marshalled
	if err != nil {
		return err
	}
	return os.WriteFile(strings.ToLower(n.Title)+".json", data, os.FileMode(0644))
}

reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
input = strings.TrimSuffix(input, "\n") // strip newline (and \r on Windows)
```

---

### 9. Todo Project (`9_todo_project`)
**Topic**: Interfaces, Type Assertions, and Interface Composition
**Description**: Combines the notes and todo packages behind a shared `outputtable` interface, and demonstrates runtime type checking with type assertions on an `interface{}`/`any` value.
**Key Concepts**:
- Interface definitions
- Interface composition (embedding one interface in another)
- Passing interfaces as function parameters (polymorphism)
- Type assertions (`value.(Type)`)
- `interface{}` / `any` for accepting arbitrary types
- Commented-out `switch value.(type)` as an alternative to chained assertions

**Syntax**:
```go
type saver interface {
	Save() error
}

// interface composition: outputtable requires both Save() and Display()
type outputtable interface {
	saver
	Display()
}

// any concrete type implementing both methods satisfies outputtable
func outputData(data outputtable) error {
	data.Display()
	return data.Save()
}

// type assertion: "comma, ok" form avoids a panic on mismatch
func printSomething(value interface{}) {
	if intValue, ok := value.(int); ok {
		fmt.Println("It's an integer:", intValue)
		return
	}
	// equivalent alternative:
	// switch value.(type) {
	// case int: ...
	// }
}
```

---

### 10. Generics (`10_generics`)
**Topic**: Generic Functions with Type Constraints
**Description**: A single `add` function that works across multiple numeric/string types using a generic type parameter constrained to a union of allowed types.
**Key Concepts**:
- Generic type parameters (`[T ...]`)
- Type constraints via type unions (`int | float64 | string`)
- Type inference at the call site (no explicit `[T]` needed)

**Syntax**:
```go
// T can be any of the listed types; the constraint is an inline union
func add[T int | float64 | string](a, b T) T {
	return a + b
}

add(5, 10)        // T inferred as int
add(5.5, 10.3)    // T inferred as float64
```

---

### 11. Arrays (`11_arrays`)
**Topic**: Arrays and Slices
**Description**: Explores fixed-size arrays vs. dynamic slices, slicing syntax, slice capacity vs. length, reslicing, and combining slices with `append`.
**Key Concepts**:
- Fixed-size arrays (`[N]Type{...}`)
- Dynamic slices (`[]Type{...}`)
- Slicing syntax: `arr[low:high]`, `arr[:high]`, `arr[low:]`
- `len()` vs `cap()`
- Reslicing within capacity
- `append()` and spreading a slice with `...`

**Syntax**:
```go
prices := [4]float64{19.99, 29.99, 39.99, 49.99} // fixed-size array

featured := prices[1:3]  // slice: index 1 up to (excl.) 3
featured = prices[:3]    // from start to index 3
featured = prices[1:]    // from index 1 to end

len(featured)  // number of elements
cap(featured)  // capacity from the slice's start to the underlying array's end
extended := featured[:1][:3] // reslicing rightward is allowed within capacity

dynamic := []float64{10.99, 20.99}   // slice literal (no size = dynamic)
dynamic = append(dynamic, 5.99)      // grow a slice

more := []float64{8.99, 18.99}
combined := append(dynamic, more...) // `...` spreads a slice's elements as args
```

---

### 12. Maps (`12_maps`)
**Topic**: Maps and the `make` Built-in
**Description**: Covers map creation, reading/writing/deleting keys, custom map types with methods, and using `make` to pre-allocate slices and maps.
**Key Concepts**:
- Map literals (`map[KeyType]ValueType{...}`)
- Reading, adding, and deleting keys (`delete()`)
- Custom named map types with methods
- `make()` for slices (with length & capacity) and maps
- `range` over slices (index, value) and maps (key, value)

**Syntax**:
```go
websites := map[string]string{"Google": "https://www.google.com"}
websites["LinkedIn"] = "https://www.linkedin.com" // add/update
delete(websites, "Twitter")                        // remove a key

type floatMap map[string]float64 // named map type

func (f floatMap) output() { fmt.Println(f) } // methods work on named map types

names := make([]string, 2, 5)   // slice: length 2, capacity 5
courses := make(floatMap, 3)    // map: pre-sized hint for 3 entries

for index, value := range names { /* ... */ }
for key, value := range courses { /* ... */ }
```

---

### 13. Functions (`13_functions`)
**Topic**: First-Class Functions, Closures, Recursion, and Variadic Parameters
**Description**: Demonstrates passing functions as parameters, returning functions from functions (closures that capture their environment), recursion, and variadic arguments.
**Key Concepts**:
- Function types (`type transformFn func(int) int`)
- Passing functions as parameters
- Closures (functions that capture surrounding variables)
- Functions that return functions
- Recursion
- Variadic parameters (`...int`) and spreading a slice into them

**Syntax**:
```go
type transformFn func(int) int // named function type

// function passed as a parameter
func transformNumbers(nums *[]int, transformFunc transformFn) []int {
	result := make([]int, len(*nums))
	for i, v := range *nums {
		result[i] = transformFunc(v)
	}
	return result
}

// closure: returned func "remembers" factor after transformerFunctionClosure returns
func transformerFunctionClosure(factor int) transformFn {
	return func(num int) int { return num * factor }
}
double := transformerFunctionClosure(2)
transformNumbers(&numbers, double)

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// variadic parameter: accepts any number of trailing ints
func sumUp(first int, numbers ...int) int {
	sum := first
	for _, v := range numbers {
		sum += v
	}
	return sum
}
sumUp(1, 2, 3, 4, 5)      // individual args
sumUp(1, numbers...)      // `...` spreads a slice into the variadic param
```

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
8. Apply concepts in a complete project (JSON + file I/O)
9. Use interfaces and type assertions to unify multiple types
10. Write generic functions with type constraints
11. Work with arrays and slices
12. Work with maps and the `make` built-in
13. Master functions, closures, recursion, and variadic parameters

---

*Happy Learning! 🚀*
