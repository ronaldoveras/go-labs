package main

import "fmt"

func main() {
	protocols := make(map[string]int)
	protocols["HTTP"] = 80
	protocols["HTTPS"] = 443
	protocols["FTP"] = 110

	for protocolo, porta := range protocols {
		fmt.Printf("Porta %d para protocolo %s. \n", porta, protocolo)
	}
}
