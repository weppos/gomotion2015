# gomotion2015

This project was used as a practical introduction to Golang for my talk at [Codemotion 2015](http://rome2015.codemotionworld.com/).

The app is not just an example, (believe it or not) it works for real! The app exposes a system binary, for example `dig`, via an HTTP interface.

This is just an example of something you can create with Go. The project was used to introduce Golang features and design to a beginner audience, therefore expect the final app to be fairly simple. Each commit of this repo represents a baby-step towards the final app.

## Outline

This outline details the Go features the code aims to highlight on each task (represented by a single commit).

### Let's get it started

- Create folder
- Empty git repo
- Special workspace rules

### All belongs to main()

- Create main file with arbitrary name
- main package
- main() function
- Note: indentation with tabs
- $ go run app.go
- $ go run app.go && echo $?

### Hello Simone!

- single-line import
- package method invocation
- Note: title-case method
- $ go run app.go
- $ go build app.go && ./app

### Let's Dig into something cool

- function signature
- Note: title-case vs lower-case
- Note: return arguments
- Note: error management
- if works as in any other language
- boring variable declaration
- error checking in main()

### What are your Args?

- Reading CLI arguments
- Validating CLI arguments
- Note: := assignment

### Go style

- := assignment

### (one-line) Go style

- one-line assignment-if-check
- Note: scope
- Note: performance

### Exec this

- copy&paste
- execute a real dig request

### Do you HTTP?

- HTTP standard library

### More Go goodness

- defer

### Spring cleanup

- go fmt

