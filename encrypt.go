package main

import (
	b64 "encoding/base64"
	"log"
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

	//Load text from source_file_name
	text := utils.GetTextFromSrcFile(err, srcFileName)

	// obtém chunksize
	chunkSize := utils.BlockSize(*modulus)
	//codedText = base64encode(text)
	codedText := b64.StdEncoding.EncodeToString([]byte(text))

	// itera por cada grupo de caracteres quebrados pelo blockSize
	//for chunk in codedText.split(by chunk_size) do
	for _, chunk := range utils.SplitByWidth(codedText, chunkSize) {
		//originalChunk = convertToBigInt(chunk)
		originalChunk := utils.NewString(chunk).BigIntValue()
		//encodedChunk = originalChunk.modPow(key, modulus)
		encodedChunk := originalChunk.Exp(originalChunk, key, modulus)

		//save(encodedChunk, dest_file_name)
		_, _ = dstWriter.WriteString(encodedChunk.Text(10) + "\n")
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
