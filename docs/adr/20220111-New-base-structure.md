# New base structure
Adapting project structure to a more idiomatic one for Go without departing from the postulates of DDD.

## Status
Starting implementation.

## Context
### "Avoid elaborate package hierarchies, resist the desire to apply taxonomy"
A too meticulous and nested organization of the project does not seem to provide many benefits and is not idiomatic within the Go ecosystem.
Empty intermediate directories are an indicator that the design can be improved.

## Decision
Flatten the directory structure that, in Go, do not represent a hierarchical organization in terms of packages, they only determine their location in the file system.

## Consequences
Faster location of types, functions, etc.
More familiar structure for a Go project. 

## Date
2022-01-11