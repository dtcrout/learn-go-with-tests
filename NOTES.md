# Notes

## Resources
Wiki on GitHub: https://github.com/golang/go/wiki

* Golang docs are available locally and can be launched via `$ godoc -http :8000`
* Having too many tests can become a problem in the long run; every test has its cost. Therefore the goal is not to have as many tests as possible, rather to have confidence in your codebase
* Interfaces allow you to make functions that can be used with different types
* Remember to import "errors" package into code
* `t.Fatal` will stop the test if it is called
* When refactoring your code, wrap errors into separate functions

### Structs, Methods and Interfaces

- Method syntax: `func (receiverName ReceiverType) MethodName(args) {}`
  - When your method is called on a variable of that type, you get your reference to its data via the `receiverName`

### Pointers & Errors

- If a symbol starts with a lowercase symbol then it is private outside the package it's defined in
- `&` in front of a variable recieves the memory address
- struct pointers are automatically dereferenced
- `var` keyword defines values globally to the package
- when a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception

### Maps


