default:
    just --list

compose:
    docker-compose up -d

dev: compose
    air


