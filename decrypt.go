package main

import (
	"bufio"
	b64 "encoding/base64"
	"log"
	"math/big"
	"os"
	"rsago/utils"
)

func main() {
	// argumentos
	// [encrypt key_file_name source_file_name dest_file_name]

	// Read key_file_name (public key)
	keyFileName := os.Args[1]
	// Read source_file_name (plain)
	srcFileName := os.Args[2]
	// Read dest_file_name (crypted)
	dstFileName := os.Args[3]

	// abre o arquivo de chaves
	err, keyFileReader := utils.GetKeyFileReader(keyFileName)
	if err != nil {
		log.Fatalln("erro abrindo arquivo de chaves: ", err)
	}

	// abre o arquivo de saída
	err, dstFile, dstWriter := utils.GetDstFileWriter(dstFileName)
	if err != nil {
		log.Fatalln("erro abrindo arquivo de saída: ", err)
	}

	//Load key, modulus from key_file_name
	modulus, key := utils.GetKeyFromFile(keyFileReader)

	originalEncodedText := ""

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	srcFileScanner := bufio.NewScanner(srcFile)

	for srcFileScanner.Scan() {
		line := srcFileScanner.Text()
		encodedChunk, _ := new(big.Int).SetString(line, 10)
		originalChunk := encodedChunk.Exp(encodedChunk, key, modulus)

		base64EncodedChunk := utils.NewBigInt(originalChunk).Text()

		originalEncodedText += base64EncodedChunk
	}

	decryptedTextBytes, _ := b64.StdEncoding.DecodeString(originalEncodedText)
	decryptedText := string(decryptedTextBytes)

	_, err = dstWriter.WriteString(decryptedText)
	if err != nil {
		log.Fatalln(err)
	}

	err = dstWriter.Flush()
	if err != nil {
		log.Fatalln("Erro ao fazer flush do arquivo de destino.", err)
	}
	err = dstFile.Close()
	if err != nil {
		log.Fatalln("Erro ao fechar o arquivo de destino.", err)
	}
}
