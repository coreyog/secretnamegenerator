//go:generate go-bindata -pkg embedded -prefix assets/ -o ./embedded/embedded.go -pkg embedded assets
package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/binary"
	"fmt"
)

const (
	AdjectiveCount uint32 = 1305
	NounCount      uint32 = 1520
)

//go:embed assets/adjectives.txt
var rawAdj []byte

//go:embed assets/nouns.txt
var rawNoun []byte

func main() {
	adjScan := bufio.NewScanner(bytes.NewBuffer(rawAdj))
	num := cryptoInt32() % AdjectiveCount
	adjScan.Scan() // must scan at least once to get anything from the scanner
	for i := uint32(0); i < num; i++ {
		adjScan.Scan()
	}
	adj := adjScan.Text()

	nounScan := bufio.NewScanner(bytes.NewBuffer(rawNoun))
	num = cryptoInt32() % NounCount
	nounScan.Scan() // must scan at least once to get anything from the scanner
	for i := uint32(0); i < num; i++ {
		nounScan.Scan()
	}
	noun := nounScan.Text()

	fmt.Printf("%s %s\n", adj, noun)
}

func cryptoInt32() uint32 {
	buffer := make([]byte, 4)
	rand.Read(buffer) // nolint, shut up linter, this won't error
	return binary.BigEndian.Uint32(buffer)
}
