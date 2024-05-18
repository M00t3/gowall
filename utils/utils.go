package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ppJson(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func CreateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func DownloadFile(filepath string, url string) error {
	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a new file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the response body to the file
	_, err = io.Copy(out, resp.Body)
	return err
}
