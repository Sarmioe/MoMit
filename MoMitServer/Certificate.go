package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func generateCertificate(ip string) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate private key: %v\n", err)
		return
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		fmt.Printf("Failed to generate serial number: %v\n", err)
		return
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My Organization"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(100 * 365 * 24 * time.Hour),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		fmt.Printf("Failed to create certificate: %v\n", err)
		return
	}

	certFile := ip + ".crt"
	certOut, err := os.Create(certFile)
	if err != nil {
		fmt.Printf("Failed to open %s for writing: %v\n", certFile, err)
		return
	}
	defer certOut.Close()

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		fmt.Printf("Failed to write certificate to %s: %v\n", certFile, err)
		return
	}

	privFile := ip + ".key"
	privOut, err := os.Create(privFile)
	if err != nil {
		fmt.Printf("Failed to open %s for writing: %v\n", privFile, err)
		return
	}
	defer privOut.Close()

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		fmt.Printf("Failed to marshal private key: %v\n", err)
		return
	}

	if err := pem.Encode(privOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes}); err != nil {
		fmt.Printf("Failed to write private key to %s: %v\n", privFile, err)
		return
	}

	fmt.Printf("Certificate and private key have been generated for IP %s:\n", ip)
	fmt.Printf("Certificate: %s\n", certFile)
	fmt.Printf("Private Key: %s\n", privFile)
}
