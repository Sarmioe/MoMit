package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func main() {
	ipv4, err := getIP("https://ipinfo.io/ip")
	if err != nil {
		fmt.Println("Failed to get public IPv4: %w", err)
		ipv4 = "Not available (IPv4 unsupported)"
	}

	ipv6, err := getIP("https://v6.ipinfo.io/ip")
	if err != nil {
		fmt.Println("Failed to get public IPv6: %w", err)
		ipv6 = "Not available (IPv6 unsupported)"
	}
	fmt.Println(" __  __       __  __ _ _   ____")
	fmt.Println("|  \\/  | ___ |  \\/  (_) |_/ ___| ")
	fmt.Println("| |\\/| |/ _ \\| |\\/| | | __\\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|")
	fmt.Println("| |  | | (_) | |  | | | |_ ___) |  __/ |   \\ V /  __/ |")
	fmt.Println("|_|  |_|\\___/|_|  |_|_|\\__|____/ \\___|_|    \\_/ \\___|_|")
	fmt.Println("Server Public IPv4 Address:", ipv4)
	fmt.Println("Server Public IPv6 Address:", ipv6)
	if ipv4 == "Not available (IPv4 unsupported)" {
		fmt.Println("Public IPv4 address is not available.")
		if ipv6 == "Not available (IPv6 unsupported)" {
			fmt.Println("No public IP address is available. Exiting.")
			os.Exit(1)
		}
	}
	generateCertificate(ipv4)
	fmt.Println("Created a TLS certificate for Websocket connections by IPV4.")
	fmt.Println("Waiting for incoming connections...")
}
