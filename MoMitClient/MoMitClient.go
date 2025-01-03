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

	ips := strings.Split(content, "\n")
	reachableIPs := []string{}
	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			if isLoopback(ip) {
				continue
			}
			if !isValidIP(ip) {
				log.Fatalf("Invalid IP address: %s", ip)
			}
			if pingIP(ip) {
				reachableIPs = append(reachableIPs, ip)
			}
		}
	}

	fmt.Println("Reachable IP addresses:")
	for i, ip := range reachableIPs {
		fmt.Printf("IP %d: %s\n", i+1, ip)
	}
	fmt.Printf("Total reachable IPs: %d\n", len(reachableIPs))
}
