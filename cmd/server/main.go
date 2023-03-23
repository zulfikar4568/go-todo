package main

import "github.com/zulfikar4568/go-todo/internal/factory"

func main() {
	appFactory := factory.NewFactory()
	appFactory.Load()
}
