package main

import "fmt"

var (
	version, build, buildDate string
)

func main() {
	fmt.Printf("version:   [%s]\n", version)
	fmt.Printf("build:     [%s]\n", build)
	fmt.Printf("buildDate: [%s]\n", buildDate)
}
