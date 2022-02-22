package autoexe

import (
	"io"
	"log"
	"net/http"
)

func DownloadConfig(url string) ([]byte, error) {
	var emptyBytes []byte

	log.Println("start downloading csgo config...")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get return an error: %v", err)
		return emptyBytes, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("io.ReadAll return an error: %v", err)
	}

	return bodyBytes, nil
}
