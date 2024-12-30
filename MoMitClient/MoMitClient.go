package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("__  __         __  __ _ _      ____ _ _            _    ")
	fmt.Println("|  \\/  | ___   |  \\/  (_) |_   / ___| (_) ___ _ __ | |_")
	fmt.Println("| |\\/| |/ _ \\  | |\\/| | | __| | |   | | |/ _ \\ '_ \\| __|")
	fmt.Println("| |  | | (_) | | |  | | | |_  | |___| | |  __/ | | | |_")
	fmt.Println("|_|  |_|\\___/  |_|  |_|_|\\__|  \\____|_|_|\\___|_| |_|\\__|")
	fmt.Println("Thanks for installing MoMit, it is running now.")
	fmt.Println("It is an encryption protocol on the Internet.")
	fmt.Println("It makes your packets look like normal Internet activity.")
	fmt.Println("So even if a man in the middle can eavesdrop on your packets, he can't get anything.")
	fmt.Println("It is an open source project, so you don't have to worry about being eavesdropped by MoMit.")
	rand.Seed(time.Now().UnixNano())
	RandomIV1 := rand.Intn(8) + 1
	IV1(RandomIV1)
	RandomIV2 := rand.Intn(7) + 9
	IV2(RandomIV2)
}

func IV1(RandomIV1 int) {
	fmt.Print("Random data is:", RandomIV1, " So, using ")
	if RandomIV1 == 1 {
		fmt.Print("TCP")
	}
	if RandomIV1 == 2 {
		fmt.Print("UDP")
	}
	if RandomIV1 == 3 {
		fmt.Print("TLS")
	}
	if RandomIV1 == 4 {
		fmt.Print("HTTPS")
	}
	if RandomIV1 == 5 {
		fmt.Print("DOT")
	}
	if RandomIV1 == 6 {
		fmt.Print("DOH")
	}
	if RandomIV1 == 7 {
		fmt.Print("mKCP")
	}
	if RandomIV1 == 8 {
		fmt.Print("gRCP")
	}
	fmt.Println(" to transmit your data packets.")
}

func IV2(RandomIV2 int) {
	fmt.Print("Random data is:", RandomIV2, " So, encrypting your data packets looks like ")
	if RandomIV2 == 9 {
		fmt.Println("watching streaming movies")
	}
	if RandomIV2 == 10 {
		fmt.Println("listen streaming music")
	}
	if RandomIV2 == 11 {
		fmt.Println("download files")
	}
	if RandomIV2 == 12 {
		fmt.Println("login a normal website")
	}
	if RandomIV2 == 13 {
		fmt.Println("play online games")
	}
	if RandomIV2 == 14 {
		fmt.Println("video chatting")
	}
	if RandomIV2 == 15 {
		fmt.Println("random datas")
	}
	fmt.Println("Let's MoMit to encrypt your data packets.")
}
