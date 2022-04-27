package main

import (
	"os"
	"time"
)

func main() {
	path := os.Getenv("WRITE_TO")
	content := os.Getenv("TEXT")

	date := time.Now().Format("2006-01-02T15:04:05Z07:00")
	data := []byte(date)
	if content != "" {
		data = []byte(content)
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}
