package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("__  __         __  __ _ _      ____ _ _            _    ")
	fmt.Println("|  \\/  | ___   |  \\/  (_) |_   / ___| (_) ___ _ __ | |_")
	fmt.Println("| |\\/| |/ _ \\  | |\\/| | | __| | |   | | |/ _ \\ '_ \\| __|")
	fmt.Println("| |  | | (_) | | |  | | | |_  | |___| | |  __/ | | | |_")
	fmt.Println("|_|  |_|\\___/  |_|  |_|_|\\__|  \\____|_|_|\\___|_| |_|\\__|")
	fmt.Println("Thanks for install MoMit, it is running now.")
	rand.Seed(time.Now().UnixNano())
	RandomIV1 := rand.Intn(8) + 1
	IV1(RandomIV1)
	RandomIV2 := rand.Intn(7) + 9
	IV2(RandomIV2)

	filePath := "ip.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	content := string(data)
	fmt.Println("File content:", content)

	entries := strings.Split(content, "\n")
	validEntries := make([]string, 0)

	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}
		parts := strings.Fields(entry)
		if len(parts) != 3 {
			log.Fatalf("Invalid entry format: %s", entry)
		}
		validEntries = append(validEntries, entry)
	}

	if len(validEntries) < 2 || len(validEntries) > 10 {
		log.Fatalf("Please provide between 2 and 10 valid entries in the file.")
	}

	for _, entry := range validEntries {
		parts := strings.Fields(entry)
		ip := parts[0]
		port := parts[1]
		publicKeyFile := ip + ".pem"

		if isLoopback(ip) {
			continue
		}
		if !isValidIP(ip) {
			log.Fatalf("Invalid IP address: %s", ip)
		}

		publicKeyData, err := ioutil.ReadFile(publicKeyFile)
		if err != nil {
			log.Fatalf("Failed to read public key file %s: %v", publicKeyFile, err)
		}

		err = connectWebSocketTLS(ip, port, string(publicKeyData))
		if err != nil {
			log.Printf("Failed to connect to %s:%s: %v", ip, port, err)
		}
	}
}
