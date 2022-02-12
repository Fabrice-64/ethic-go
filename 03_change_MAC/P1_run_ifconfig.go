package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func ExecCmd(cmd string, args_arr []string) (err error) {
	args := args_arr
	cmd_obj := exec.Command(cmd, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr
	cmd_obj.Stdin = os.Stdin

	err = cmd_obj.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil
}

func main() {
	iface := flag.String("iface", "eth0", "Interface with which you want to change the MAC")
	newMac := flag.String("newMac", "", "Provide the new MAC address")

	flag.Parse()

	ExecCmd("sudo", []string{"ifconfig", *iface, "down"})
	ExecCmd("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
	ExecCmd("sudo", []string{"ifconfig", *iface, "up"})

}

// to change the MAC address using cli, type:
// go run P1_run_ifconfig.go -iface=eth0 -newMac=00:11:22:33:44:66
