package parser

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const DNSServerAddr string = "8.8.8.8:53"

func CallUDP(domain_name string, record_type uint16, class_type uint16) {

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", DNSServerAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	// defer conn.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Send a message to the server
	query := build_query(domain_name, record_type, class_type)
	_, err = conn.Write(query)
	fmt.Println("send...")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is send
	buf := make([]byte, 100)
	n, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, buf)

	var header = DNSHeader{}
	header = header.decode(buf[:12])
	fmt.Printf("%+v\n", header)

	var dname = DNSDomainName{}
	dname = dname.decode(buf[12:])
	fmt.Printf("%+v\n", dname)

	var question = DNSQuestion{}
	question = question.decode(buf[12:])
	fmt.Printf("%+v\n", question)
}
