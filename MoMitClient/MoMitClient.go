package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/websocket"
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
	if len(entries) < 2 || len(entries) > 10 {
		log.Fatalf("Please provide between 2 and 10 entries in the file.")
	}

	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}

		parts := strings.Fields(entry)
		if len(parts) != 3 {
			log.Fatalf("Invalid entry format: %s", entry)
		}

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

func connectWebSocketTLS(ip string, port string, publicKey string) error {
	tlsConfig, err := createTLSConfig(publicKey)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("wss://%s:%s", ip, port)

	dialer := websocket.Dialer{
		TLSClientConfig:  tlsConfig,
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	conn, _, err := dialer.Dial(url, http.Header{})
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Successfully connected to %s:%s\n", ip, port)
	return nil
}

func createTLSConfig(publicKey string) (*tls.Config, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to parse public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(&x509.Certificate{PublicKey: parsedKey})

	return &tls.Config{RootCAs: certPool}, nil
}

func IV1(n int) {
}

func IV2(n int) {
}
