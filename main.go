package main

import (
	"encoding/json"
	"fmt"
	"gowall/utils"
	"net/http"
	"strconv"
)

type pic struct {
	name string
	url  string
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
	utils.CreateDirectory(save_directory)
	for index, value := range data {
		// pp_json(i)
		file_url := value.(map[string]interface{})["path"].(string)
		file_name := strconv.Itoa(index)
		pics = append(pics, pic{name: file_name, url: file_url})

		err := utils.DownloadFile(save_directory+"/"+file_name+".jpg", file_url)
		if err != nil {
		}
		break
	}
	for _, pic := range pics {
		fmt.Printf("%v : %v \n", pic.name, pic.url)
	}

}
