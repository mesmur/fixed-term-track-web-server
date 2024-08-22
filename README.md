# Fixed Term Track - Web Server

![Go Version](https://img.shields.io/github/go-mod/go-version/mesmur/fixed-term-track-web-server)
![License](https://img.shields.io/github/license/mesmur/fixed-term-track-web-server)
![Go Report Card](https://goreportcard.com/badge/github.com/mesmur/fixed-term-track-web-server)

![Last Commit](https://img.shields.io/github/last-commit/mesmur/fixed-term-track-web-server)

![Issues](https://img.shields.io/github/issues/mesmur/fixed-term-track-web-server)
![Pull Requests](https://img.shields.io/github/issues-pr/mesmur/fixed-term-track-web-server)
![Pull Requests](https://img.shields.io/github/issues-pr/mesmur/fixed-term-track-web-server)


## Status

- Development, `In Progress` (In Progress | Maintaining | Archived)

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
- [x] Docker Image for deployment
- Comprehensive Testing
  - [ ] Unit Tests
  - [ ] Integration Tests
- Documentation
  - [ ] GoDoc (?)
  - [ ] Openapi Specification for APIs

## Technologies Used

### Language

**Golang**

### Technologies

- Database: [PostgreSQL](https://www.postgresql.org/)
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
- A telegram bot token and your chat ID in the `.env` file

Do the below:

```bash
go mod download # Get the dependencies
go mod verify # Verify the dependencies
cp .env.example .env # Copy the example env file, make sure to get your telegram bot token!
just dev # Starts the server
```

## Examples

Check the `examples` directory for examples on setting up the server.

1. `running-server-with-compose`
   1. There's an example of a compose file that runs the server with a postgres instance, it pulls the latest image from dockerhub
   2. The `.env` is slightly different (since it connects to postgres on the internal network), and the telegram credentials are required

## Contributing

Feel free to fork and raise a PR if you're interested in contributing to this project :)

You can create an `issue` to raise any bugs or have a suggestion.

If you do use this project as a reference, let me know! I'd love to hear about it. 
