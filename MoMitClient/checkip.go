package main

import (
	"log"
	"net"
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
