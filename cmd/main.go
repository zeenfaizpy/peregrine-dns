package main

import (
	"peregrinedns/parser"
)

func main() {
	domain_name := "www.twitter.com"
	record_type := parser.TypeA
	class_type := parser.ClassIn
	parser.CallUDP(domain_name, record_type, class_type)

}
