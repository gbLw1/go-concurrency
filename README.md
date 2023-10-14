# Go Concurrency

Testing out the Go main feature.

## Concurrency

Golang is known for several key features, but one of its most notable features is its strong support for concurrent and parallel programming.

Go provides built-in support for goroutines, which are lightweight threads, and channels for communication between goroutines.

This makes it easy to write highly concurrent and efficient code, which is particularly useful for developing concurrent network servers, distributed systems, and other applications that require handling many tasks concurrently.

## Overview

This example demonstrates how to use Go concurrency features and test them comparing the time it takes to run the two versions of the code.

For this test I'm using the [`Public APIs`](https://api.publicapis.org/entries) API.

## Run

run no concurrency fetch before to see how long it takes:

```sh
go run .\cmd\no-concurrency\main.go
```

and then run the concurrency one which runs requests in parallel with [Go routines](https://golang.org/pkg/runtime/#Goroutine):

```sh
go run .\cmd\concurrency\main.go
```
