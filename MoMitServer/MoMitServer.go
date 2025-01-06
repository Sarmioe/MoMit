package main

import (
	"fmt"
	"os"
)

func main() {
	var ipv4, ipv6 string
	var nocreatecert bool = false
	var noCheckIP bool = false
	config, err := readConfig("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("INFO: config.json not found, default behavior will be applied.")
		} else {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	} else {
		if config.NocreateCert {
			nocreatecert = true
		}
	}

	if !noCheckIP {
		ipv4, err = getIP("https://ipinfo.io/ip")
		if err != nil {
			fmt.Println("Failed to get public IPv4:", err)
			ipv4 = "Not available"
		}

		ipv6, err = getIP("https://v6.ipinfo.io/ip")
		if err != nil {
			fmt.Println("Failed to get public IPv6:", err)
			ipv6 = "Not available"
		}
	} else {
		fmt.Println("INFO: IP address checking is disabled, so the certificate will not be generated.")
		ipv4 = "Skipped"
		ipv6 = "Skipped"
		nocreatecert = true
	}
	rip := ipv4
	printOutput(ipv4, ipv6, nocreatecert, rip)
}
