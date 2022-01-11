# Todo

## Description

A todo list implementation as an excuse to explore building Go microservices and micromonoliths based on DDD practices.

The real advantage of this approach may not be visible in small service projects or those that do not require maintenance or extension in the long term, but it notoriously simplifies the addition of new business features on a consistent basis. On the other hand, as the implementation is defined in real business terms (ubiquitous language) and consequently freed from needing to adjust directly to a REST model (*), the code navigation and discovery becomes more natural.

Some choices may seem overly ceremonial within the Go ecosystem, particularly when trying to employ them in a trivial example such as the one proposed here but the intention is to use it as a reference to later create a generator utility based on those patterns.

Request processing in a more RESTy way will be also considered if required (i.e .: during signin, signout, etc.).

(*) A REST adapter can still make use of these commands and queries but now GET and POST requests will be enough the only meaningful information in the URL will be the name of the command to be executed or a suitable mapping. Also API documentation (OpenAPI) becomes more humanly understandable [TODO: show an example], this especially to people with less technical background.

Testing is also simplified since we are isolated from infrastructure issues (http, grpc, etc).

Finally it is possible to manually send queries and commands from console if required since it is trivial to create a CLI adapter that operates on the elements of the bounded context. This can be particularity useful during development stage.

## Structure [TODO: Update chart to reflect new structure]
![Draft](docs/images/first-draft.png?raw=true "Draft")

[WIP] This is a very basic rough draft, it's more of a placeholder for the actual chart (does GitHub render Mermaid?)

## Run
```shell
$ make run
go run cmd/todo.go
[INF] 2022/01/09 17:31:57.937272 Server json-api-server initializing at port :8081
[DBG] 2022/01/09 17:32:07.651828 Processing create-list with {Name:Home Description:What needs to be done in the house}
[INF] 2022/01/09 17:32:07.651845 CreateList name: 'Home', description: 'What needs to be done in the house'
[ERR] 2022/01/09 17:32:07.651866 error: create-list handle error: not implemented
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

### ADR

[20220111 - New base structure](docs/adr/20220111-New-base-structure.md)

## Notes

* There are a couple of type assertions associated with command processing that don't seem quite fancy. See how to avoid them.

* Directory nesting to organize package definitions can be cumbersome but for now it simplifies mapping concepts from DDD to code.

## Todo
* Simplify the structure if it really make some sense.
* Define the sequence of steps required to manually create a new command and associate it with driving adapters.
* Implement a gen code tool (command line) in order to simplify the creation of aggregates, entities, value objects, commands, adapters, etc.
* Define which is the most convenient way to feed this tool:
  * CLI shell arguments
  * JSON
  * YAML
  * SAML
  * TOML
  * HCL

