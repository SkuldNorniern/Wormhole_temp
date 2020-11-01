package network

import (
	"log"

	"github.com/stacktitan/smb/smb"
)

func Hostsmb() {
	host := GetOutboundIP()
	//addrs, _ := net.LookupIP(host)
	options := smb.Options{
		Host:        host,
		Port:        445,
		User:        "alice",
		Domain:      "corp",
		Workstation: "",
		Password:    "Password123!",
	}
	debug := false
	session, err := smb.NewSession(options, debug)
	if err != nil {
		log.Fatalln("[!]", err)
	}
	defer session.Close()

	if session.IsSigningRequired {
		log.Println("[-] Signing is required")
	} else {
		log.Println("[+] Signing is NOT required")
	}

	if session.IsAuthenticated {
		log.Println("[+] Login successful")
	} else {
		log.Println("[-] Login failed")
	}

	if err != nil {
		log.Fatalln("[!]", err)
	}
}
