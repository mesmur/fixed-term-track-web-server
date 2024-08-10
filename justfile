default:
    just --list

development-reload:
    air

development-pretty:
    go run cmd/main.go 2>@1 | jq

