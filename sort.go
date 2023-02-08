package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) == 2 {
		godotenv.Write(map[string]string{
			"DOWNLOAD_PATH": os.Args[1],
		}, "./.env")
		err := godotenv.Load()
		check_err(err, "loading .env file")
		if _, err = os.Stat(os.Getenv("DOWNLOAD_PATH")); err != nil {
			godotenv.Write(map[string]string{
				"DOWNLOAD_PATH": "~/Downloads",
			}, "./.env")
			log.Fatal(err)
		}
	}
	err := godotenv.Load()
	check_err(err, "loading .env file")
	download_path := os.Getenv("DOWNLOAD_PATH")
	// log.Println(download_path)
	files, err := ioutil.ReadDir(download_path)
	check_err(err, "reading dir")
	dirs := []string {"music", "vid", "img", "font", "drawio", "docs", "installer", "codesample"}
	for idx := range dirs {
		dirs[idx] = download_path + "/" + dirs[idx]
	}
	extmap := map[string][]string {
		dirs[0]: 	{"mp3"},
		dirs[1]: 	{"mp4"},
		dirs[2]:	{"jpg", "jpeg", "png", "psd", "ico", "bmp", "svg", "heic"},
		dirs[3]:  	{"ttf", "woff"},
		dirs[4]:	{"drawio"},
		dirs[5]:	{"pdf", "docx", "pptx", "pptm", "xlsx"},
		dirs[6]:	{"msi", "exe", "dmg"},
		dirs[7]:	{"cpp", "go", "py"},
	}

	for _, file := range files {
		if !file.IsDir() {
			str := strings.Split(file.Name(), ".")
			for _, dir := range dirs {
				//log.Println(str)
				if contains_ext(strings.ToLower(str[len(str)-1]), extmap[dir]) {
					check_dir(dir)
					new_name := strings.Join([]string{dir, "/", file.Name()}, "")
					log.Printf("dir - %s\nfilename - %s\nextmap - %s\nnewname - %s\n",dir, str, extmap[dir], new_name)
					err = os.Rename(download_path + "/" + file.Name(), new_name)
					check_err(err, "rename")
				}
			}
		}
	}
}

func check_dir(name string) {
	_, err := os.ReadDir(name)
	if err != nil {
		os.Mkdir(name, os.ModePerm)
	}
}

func contains_ext(name string, extensions []string) bool {
	for _, ext := range extensions {
		if ext == name {
			return true
		}
	}
	return false
}

func check_err(err error, msg string) {
	if err != nil {
		msg := strings.ToUpper(msg)
		log.Fatal(msg, "\t", err)
	}
}
