# Go Concurrency

Testing out the Go most notable feature: concurrency.

## Overview

This example demonstrates how to use Go concurrency features and test them comparing the time it takes to run the two versions of the code.

For this test I'm using the [`Dog API`](https://dog.ceo/dog-api/) to fetch 20 random images of dogs.

## Run

Running the following command will run the normal fetching code (one by one) and then the concurrency code (in parallel).

ps: Each one of them shows the time it took to complete the requests.

```sh
go run .\cmd\main.go
```
