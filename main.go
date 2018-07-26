package main

import (
	"fmt"
	"flag"
	"os"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base32"
	"encoding/pem"
	"strings"
)


const nl = byte('\n')

var size = flag.Int("size", 1024, "number of bits for key")
var num = flag.Int("num", 1, "number of keys to generate")
var numProcs = flag.Int("num-procs", 1 , "number of processors")
var v2 = flag.Bool("version", true, "generate v2 (current) onion service addresses")
var v3 = flag.Bool("version", false, "generate v2 (current) onion service addresses")

var pe = os.Stderr.Write
var sprintf = fmt.Sprintf

func pes(str string) {
	pe([]byte(str))
}

func main() {
	flag.Parse()

	pes(sprintf("size: %d, numProcs: %d\n", *size, *numProcs))

	for i:=0; i<(*num); i++ {
		generateV2Address()
	}

	pes(sprintf("Done, byee\n"))
}

func generateV2Address() {
	sk, err := rsa.GenerateKey(rand.Reader, *size)

	if err != nil {
		panic("SK - Crap! " + err.Error() + "\n")
	}
	
	err = sk.Validate()
	if err != nil {
		panic("Valid - Crap! " + err.Error() + "\n")
	}

	pk := sk.PublicKey
	skBytes := x509.MarshalPKCS1PrivateKey(sk)
	pkBytes := x509.MarshalPKCS1PublicKey(&pk)

	sha1 := sha1.Sum(pkBytes)

	addr := strings.ToLower(base32.StdEncoding.EncodeToString(sha1[:10])) // 10 bytes or 80 bits => 16 bytes in base32
	pes(addr + ".onion\n")
	//pes(sprintf("D:%d\nN:%d\nE:%d\n", sk.D, pk.N, pk.E))
	pes("\n\n")

	skPemBlock := pem.Block{
		Type : "RSA PRIVATE KEY",
		Headers : nil,
		Bytes : skBytes,
	}
	pem.Encode(os.Stdout, &skPemBlock)

	pes("\n\n")

	pkPemBlock := pem.Block{
		Type : "RSA PUBLIC KEY",
		Headers : nil,
		Bytes : pkBytes,
	}
	pem.Encode(os.Stdout, &pkPemBlock)

	pes("\n\n")

}
