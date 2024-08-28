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

func (h DNSHeader) to_bytes() []byte {
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
	type_  uint16
	class_ uint16
}

func (q DNSQuestion) to_bytes() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, q)
	if err != nil {
		fmt.Println(err)
	}
	// return fmt.Sprintf("% x", buf.Bytes())
	fmt.Printf("question2 = % x\n", buf)
	return buf.Bytes()
}

// func (q DNSQuestion) to_gob_bytes() []byte {
// 	buf := new(bytes.Buffer)

// 	enc := gob.NewEncoder(buf)
// 	if err := enc.Encode(q); err != nil {
// 		fmt.Println(err)
// 	}
// 	// return fmt.Sprintf("% x", buf.Bytes())
// 	fmt.Printf("question = % x\n", buf)
// 	return buf.Bytes()
// }

func encode_domain_name(name string) []byte {
	buf := new(bytes.Buffer)
	for _, part := range strings.Split(name, ".") {
		length := len(part)
		buf.WriteByte(byte(length))
		buf.Write([]byte(part))
	}
	buf.WriteByte(byte(0))
	fmt.Printf("question1 = % x\n", buf)
	return buf.Bytes()
}

func build_query(domain_name string, record_type uint16, class_type uint16) []byte {
	header := DNSHeader{
		id:              0x1314,
		flags:           1 << 8,
		num_questions:   1,
		num_answers:     0,
		num_authorities: 0,
		num_additionals: 0,
	}
	question := DNSQuestion{
		type_:  record_type,
		class_: class_type,
	}

	buf := new(bytes.Buffer)
	buf.Write(header.to_bytes())
	name := encode_domain_name(domain_name)
	buf.Write(name)
	buf.Write(question.to_bytes())

	fmt.Printf("final = % x\n", buf)

	return buf.Bytes()
}
