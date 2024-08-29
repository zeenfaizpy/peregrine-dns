package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"strings"
)

type DNSHeader struct {
	Id             uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
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

func (h DNSHeader) decode(dataIn []byte) DNSHeader {
	buf := new(bytes.Buffer)
	buf.Write(dataIn)

	err := binary.Read(buf, binary.BigEndian, &h)
	if err != nil {
		fmt.Println(err)
	}
	return h
}

type DNSDomainName struct {
	name string
}

func (n DNSDomainName) encode() []byte {
	buf := new(bytes.Buffer)
	// Encode domain name
	for _, part := range strings.Split(n.name, ".") {
		length := len(part)
		buf.WriteByte(byte(length))
		buf.Write([]byte(part))
	}
	buf.WriteByte(byte(0))
	return buf.Bytes()
}

func (n DNSDomainName) decode(dataIn []byte) DNSDomainName {
	buf := new(bytes.Buffer)
	buf.Write(dataIn)

	var name_parts []string

	for {
		var namePartLen uint8
		err := binary.Read(buf, binary.BigEndian, &namePartLen)
		if err != nil {
			fmt.Println(err)
		}

		if int(namePartLen) == 0 {
			break
		} else {
			var tempBuf []byte = buf.Next(int(namePartLen))
			name_parts = append(name_parts, string(tempBuf[:]))
		}
	}
	return DNSDomainName{name: strings.Join(name_parts, ".")}

}

type DNSQuestion struct {
	type_  uint16
	class_ uint16
}

func (q DNSQuestion) encode() []byte {
	buf := new(bytes.Buffer)

	// Encode type and class
	binary.Write(buf, binary.BigEndian, q.type_)
	binary.Write(buf, binary.BigEndian, q.class_)

	fmt.Printf("question = % x\n", buf)
	return buf.Bytes()
}

func (q DNSQuestion) decode(dataIn []byte) DNSQuestion {
	buf := new(bytes.Buffer)
	buf.Write(dataIn)

	err := binary.Read(buf, binary.BigEndian, &q)
	if err != nil {
		fmt.Println(err)
	}
	return q
}

func build_query(domain_name string, record_type uint16, class_type uint16) []byte {
	header := DNSHeader{
		Id:             uint16(rand.IntN(65535)), // todo: use random value
		Flags:          1 << 8,                   // todo: read RFC
		NumQuestions:   1,
		NumAnswers:     0,
		NumAuthorities: 0,
		NumAdditionals: 0,
	}
	dns_dname := DNSDomainName{
		name: domain_name,
	}
	question := DNSQuestion{
		type_:  record_type,
		class_: class_type,
	}

	buf := new(bytes.Buffer)
	buf.Write(header.encode())
	buf.Write(dns_dname.encode())
	buf.Write(question.encode())

	fmt.Printf("final = % x\n", buf)

	return buf.Bytes()
}
