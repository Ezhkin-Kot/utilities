package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func Start(routerName string, serverMAC string) {
}

func Connect(hostName string, userName string, serverPort string) {
	// Start Tailscale service
	err := TailscaleUp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Find host IP and connect to it using SSH
	hostIp := findHostIp(hostName)
	connectTarget := fmt.Sprintf("%s@%s", userName, hostIp)
	cmd := exec.Command("ssh", "-p", serverPort, connectTarget)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start connection and finish program
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error while connecting to server: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func Stop(hostName string, userName string, serverPort string) {
}

// Result of `tailscale status --json`
type TailscaleStatus struct {
	Peer map[string]struct {
		HostName     string   `json:"HostName"`
		TailscaleIPs []string `json:"TailscaleIPs"`
	} `json:"Peer"`
}

// Start Tailscale service
func TailscaleUp() error {
	err := exec.Command("tailscale", "up").Run()
	return err
}

// Find host IP in Tailscale by host name
func findHostIp(hostName string) string {
	// Get Tailscale status in JSON
	out, err := exec.Command("tailscale", "status", "--json").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var status TailscaleStatus
	err = json.Unmarshal(out, &status)
	if err != nil {
		fmt.Println("Tailscale JSON parse error:", err)
		os.Exit(1)
	}

	// Extract host IP
	var ip string
	for _, peer := range status.Peer {
		if peer.HostName == hostName {
			ip = peer.TailscaleIPs[0]
			break
		}
	}

	if ip == "" {
		fmt.Println("Host not found in Tailscale")
		os.Exit(1)
	}

	return ip
}
