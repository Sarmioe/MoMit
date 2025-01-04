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

func generateCertificate(ip string) {
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
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(24 * time.Hour),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		logError(fmt.Errorf("Failed to create certificate: %w", err))
		return
	}

	publicKeyFile := ip + ".pem"
	publicKeyOut, err := os.Create(publicKeyFile)
	if err != nil {
		logError(fmt.Errorf("Failed to open %s for writing: %w", publicKeyFile, err))
		return
	}
	defer publicKeyOut.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		logError(fmt.Errorf("Failed to marshal public key: %w", err))
		return
	}

	if err := pem.Encode(publicKeyOut, &pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes}); err != nil {
		logError(fmt.Errorf("Failed to write data to %s: %w", publicKeyFile, err))
		return
	}

	privateKeyFile := ip + "_key.key"
	privateKeyOut, err := os.Create(privateKeyFile)
	if err != nil {
		logError(fmt.Errorf("Failed to open %s for writing: %w", privateKeyFile, err))
		return
	}
	defer privateKeyOut.Close()

	privateKeyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		logError(fmt.Errorf("Failed to marshal private key: %w", err))
		return
	}

	if err := pem.Encode(privateKeyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes}); err != nil {
		logError(fmt.Errorf("Failed to write data to %s: %w", privateKeyFile, err))
		return
	}

	fmt.Printf("Self-signed certificate and key have been generated for IP %s:\n", ip)
	fmt.Printf("Public Key: %s\n", publicKeyFile)
	fmt.Printf("Private Key: %s\n", privateKeyFile)
}
