package main

import (
	"github.com/steinfletcher/platform/functions/user/app"
)

func main() {
	a := app.New()
	a.Start()
}
