package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func encode(buf []byte) {
	ts := uint32(time.Now().Unix())
	binary.BigEndian.PutUint16(buf[0:], 0xa20c) // sensorID
	binary.BigEndian.PutUint16(buf[2:], 0x04af) // locationID
	binary.BigEndian.PutUint32(buf[4:], ts)     // timestamp
	binary.BigEndian.PutUint16(buf[8:], 479)    // temp
	fmt.Printf("% x\n", buf)
}

func decode(buf []byte) {
	sensorID := binary.BigEndian.Uint16(buf[0:])
	locID := binary.BigEndian.Uint16(buf[2:])
	tstamp := binary.BigEndian.Uint32(buf[4:])
	temp := binary.BigEndian.Uint16(buf[8:])
	fmt.Printf("sid: %0#x, locID %0#x ts: %0#x, temp:%d\n",
		sensorID, locID, tstamp, temp)
}

func encode_domain_name(name string) []byte {
	buf := new(bytes.Buffer)
	for _, part := range strings.Split(name, ".") {
		length := len(part)
		buf.WriteByte(byte(length))
		buf.Write([]byte(part))
		buf.WriteByte(byte(0))
	}
	return buf.Bytes()
}

func main() {
	buf := make([]byte, 10)

	encode(buf)
	decode(buf)

	val := encode_domain_name("google.com")
	fmt.Printf("% x\n", val)
}
