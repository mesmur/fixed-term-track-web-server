# Fixed Term Track - Web Server

![Go Version](https://img.shields.io/github/go-mod/go-version/mesmur/fixed-term-track-web-server)
![License](https://img.shields.io/github/license/mesmur/fixed-term-track-web-server)
![Go Report Card](https://goreportcard.com/badge/github.com/mesmur/fixed-term-track-web-server)

![Last Commit](https://img.shields.io/github/last-commit/mesmur/fixed-term-track-web-server)

![Issues](https://img.shields.io/github/issues/mesmur/fixed-term-track-web-server)
![Pull Requests](https://img.shields.io/github/issues-pr/mesmur/fixed-term-track-web-server)
![Pull Requests](https://img.shields.io/github/issues-pr/mesmur/fixed-term-track-web-server)

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

- [x] Creating a record of investment
- [x] Creating a record of investment returns
- [x] Retrieving metrics on investments and associated returns
- [x] Auto-scheduling reminders for upcoming returns through telegram (so you don't forget to check ;>)
- [ ] Validations on input data
- [ ] Error handling
- Comprehensive Testing
  - [ ] Unit Tests
  - [ ] Integrations
  - [ ] E2E Tests
- Documentation
  - [ ] GoDoc
  - [ ] Openapi Specification for APIs

## Technologies Used

### Language

**Golang**

Some nifty parts of it:

- Scheduling: Native `time.Ticker` with a PG backed job queue running on a separate goroutine

### Technologies

- Database: PostgreSQL (courtesy of [Supabase](https://supabase.com/))
- Deployment: (TBD)
- Notifications: [Telegram Bot API](https://core.telegram.org/bots/api)

### Tools and Libraries

- ORM: [GORM](https://github.com/go-gorm/gorm)
- Web Framework: [Gin](https://github.com/gin-gonic/gin)
- Logging: [Zap](https://github.com/uber-go/zap)
- Environment Management: [Viper](https://github.com/spf13/viper)
- Development:
    - [Docker (Compose)](https://docs.docker.com/compose/) for local development
    - [Air](https://github.com/air-verse/air) for hot reloading
    - [Just](https://github.com/casey/just) for command running
    - [Mise](https://github.com/jdx/mise) for version management

## Developing

Prerequisites:
- Go 1.22
- The tools listed above under `Development` (`mise` is optional)

Do the below:

```bash
go mod tidy # Get the dependencies
just dev # Starts the server
```

## Contributing

I doubt anyone's interested in contributing to this, but if you are, feel free to fork and PR.

If you do use this project as a reference, let me know! I'd love to hear about it. You can create an `issue` to raise any bugs or suggested features.
