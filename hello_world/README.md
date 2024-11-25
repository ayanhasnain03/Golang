# Hello World in Go

This repository contains a simple Go program that prints "Hello World" to the console.

---

## Code Explanation

### `package main`
The program starts with the `main` package, which serves as the entry point for the Go application. Every executable Go program must have a `main` package.

---

### `import "fmt"`
The `"fmt"` package is imported to enable input and output operations. It provides functions like:
- `Println`: Prints a line of text followed by a newline.
- `Printf`: Allows formatted output.

---

### The Code
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

Purpose: Prints Hello World to the console.
How to Run the Program
Prerequisites:

Install Go on your system. You can download it from Go's official website.
Steps:

Save the program in a file named hello_world.go.

Open your terminal and navigate to the directory containing the file.
Run the program using:

go run hello_world.go
Expected Output:
Hello World
