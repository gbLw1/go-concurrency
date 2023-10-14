package main

import (
	"go-concurrency/pkg/concurrency"
	noConcurrency "go-concurrency/pkg/no-concurrency"
)

func main() {
	noConcurrency.RunNoConcurrency()
	concurrency.RunConcurrency()
}
