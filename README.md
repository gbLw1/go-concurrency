# Go Concurrency

Testing out the Go most notable feature: concurrency.

## Overview

This example demonstrates how to use Go concurrency features and test them comparing the time it takes to run the two versions of the code.

For this test I'm using the [`Dog API`](https://dog.ceo/dog-api/).

## Run

run no concurrency fetch before to see how long it takes:

```sh
go run .\cmd\no-concurrency\main.go
```

and then run the concurrency one which runs requests in parallel

```sh
go run .\cmd\concurrency\main.go
```
