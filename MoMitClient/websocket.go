package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

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
