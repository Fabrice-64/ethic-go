# Udemy Course on Ehtical Hacking with Golang
Author : Fahad Sarwar <br>
Course available on Udemy
**Requirements: everything runs on Virtual Machines**

## Folders: <p>
- 01_test :
    * misc files for testing purpose.
- 02_cli_commands: 
    * create commands to execute using shell
- 03_change_MAC: 
    * simply change a mac address
- 04_TCP_scanning: 
    * how to scan ports
- 05_network_scanner: 
    * use nmap to scan network
- 06_address_resolution_protocol: 
    * No Use of Golang
    * used with man in the middle (needs dsniff)
    * Build 2 VM and a network on mode "host"
    * Start the VM and Kali
    * On Kali, switch to "root": ```$sudo su```
    * Install dsniff : ```$sudo apt-get install dsniff```
    * set shell: ```$sysctl -w net.ipv4.ip_forward=1```
    * Then check: ```$arp -a```
    * Syntax for spoofing: ```$arpspoof -i eth0 -t <ip victim machine or router> <ip attacker>``` 
    * On Kali : open 2 terminals on root and spoof the victim and then the router.
- 07_packet_intercept:
    - 01: find all devices. 
    Need to install library```gopacket``` and ```libcap-dev```. <br>
    Don't forget to have network on mode host<br>
    - 02 : Intercept packets
        * build up a MITM spoof using cli with 2 VM
        * launch the go script and check the interception of packets.

    
    
    



