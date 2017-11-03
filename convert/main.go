package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

func main() {
	num := int64(math.MaxInt64)
	fmt.Println("IntToHex(", num, ") =", IntToHex(num))
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bianry.BigEndian: %#v num: %v\n", binary.BigEndian, num)

	return buff.Bytes()
}

/*
bianry.BigEndian: binary.BigEndian num: 9223372036854775807
IntToHex( 9223372036854775807 ) = [127 255 255 255 255 255 255 255]
*/
