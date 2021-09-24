# Cirrus Todo

A todo list implementation as an excuse to explore building microservices in Go based on DDD practices.

Some choices may seem overly ceremonial within the Go ecosystem, particularly when trying to employ them in a trivial example such as the one proposed here but the intention is to use this example as a reference to later create a generator utility based on these patterns.

The driving idea at the moment is to try to minimize the leap between DDD concepts and terminology and final implementation in order to decrease friction when mind mapping modeling concepts from DDD, Event Sourcing, CQRS, etc. to code and the package structure that contains it.

Request processing in a more RESTy way will be also considered if required (i.e .: during signin, signout, etc.).

Some base functionality is provided by cirrus lib but once basic structure of the service is defined those features will be replicated by custom code inside this module and used as a starting point to build the code generator.

This is a proof of concept, needs some polishing and tests implementation.

