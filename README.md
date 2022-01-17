# Todo

## Description

A todo list implementation as an excuse to explore building Go microservices and micromonoliths based on DDD practices.

The real advantage of this approach may not be visible in small service projects or those that do not require maintenance or extension in the long term, but it notoriously simplifies the addition of new business features on a consistent basis. On the other hand, as the implementation is defined in real business terms (ubiquitous language) and consequently freed from needing to adjust directly to a REST model (*), the code navigation and discovery becomes more natural.

Some choices may seem overly ceremonial within the Go ecosystem, particularly when trying to employ them in a trivial example such as the one proposed here but the intention is to use it as a reference to later create a generator utility based on those patterns.

Request processing in a more RESTy way will be also considered if required (i.e .: during signin, signout, etc.).

(*) A REST adapter can still make use of these commands and queries but now GET and POST requests will be enough the only meaningful information in the URL will be the name of the command to be executed or a suitable mapping. Also API documentation (OpenAPI) becomes more humanly understandable [TODO: show an example], this especially to people with less technical background.

Testing is also simplified since we are isolated from infrastructure issues (http, grpc, etc).

Finally, it is possible to manually send queries and commands from console if required since it is trivial to create a CLI adapter that operates on the elements of the bounded context. This can be particularity useful during development stage.

## Structure [TODO: Update chart to reflect new structure]
<img src="docs/images/first-draft.png?raw=true" alt="Draft" width="320">

# Prerequisites
Install NATS server
```shell
$ GO111MODULE=on go get github.com/nats-io/nats-server/v2
(...)
$ nats-server -m 8222
```

Alternatively
```shell
$ docker pull nats:latest
$ docker run -p 4222:4222 -ti nats:latest
````

## Run
```shell
$ make run
go run cmd/todo.go
[INF] 2022/01/17 23:55:58.549949 NATS client connecting to nats://localhost:4222
[INF] 2022/01/17 23:55:58.550059 Server rest-server initializing at port :8081
[INF] 2022/01/17 23:55:58.550754 NATS subscribed through: 127.0.0.1:4222
[INF] 2022/01/17 23:55:58.550858 Listening on 'commands' subject
```

## Call
```shell
curl --location --request POST 'http://localhost:8081/api/v1/cmd/create-list' \
--header 'Content-Type: application/json' \
--data-raw '{
  "userUUID": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "name": "Home",
  "description": "What needs to be done in the house"
}'
```

```shell
[INF] 2022/01/17 23:55:58.549949 NATS client connecting to nats://localhost:4222
[INF] 2022/01/17 23:55:58.550059 Server rest-server initializing at port :8081
[INF] 2022/01/17 23:55:58.550754 NATS subscribed through: 127.0.0.1:4222
[INF] 2022/01/17 23:55:58.550858 Listening on 'commands' subject
[INF] 2022/01/17 23:56:08.128877 NATS publishing through: 127.0.0.1:4222
[INF] 2022/01/17 23:56:08.129185 Received a command event with ID: 4f587bb9-5218-4ad1-8735-c2ee5e0b4ecb
[INF] 2022/01/17 23:56:10.709405 NATS publishing through: 127.0.0.1:4222
[INF] 2022/01/17 23:56:10.709648 Received a command event with ID: 54ac4ec9-735d-428e-baba-1078184c44e9
[INF] 2022/01/17 23:56:13.860968 NATS publishing through: 127.0.0.1:4222
[INF] 2022/01/17 23:56:13.861305 Received a command event with ID: f7f47aea-acc3-4f66-98d2-96dd442638da
```

### ADR

* [20220111-01 - New base structure](docs/adr/20220111-01-New-base-structure.md)
* [20220111-02 - Command bus](docs/adr/20220111-02-Command-bus.md)

## Notes

* There are a couple of type assertions associated with command processing that don't seem quite fancy. See how to avoid them.

* Directory nesting to organize package definitions can be cumbersome but for now it simplifies mapping concepts from DDD to code.

## GTD
### Inbox
* Implement Command Bus

### Next
* Move to a multirepo / organization
* Create three services
  * Dispatcher (commands gateway)
  * Split into three services
    * Commands gateway
    * Todo
    * Dashboard
  * Docker compose / k8s setup to orchestrate these three services, mongo and NATS server

### Someday
* Implement a gen code tool (command line) in order to simplify the creation of aggregates, entities, value objects, commands, adapters, etc.

