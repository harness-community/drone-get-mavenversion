package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	pomPath := os.Getenv("PLUGIN_POM_PATH")

	if pomPath == "" {
		fmt.Println("POM Path is empty, exiting...")
		os.Exit(1)
	}

	var pathSeparator string

	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	} else {
		pathSeparator = "/"
	}

	fmt.Println("POM Path: ", pomPath)

	// cmd := exec.Command("mvn", "-f", fmt.Sprintf("%s/pom.xml", pomPath), "help:evaluate", "-Dexpression=project.version", "-q", "-DforceStdout")
	cmd := exec.Command("mvn", "-f", fmt.Sprintf("%s%s%s", pomPath, pathSeparator, "pom.xml"), "help:evaluate", "-Dexpression=project.version", "-q", "-DforceStdout")
	output, err := cmd.Output()

	// check if os is windows

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	pomVersion := strings.TrimSpace(string(output))

	fmt.Println("POM Version: ", pomVersion)
	os.Setenv("POM_VERSION", pomVersion)
}