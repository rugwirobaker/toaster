[![Build Status](https://cloud.drone.io/api/badges/rugwirobaker/toaster/status.svg)](https://cloud.drone.io/rugwirobaker/toaster)
[![codebeat badge](https://codebeat.co/badges/007155a0-c7a5-495e-bda9-6e4b786f69e4)](https://codebeat.co/projects/github-com-rugwirobaker-toaster-master)

# Toaster

Toaster is  a rudimentary sql database that implements a subset of the postgres protocol.
Toaster will uses dgraph's [badger](https://github.com/dgraph-io/badger)(introduced in this nice blog post: https://blog.dgraph.io/post/badger) as the storage backend to build on the shoulder of giants. Since it's main purpose is learning it will forgo many of the complex features of modern DBMS.

## planned features

I really don't know how this will look in the end but here some of my current goals:

1. subset of postgres sql(create, insert, select, delete, update statements)
2. support all major data types(numeric, text, bool,...)
3. support go's `database/sql`
4. more....

## current state

1. a small subset of sql vocabulary.
2. a simple repl to check which keywords are currently supported.

## development

NOTE: you need to install go for the following to work

1. **run tests**: go test ./...
2. **run the repl**: go run main.go

## example

```
➜ go run cmd/main.go
Hello rugwirobaker! welcome to toaster!
Feel free to type in sql queries
To exit this prompt type in: '\q'

>> CREATE TABLE users (id INT, name TEXT, active BOOL);
{Literal:create Kind:CREATE}
{Literal:table Kind:TABLE}
{Literal:users Kind:IDENT}
{Literal:( Kind:(}
{Literal:id Kind:IDENT}
{Literal:int Kind:INT}
{Literal:, Kind:,}
{Literal:name Kind:IDENT}
{Literal:text Kind:TEXT}
{Literal:, Kind:,}
{Literal:active Kind:IDENT}
{Literal:bool Kind:BOOL}
{Literal:) Kind:)}
{Literal:; Kind:;}
>> \q
```

## trivia

I like to think that toaser is an sql version of the monkey language described in [writing an interpreter in go](https://interpreterbook.com/) book
