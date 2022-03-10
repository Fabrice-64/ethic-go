package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}
	for idx, dev := range devices {
		fmt.Println(idx, "->", dev.Name)
		for _, addr := range dev.Addresses {
			fmt.Println("/tIP: ", addr.IP)
			fmt.Println("/tNetMask:", addr.Netmask)
		}
	}
}
