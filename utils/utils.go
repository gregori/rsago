package utils

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func SplitByWidth(str string, size int) []string {
	strLength := len(str)
	var split []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		split = append(split, str[i:stop])
	}
	return split
}

func GetTextFromSrcFile(err error, srcFileName string) string {
	buf, err := os.ReadFile(srcFileName)
	if err != nil {
		fmt.Printf("erro abrindo arquivo source: %v\n", err)
	}
	text := string(buf)
	return text
}

func GetDstFileWriter(dstFileName string) (error, *os.File, *bufio.Writer) {
	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0644)
	dstWriter := bufio.NewWriter(dstFile)
	return err, dstFile, dstWriter
}

func GetKeyFileReader(keyFileName string) (error, *bufio.Reader) {
	keyFile, err := os.Open(keyFileName)
	keyFileReader := bufio.NewReader(keyFile)
	return err, keyFileReader
}

func GetKeyFromFile(keyFileReader *bufio.Reader) (*big.Int, *big.Int) {
	modulusStr, _ := Readln(keyFileReader)
	modulus, _ := new(big.Int).SetString(modulusStr, 10)
	keyStr, _ := Readln(keyFileReader)
	key, _ := new(big.Int).SetString(keyStr, 10)
	return modulus, key
}
