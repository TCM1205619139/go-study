package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var BaseURL = "./assets"

type PageData struct {
	Links []Link // 链接列表
}

type Link struct {
	URL  string // 链接地址
	Text string // 显示文本
}

func main() {
	fmt.Println("Starting server on :8080...")
	go walkDirFunc(BaseURL)

	http.ListenAndServe(":8080", nil)
}

func walkDirFunc(path string) error {
	var links []string
	files, err := os.ReadDir(path)
	route := strings.Replace(path, "./assets", "/static", 1)

	if err != nil {
		return fmt.Errorf("Error reading directory:", err)
	}

	for _, file := range files {
		links = append(links, file.Name())
		if file.IsDir() {
			go walkDirFunc(path + "/" + file.Name())
		} else {
			mimeType, err := detectMIME(path + "/" + file.Name())
			if err != nil {
				return fmt.Errorf("Error detecting MIME type:", err)
			}
			http.HandleFunc(route+"/"+file.Name(), func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", mimeType)
				http.ServeFile(w, r, path+"/"+file.Name())
			})
		}
	}

	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		teml := template.Must(template.ParseFiles("./index.html"))
		data := PageData{
			Links: parseLinks(route, links),
		}
		w.Header().Set("Content-Type", "text/html")
		teml.Execute(w, data)
	})

	return nil
}

func parseLinks(base string, array []string) []Link {
	links := []Link{}
	for _, link := range array {
		links = append(links, Link{
			URL:  base + "/" + link,
			Text: link,
		})
	}
	return links
}

func detectMIME(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	mimeType := http.DetectContentType(buffer)

	return mimeType, nil
}
