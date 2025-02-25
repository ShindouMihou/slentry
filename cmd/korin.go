package main

import "github.com/ShindouMihou/korin/pkg/korin"

func main() {
	k := korin.New()
	k.Logger = korin.NoOpLogger
	k.Run("cmd/app.go")
}
