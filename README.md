# Fixed Term Track - Web Server

## Status

- Development, `In Progress` (In Progress | Maintaining | Complete)
- In Production, `False` (True | False)

## Description

FYI: This is my first time writing `go` code in a 'realistic' project so - don't hate ;)

Fixed Term Track was built to help me track my Fixed Term investments (not that I have a lot of them ha.)

It would help me aggregate these term investments and track their maturity dates, giving me a heads-up on any upcoming
maturities / return dates.

At the moment though, this is a work in progress, and I'm still figuring out the best way to structure the data and the
application.

Aaaand how to interface with it, although I'm just leaning towards a mobile app to get some practice >:)

## Features

## Technologies Used

### Language

**Golang**

Some nifty parts of it:

- Scheduling: Native `time.Ticker` with a PostgreSQL backed job queue running on a separate goroutine

### Technologies

- Database: PostgreSQL
- CI/CD: GitHub Actions (WIP)
- Build: --
- Deployment: --
- Notifications: Telegram Bot API

### Tools and Libraries

- ORM: [GORM](https://github.com/go-gorm/gorm)
- Web Framework: [Gin](https://github.com/gin-gonic/gin)
- Logging: [Zap](https://github.com/uber-go/zap)
- Environment Management: [Viper](https://github.com/spf13/viper)
- Development
    - [Docker (Compose)](https://docs.docker.com/compose/) for local development
    - [Air](https://github.com/air-verse/air) for hot reloading
    - [Just](https://github.com/casey/just) for command running
    - [Mise](https://github.com/jdx/mise) for version management

## Using

## Developing

