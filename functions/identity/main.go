package main

import (
	"github.com/steinfletcher/platform/functions/identity/app"
)

func main() {
	a := app.New()
	a.Start()
}
