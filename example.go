package main

import "fmt"

// Compiler passing -ldflags "-X main.Build=build-1 -X main.Version=1.0.0-SNAPSHOT -X main.Id=M7160FC"
var (
	build, buildDate, version, id string
)

func main() {
	fmt.Printf("build: [%s]\n", build)
	fmt.Printf("buildDate: [%s]\n", buildDate)
	fmt.Printf("version: [%s]\n", version)
	fmt.Printf("id: [%s]\n", id)
}
