# Notes

* Golang docs are available locally and can be launched via `$ godoc -http :8000`
* Having too many tests can become a problem in the long run; every test has its cost. Therefore the goal is not to have as many tests as possible, rather to have confidence in your codebase
* Interfaces allow you to make functions that can be used with different types
* Remember to import "errors" package into code
* `t.Fatal` will stop the test if it is called
* When refactoring your code, wrap errors into separate functions
