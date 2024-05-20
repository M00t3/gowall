package main

import (
	"encoding/json"
	"fmt"
	"gowall/utils"
	"net/http"
	"strconv"
	"sync"
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

	for _, value := range data {
		counter++
		file_url, ok := value.(map[string]interface{})["path"].(string)
		if !ok {
			continue
		}

		file_name := strconv.Itoa(counter)

		*pics = append(*pics, pic{name: file_name, url: file_url})

	}
}

func DownlaodInParallel(pics []pic, save_directory string, pic_num int) {
	var wg sync.WaitGroup
	sem := make(chan bool, 3) // Limit to 3 parallel downloads

	counter := 0
	for _, item := range pics {
		wg.Add(1)
		go func(counter int, pic pic) {
			defer wg.Done()
			sem <- true // Will block if there are already 3 downloads in progress
			file := strconv.Itoa(counter + 1)
			err := utils.DownloadFile(save_directory+"/"+file+".jpg", pic.url)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Downloaded: ", counter+1)
			}
			<-sem // Release a slot
		}(counter, item)
		counter++
		if counter >= pic_num {
			break
		}
	}

	wg.Wait() // Wait for all downloads to finish
}

func main() {
	pics := []pic{}
	utils.CreateDirectory(save_directory)
	var err error

	var total_pic int
	fmt.Print("Enter an integer for number of wallpaper that you want downlaod: ")
	_, err = fmt.Scan(&total_pic)
	if err != nil {
		panic(err)
	}

	page := 0
	for {
		if len(pics) >= total_pic {
			break
		}
		page++
		GetPics(page, &pics)
	}

	DownlaodInParallel(pics, save_directory, total_pic)
}
