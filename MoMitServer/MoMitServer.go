package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func logError(err error) {
	if err != nil {
		logFile := time.Now().Format("2006-01-02_15-04-05") + ".log"
		f, fileErr := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if fileErr != nil {
			fmt.Println("Failed to create log file:", fileErr)
			return
		}
		defer f.Close()
		_, _ = f.WriteString(err.Error() + "\n")
	}
}

func getIP(url string) (string, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}

func generateCertificate() {
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		logError(fmt.Errorf("Failed to generate private key: %w", err))
		return
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		logError(fmt.Errorf("Failed to generate serial number: %w", err))
		return
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My Organization"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		logError(fmt.Errorf("Failed to create certificate: %w", err))
		return
	}

	certFile, err := os.Create("cert.crt")
	if err != nil {
		logError(fmt.Errorf("Failed to open cert.crt for writing: %w", err))
		return
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		logError(fmt.Errorf("Failed to write data to cert.crt: %w", err))
		return
	}

	keyFile, err := os.Create("key.key")
	if err != nil {
		logError(fmt.Errorf("Failed to open key.key for writing: %w", err))
		return
	}
	defer keyFile.Close()

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		logError(fmt.Errorf("Failed to marshal private key: %w", err))
		return
	}

	if err := pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes}); err != nil {
		logError(fmt.Errorf("Failed to write data to key.key: %w", err))
		return
	}

	fmt.Println("Self-signed certificate and key have been generated:")
	fmt.Println("Certificate: cert.crt")
	fmt.Println("Private Key: key.key")
}

func main() {
	ipv4, err := getIP("https://ipinfo.io/ip")
	if err != nil {
		logError(fmt.Errorf("Failed to get public IPv4: %w", err))
		ipv4 = "Not available (IPv4 unsupported)"
	}

	ipv6, err := getIP("https://ipv6.icanhazip.com")
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
