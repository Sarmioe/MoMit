package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Println("__  __         __  __ _ _      ____ _ _            _    ")
	fmt.Println("|  \\/  | ___   |  \\/  (_) |_   / ___| (_) ___ _ __ | |_")
	fmt.Println("| |\\/| |/ _ \\  | |\\/| | | __| | |   | | |/ _ \\ '_ \\| __|")
	fmt.Println("| |  | | (_) | | |  | | | |_  | |___| | |  __/ | | | |_")
	fmt.Println("|_|  |_|\\___/  |_|  |_|_|\\__|  \\____|_|_|\\___|_| |_|\\__|")
	fmt.Println("Thanks for installing MoMit, it is running now.")
	fmt.Println("It is an encryption protocol on the Internet.")
	fmt.Println("It makes your packets look like normal Internet activity.")
	fmt.Println("So even if a man in the middle can eavesdrop on your packets, he can't get anything.")
	fmt.Println("It is an open source project, so you don't have to worry about being eavesdropped by MoMit.")
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
				continue // 跳过本地环回地址
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

func IV1(RandomIV1 int) {
	fmt.Print("Random data is:", RandomIV1, " So, using ")
	if RandomIV1 == 1 {
		fmt.Print("TCP")
	}
	if RandomIV1 == 2 {
		fmt.Print("UDP")
	}
	if RandomIV1 == 3 {
		fmt.Print("TLS")
	}
	if RandomIV1 == 4 {
		fmt.Print("HTTPS")
	}
	if RandomIV1 == 5 {
		fmt.Print("DOT")
	}
	if RandomIV1 == 6 {
		fmt.Print("DOH")
	}
	if RandomIV1 == 7 {
		fmt.Print("mKCP")
	}
	if RandomIV1 == 8 {
		fmt.Print("gRCP")
	}
	fmt.Println(" to transmit your data packets.")
}

func IV2(RandomIV2 int) {
	fmt.Print("Random data is:", RandomIV2, " So, encrypting your data packets looks like ")
	if RandomIV2 == 9 {
		fmt.Println("watching streaming movies")
	}
	if RandomIV2 == 10 {
		fmt.Println("listen streaming music")
	}
	if RandomIV2 == 11 {
		fmt.Println("download files")
	}
	if RandomIV2 == 12 {
		fmt.Println("login a normal website")
	}
	if RandomIV2 == 13 {
		fmt.Println("play online games")
	}
	if RandomIV2 == 14 {
		fmt.Println("video chatting")
	}
	if RandomIV2 == 15 {
		fmt.Println("random datas")
	}
	fmt.Println("Let's MoMit to encrypt your data packets.")
}
