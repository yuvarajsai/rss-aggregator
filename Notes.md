### Initialize a new module

```bash
go mod init github.com/yuvarajsai/rss-aggregator
```

### Run the whole project

```bash
go run .
```

### Build and run the project

```bash
go build && ./rss-aggregator
```

### Read from `.env` file

```go
import "os"

secret := os.Getenv("SECRET")
```

## Log and error and exit the program

```go
import "log"

log.Fatal("Internal server error is the worst error messge. But here we go.")
```

## Install a package (godotenv)

```bash
go get github.com/joho/godotenv
```

- This command adds the package to the `go.mod` file.
- If the package is not used in the code, a yellow squiggly line may appear in some editors, indicating it is unused.
- Running `go mod tidy` will remove unused packages and their associated files from the project.

## Copy the package code into the vendor directory

```bash
go mod vendor
```

- This command creates a `vendor` directory containing all dependencies, along with a `modules.txt` file.
- The `vendor` directory ensures dependencies are locally available for builds and offline use.
- Running `go mod vendor` ensures the package files are fetched and stored locally in the `vendor` directory.
- Note: The `vendor` directory is typically used in projects where dependencies need to be bundled for offline use or to ensure consistent builds.
