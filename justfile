default:
    just --list

compose:
    docker-compose up -d

development-reload: compose
    air

development-pretty: compose
    go run cmd/main.go 2>@1 | jq

