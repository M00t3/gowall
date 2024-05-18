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

func GetPics(page int) []pic {
	var pics []pic
	counter := 0
	var url string
	if api_key != "" {
		url = fmt.Sprintf("%s/%s?categories=%s&atleast=%s&ratios=%s&topRange=%s&sorting=%s&page=%s&purity=%s&apikey=%s&order=desc&ai_art_filter=1",
			site, mode, categories, atleast, ratios, topRange, sorting, strconv.Itoa(page), purity, api_key)
	} else {
		url = fmt.Sprintf("%s/%s?categories=%s&atleast=%s&ratios=%s&topRange=%s&sorting=%s&page=%s&purity=%s&order=desc&ai_art_filter=1",
			site, mode, categories, atleast, ratios, topRange, sorting, strconv.Itoa(page), purity)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var json_data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&json_data)
	data := json_data["data"].([]interface{})

	utils.CreateDirectory(save_directory)
	for _, value := range data {
		counter++
		// pp_json(i)
		file_url := value.(map[string]interface{})["path"].(string)
		file_name := strconv.Itoa(counter)
		pics = append(pics, pic{name: file_name, url: file_url})

	}
	return pics
}

func main() {
	pics := []pic{}
	for i := 1; i <= pages; i++ {
		pics = GetPics(i)
	}

	for _, pic := range pics {
		err := utils.DownloadFile(save_directory+"/"+pic.name+".jpg", pic.url)
		if err != nil {
		}
		fmt.Println("Downloaded: ", pic.name)
	}

}
