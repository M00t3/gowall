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

func GetPics(page int, pics *[]pic) {
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
		*pics = append(*pics, pic{name: file_name, url: file_url})

	}
}

func main() {
	pics := []pic{}
	var err error

	var page_number int
	fmt.Print("Enter an integer for number of wallpaper that you want downlaod: ")
	_, err = fmt.Scan(&page_number)
	if err != nil {
		panic(err)
	}

	page := 0
	for {
		if len(pics) >= page_number {
			break
		}
		page++
		GetPics(page, &pics)
	}

	counter := 0
	for _, pic := range pics {
		counter++
		if counter > page_number {
			break
		}
		file := strconv.Itoa(counter)
		err = utils.DownloadFile(save_directory+"/"+file+".jpg", pic.url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Downloaded: ", counter)
	}

}
