package main

import "bsuir-schedule/cmd"

var (
	Version string = "latest"
)

func main() {
	cmd.Execute(Version)
}
