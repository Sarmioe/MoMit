package main

import (
	"fmt"
	"os"
)

func printOutput(ipv4 string, ipv6 string, nocreatecert bool, dip bool) {
	var rip string
	rip = ipv4
	fmt.Println(" __  __       __  __ _ _   ____")
	fmt.Println("|  \\/  | ___ |  \\/  (_) |_/ ___| ")
	fmt.Println("| |\\/| |/ _ \\| |\\/| | | __\\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|")
	fmt.Println("| |  | | (_) | |  | | | |_ ___) |  __/ |   \\ V /  __/ |")
	fmt.Println("|_|  |_|\\___/|_|  |_|_|\\__|____/ \\___|_|    \\_/ \\___|_|")
	fmt.Println("Server Public IPv4 Address:", ipv4)
	fmt.Println("Server Public IPv6 Address:", ipv6)

	if ipv4 == "Not available" {
		fmt.Println("Public IPv4 address is not available.")
		if ipv6 == "Not available" {
			fmt.Println("No public IP address is available. Exiting.")
			os.Exit(1)
		}
	}
	if nocreatecert == true {
		fmt.Println("Skipping certificate generation.")
	} else {
		if dip == true {
			rip = ipv6
		}
		generateCertificate(rip)
		fmt.Println("Created a TLS certificate for Websocket connections by", rip, ".")
	}

	fmt.Println("Waiting for incoming connections...")
}
