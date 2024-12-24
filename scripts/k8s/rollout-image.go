package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
    appName := os.Args[1]
    image := os.Args[2]

    fmt.Println(appName, image)

	rolloutCmd := exec.Command("kubectl", "set", "image", "deployment/"+appName, appName+"="+image)
	rolloutCmd.Stdout = os.Stdout
	rolloutCmd.Stderr = os.Stderr

	if err := rolloutCmd.Run(); err != nil {
		log.Fatalf("Err: %v", err)
	}

	statusCmd := exec.Command("kubectl", "rollout", "status", "deployment/"+appName)
	statusCmd.Stdout = os.Stdout
	statusCmd.Stderr = os.Stderr

	if err := statusCmd.Run(); err != nil {
		log.Fatalf("Err: %v", err)
	}
}