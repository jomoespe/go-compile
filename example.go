package main

import "fmt"

// Compiler passing -ldflags "-X main.Build build-1 -X main.Version 1.0.0-SNAPSHOT -X main.Id M7160FC"
var Build, Version, Id string

func main() {
	fmt.Printf("Build: [%s]\n", Build)
	fmt.Printf("Version: [%s]\n", Version)
	fmt.Printf("Id: [%s]\n", Id)
}
