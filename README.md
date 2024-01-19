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
│── mocks                   # mocking an interfaces for unit testing
├── pkg                     # utility packages.
├── protocol                # application's protocols to listerning on incoming traffics.
|── main.go                 # entry point to run the application
```

---
### How to run ?

Here is the example to run the application on your local machine

```make start```

the command above is an example to make the application to serve REST protocol
