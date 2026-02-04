# WhereToGo
### Golang WhereTo Backend

Run: 
```bash
go run cmd/server/main.go
```

Alternatively, install air for live reloading during development:
```bash
go install github.com/cosmtrek/air@latest
```

Recommended .air.toml:
```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/app ./cmd/server"
bin = "tmp/app"
include_ext = ["go"]
exclude_dir = ["tmp", "vendor"]
```

Then run:
```bash
air
```
