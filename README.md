# Introduction
---
> This project has been designed based on the hexagonal architecture, to read more https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749?gi=fb67606a28ab


### Project Structure

```
.
├── cmd                     # application commands
├── config                  # application configuration
├── intrastructure          # dealing with the infrastructure layer such as DB connection, AMQP server, etc.
├── internal                # core implementation goes here, whether core business, handlers, repositories.
│   ├── core
│   │   ├── domain          # business entities
│   │   ├── port            # abstraction layer, to separate the core business from handlers and repositories.
│   │   └── service         # application's core business.
│   │
│   ├── handler             # application's handler for the APIs.
│   └── repostiory          # repositories to dealing with the external soruces whether DB, external service, or just a simple CSV file.
|
│── mocks                   # mock interfaces for unit testing
├── pkg                     # utility packages.
├── protocol                # application's protocols to listerning on incoming traffics.
|── main.go                 # entry point to run the application
```

---
## How to run on locally?

### Prerequisite
- `make` command
- `docker`
- `go` version 1.21+

### Get started
- create `.env` file in the root project, copy the value from `.env.example`, then edit those values which suitable for your local configuration
- `make docker.up` (you can skip this step if you already have PostgreSQL on your localhost)
- `make start` to run the application
