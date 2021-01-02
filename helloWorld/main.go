// ? How do we run the code in our project?

// * Execute a project by running "go run <filename>" in the terminal

// # go run is used to compile and execute
// # go build just compiles. It stores the file in the project folder and can executed by then running main.exe

// # other CLI commands include go fmt, goinstall go get and go test
// # go install performs a similar action to npm install in that they deal with code dependancies

// ------------------------------------

// ? What does "package main" mean?

// * A package is a collection of common source code files

// # A package can have many related files inside of it. The only requirement for each file inside a package is that each file must declare the package that it belongs to at the top of each file.

// # go has two types of packages. Executable and reusable. Reusables can be thought of like dependancies, libraries or helper functions

// # main.go is the only way to make an executable file in Go.

// # main.go must always have a function called main inside it.

package main

// -------------------------------------
// ? What does import fmt mean?

// * fmt is a standard library included in Go. It allows us to print to the console.

import "fmt"

func main() {
	fmt.Println("Hello World")
}
