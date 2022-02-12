package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanPort(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	adrs := fmt.Sprintf(ip+":%d", port)
	conn, err := net.Dial("tcp", adrs)
	if err != nil {
		return
	}
	fmt.Printf("[+] Connected... port %v\n", port)
	conn.Close()
}

func main() {
	ip := "scanme.nmap.org"
	var wg sync.WaitGroup

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go ScanPort(ip, i, &wg)

	}
	wg.Wait()
}
