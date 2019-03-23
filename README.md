# Book Notes Generator

This CLI tool generates a basic markdown outline for a book. Just execute the command, answer the prompts and you'll have a solid template for book notes.

## Commands

### Executing Tests

``` Shell
go test -coverprofile=c.out
go tool cover -html=c.out
```

### Building the Project

I chose to use packr for this project to package templates. As such, you need to use packr build.

``` Shell
packr build
packr install
```