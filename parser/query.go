package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type DNSHeader struct {
	id              uint16
	flags           uint16
	num_questions   uint16
	num_answers     uint16
	num_authorities uint16
	num_additionals uint16
}

func (h DNSHeader) encode() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, h)
	if err != nil {
		fmt.Println(err)
	}
	// return fmt.Sprintf("% x", buf.Bytes())
	fmt.Printf("header = % x\n", buf)
	return buf.Bytes()
}

type DNSQuestion struct {
	name   string
	type_  uint16
	class_ uint16
}

func (q DNSQuestion) encode() []byte {
	buf := new(bytes.Buffer)

	// Encode domain name
	for _, part := range strings.Split(q.name, ".") {
		length := len(part)
		buf.WriteByte(byte(length))
		buf.Write([]byte(part))
	}
	buf.WriteByte(byte(0))

	// Encode type and class
	binary.Write(buf, binary.BigEndian, q.type_)
	binary.Write(buf, binary.BigEndian, q.class_)

	fmt.Printf("question = % x\n", buf)
	return buf.Bytes()
}

func build_query(domain_name string, record_type uint16, class_type uint16) []byte {
	header := DNSHeader{
		id:              0x1314, // todo: use random value
		flags:           1 << 8, // todo: read RFC
		num_questions:   1,
		num_answers:     0,
		num_authorities: 0,
		num_additionals: 0,
	}
	question := DNSQuestion{
		name: domain_name,
		type_:  record_type,
		class_: class_type,
	}

	buf := new(bytes.Buffer)
	buf.Write(header.encode())
	buf.Write(question.encode())

	fmt.Printf("final = % x\n", buf)

	return buf.Bytes()
}
