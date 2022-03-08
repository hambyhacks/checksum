package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("MD5/SHA512 CHECKSUM TOOL")
	fileName := flag.String("f", "", "Filename to be checked.")
	flag.Parse()

	if os.Args[1] == "-h" || os.Args[1] == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(os.Args) < 1 {
		fmt.Printf("Check argument count.\nCurrent Argument count: %d\n", len(os.Args))
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var r = make([]byte, 1024)
	res, err := file.Read(r)
	if err != nil {
		panic(err)
	}

	md5checksum := md5.Sum(r[:res])
	md5hash := md5checksum[:]
	sha_512checksum := sha512.Sum512(r[:res])
	sha512hash := sha_512checksum[:]

	fmt.Printf("Checking the MD5 and SHA512 hash of %s...\n\n\n", *fileName)
	fmt.Printf("MD5 HASH: %s\n", hex.EncodeToString(md5hash))
	fmt.Printf("SHA512 HASH: %s\n\n", hex.EncodeToString(sha512hash))
}
