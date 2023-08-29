package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	json := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")
	if json == "" {
		log.Fatalln("Environment variable GOOGLE_APPLICATION_CREDENTIALS_JSON is not set")
	}
	outputJsonPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	if outputJsonPath == "" {
		log.Fatalln("Environment variable GOOGLE_APPLICATION_CREDENTIALS is not set")
	}
	jsonDir := filepath.Dir(outputJsonPath)

	if _, err := os.Lstat(jsonDir); err != nil {
		if err := os.MkdirAll(jsonDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	if err := os.WriteFile(outputJsonPath, []byte(json), 0755); err != nil {
		log.Fatal(err)
	}

	log.Printf("Successful set GOOGLE_APPLICATION_CREDENTIALS=%s \n", outputJsonPath)
}
