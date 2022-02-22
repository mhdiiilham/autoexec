package autoexe

import (
	"fmt"
	"log"
	"os"
)

func CreateCFGDirectory(steamDirectory, steamID32 string) (string, error) {
	dirPath := fmt.Sprintf(CONFIG_DIRECTORY, steamDirectory, steamID32)

	fmt.Printf(`saving csgo config to: "%s\n"`, dirPath)

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalf("failed creating userdata directory: %v", err)
		return "", err
	}

	log.Printf("success creating directory for csgo config at: %s", dirPath)
	return dirPath, nil
}
