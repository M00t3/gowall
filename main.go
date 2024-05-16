package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type pic struct {
	name string
	url  string
}

func pp_json(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadFile(filepath string, url string) error {
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

func main() {

	site := "https://wallhaven.cc/api/v1"
	mode := "search"

	categories := "100"
	purity := "100"
	atleast := "2560x1600"
	ratios := "landscape%2Cportrait"
	topRange := "1y"
	sorting := "favorites"
	page := "1"

	url := fmt.Sprintf("%s/%s?categories=%s&atleast=%s&ratios=%s&topRange=%s&sorting=%s&page=%s&purity=%s&order=desc&ai_art_filter=1",
		site, mode, categories, atleast, ratios, topRange, sorting, page, purity)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	var json_data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&json_data)

	data := json_data["data"].([]interface{})

	var pics []pic
	save_directory := "images"
	createDirectory(save_directory)
	for index, value := range data {
		// pp_json(i)
		file_url := value.(map[string]interface{})["path"].(string)
		file_name := strconv.Itoa(index)
		pics = append(pics, pic{name: file_name, url: file_url})

		err := downloadFile(save_directory+"/"+file_name+".jpg", file_url)
		if err != nil {
		}
	}
	for _, pic := range pics {
		fmt.Printf("%v : %v \n", pic.name, pic.url)
	}

}
