# Fixed Term Track - Web Server

## Status

- Development, `In Progress` (In Progress | Maintaining | Complete)
- In Production, `False` (True | False)

## Description

FYI: This is my first time writing `go` code in a 'realistic' project so - don't hate ;)

Fixed Term Track was built to help me track my Fixed Term investments (not that I have a lot of them ha.)

It would help me aggregate these investments and track their maturity dates, giving me a heads-up on any upcoming maturities / return dates.

At the moment though, this is a work in progress, and I'm still figuring out the best way to structure the data and the application.
(and how to interface with it, although I'm just leaning towards a mobile app to get some practice >:))

## Features

## Technologies Used

### Language

**Golang**

Some nifty parts of it:
- Scheduling: Native `time.Ticker` with a PostgreSQL backed job queue running on a separate goroutine

### Technologies

- Database: PostgreSQL
- CI/CD: GitHub Actions (WIP)
- Build: Docker (WIP)
- Deployment: Fly.io (WIP)
- Notifications: Telegram Bot API

### Tools and Libraries

- ORM: GORM
- Web Framework: Gin
- Logging: Zap
- Environment Management: Viper
- Development
  - Docker (Compose) for local development
  - Air for hot reloading
  - [Just](https://github.com/casey/just) for command running
  - [Mise](https://github.com/jdx/mise) for version management

## Using

## Developing

