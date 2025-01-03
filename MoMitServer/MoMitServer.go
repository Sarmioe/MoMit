package main

import (
	"fmt"
)

func main() {
	ipv4, err := getIP("https://ipinfo.io/ip")
	if err != nil {
		logError(fmt.Errorf("Failed to get public IPv4: %w", err))
		ipv4 = "Not available (IPv4 unsupported)"
	}

	ipv6, err := getIP("https://v6.ipinfo.io/ip")
	if err != nil {
		logError(fmt.Errorf("Failed to get public IPv6: %w", err))
		ipv6 = "Not available (IPv6 unsupported)"
	}

	fmt.Println(" __  __       __  __ _ _   ____")
	fmt.Println("|  \\/  | ___ |  \\/  (_) |_/ ___| ")
	fmt.Println("| |\\/| |/ _ \\| |\\/| | | __\\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|")
	fmt.Println("| |  | | (_) | |  | | | |_ ___) |  __/ |   \\ V /  __/ |")
	fmt.Println("|_|  |_|\\___/|_|  |_|_|\\__|____/ \\___|_|    \\_/ \\___|_|")
	fmt.Println("Server Public IPv4 Address:", ipv4)
	fmt.Println("Server Public IPv6 Address:", ipv6)

	generateCertificate()
}
