package main

import (
	"github.com/ShindouMihou/korin/pkg/korin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	korin := korin.New()
	korin.DockerBuildStep(".")
}
