package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func isLoopback(ip string) bool {
	return ip == "127.0.0.1" || strings.ToLower(ip) == "localhost"
}

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func pingIP(ip string) bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", "1", ip)
	case "linux", "darwin":
		cmd = exec.Command("ping", "-c", "1", ip)
	default:
		log.Fatalf("Unsupported operating system: %s", runtime.GOOS)
	}
	err := cmd.Run()
	return err == nil
}

func parseIPFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 3 {
			log.Printf("Invalid line format: %s", line)
			continue
		}

		ip, port, publicKey := parts[0], parts[1], parts[2]
		if isValidIP(ip) {
			log.Printf("Valid IP: %s, Port: %s, PublicKey: %s", ip, port, publicKey)
			if pingIP(ip) {
				log.Printf("Ping successful for IP: %s", ip)
			} else {
				log.Printf("Ping failed for IP: %s", ip)
			}
		} else {
			log.Printf("Invalid IP address: %s", ip)
		}
	}
}
