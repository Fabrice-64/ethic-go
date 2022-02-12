package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	ipAddr := "192.168.0.1/24"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ipAddr),
		nmap.WithPorts("80, 443"),
		nmap.WithContext(ctx))

	if err != nil {
		log.Fatalln("an error occured: ", err)
	}
	r, w, err := scanner.Run()
	if err != nil {
		log.Fatalln("Error while scanning: ", err)
	}
	if w != nil {
		log.Fatalln("Warning while scanning: ", w)
	}

	for _, host := range r.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		fmt.Printf("IP: %q\n", host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Printf("MAC: %v\n", host.Addresses[1])
		}
		for _, port := range host.Ports {
			fmt.Printf("\t Port : %v Protocol: %v State: %v Service: %v\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}
