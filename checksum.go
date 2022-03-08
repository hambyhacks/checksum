package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("MD5/SHA1/SHA256/SHA512 CHECKSUM TOOL")
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

	fmt.Printf("CALCULATING HASH OF FILE: %s\n\n", *fileName)

	go SHA1(*fileName)
	MD5(*fileName)

	wg.Add(1)
	go func() {
		defer wg.Done()
		go SHA512(*fileName)
		SHA256(*fileName)
	}()
	wg.Wait()

	fmt.Println("Done!")

}

func MD5(filename string) {
	hashmd5 := md5.New()
	md5_buf := make([]byte, 51250)

	md5_file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer md5_file.Close()
	if _, err := io.CopyBuffer(hashmd5, md5_file, md5_buf); err != nil {
		panic(err)
	}
	fmt.Printf("MD5: %x\n", hashmd5.Sum(nil))
}

func SHA1(filename string) {
	sha1_buf := make([]byte, 51250)
	sha1hash := sha1.New()

	sha1_file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer sha1_file.Close()
	if _, err := io.CopyBuffer(sha1hash, sha1_file, sha1_buf); err != nil {
		panic(err)
	}
	fmt.Printf("SHA1: %x\n", sha1hash.Sum(nil))
}

func SHA256(filename string) {
	sha256_buf := make([]byte, 51250)
	sha256hash := sha256.New()

	sha256_file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer sha256_file.Close()

	if _, err := io.CopyBuffer(sha256hash, sha256_file, sha256_buf); err != nil {
		panic(err)
	}
	fmt.Printf("SHA256: %x\n", sha256hash.Sum(nil))
}

func SHA512(filename string) {
	sha512_buf := make([]byte, 51250)
	sha512hash := sha512.New()

	sha512_file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer sha512_file.Close()
	if _, err := io.CopyBuffer(sha512hash, sha512_file, sha512_buf); err != nil {
		panic(err)
	}
	fmt.Printf("SHA512: %x\n", sha512hash.Sum(nil))
}
