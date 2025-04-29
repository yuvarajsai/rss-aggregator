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
