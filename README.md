# neo4j-go-file-system

A proof of concept demonstrating how to create
a small graph of objects (files, folders, users, organizations, and policies)
that when put together look something like a rudimentary file system.

[Neo4j](https://neo4j.com/) felt like a natural choice to represent a graph.

## Getting started
To get started, just run `make`

This will spin up Neo4j, which you can access at http://localhost:7474

Log in with `neo4j:neo4j` and it'll ask you to change your password
(I just changed mine to `password`).