# neo4j-go-file-system

A proof of concept demonstrating how to create
a small graph of objects (files, folders, users, organizations, and policies)
that when put together look something like a rudimentary file system.

[Neo4j](https://neo4j.com/) feels like a natural choice for database, 
since our goal is to represent a file system tree and users' relationships to certain nodes in that tree.

In this model, organizations (red nodes) represent file system root.
The blue nodes are the users.
The yellow nodes are the folders.
The green nodes are the files.

<img width="1075" alt="screen shot 2019-01-17 at 10 49 13 am" src="https://user-images.githubusercontent.com/5129994/51330465-936d7300-1a45-11e9-9ca9-c1c484503fb6.png">

## Getting started
This guide expects your Neo4j password to be in the `.env` file.
If you've never configured Neo4j, the default password should be `neo4j`.

Set that in your `.env` file:
```bash
$ cat .env

NEO_PASSWORD=neo4j
```

| Command        | Description                                      |
| -------------- |:------------------------------------------------:|
| `make`         | Runs containers                                  |
| `make rebuild` | Rebuilds images from scratch and runs containers |
| `make stop`    | Stops running containers                         |

Once started, the Neo4j container will be accessible via http://localhost:7474.

Log in with the username `neo4j` and whatever password you've configured (`neo4j` is the default out-of-the-box password).

## Creating a user
```bash
curl http://localhost:8080/user -H 'Origin: http://localhost:3000' -d '{"email_address": "kevin.chen@irisvr.com",  "full_name": "Kevin Chen"}'
```

## Reading
- [Docker Compose reference](https://docs.docker.com/compose/compose-file/)
- [Background on Seabolt driver](https://medium.com/neo4j/neo4j-go-driver-is-out-fbb4ba5b3a30)
- [Neo4j with Docker](https://neo4j.com/developer/docker/)
- [neo4j-go-driver](https://github.com/neo4j/neo4j-go-driver) repo
- [seabolt](https://github.com/neo4j-drivers/seabolt) repo