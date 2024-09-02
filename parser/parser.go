package parser

import (
	"bytes"
	"fmt"
)

type DNSRecord struct {
	name   string
	type_  uint16
	class_ uint16
	ttl    uint16
	data   []byte
}

func parse_response(dataIn []byte) {
	buf := new(bytes.Buffer)
	buf.Write(dataIn)

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
