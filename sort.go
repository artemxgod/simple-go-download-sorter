package main

import (
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	check_err(err, "reading dir")
	check_dir("music")
	check_dir("vid")
	check_dir("font")
	check_dir("img")
	check_dir("drawio")
	check_dir("docs")
	check_dir("codesample")
	photoext := []string{"jpg", "jpeg", "png", "psd", "ico", "bmp"}
	docext := []string{"pdf", "docx", "pptx", "xlsx"}
	installerext := []string{"msi", "exe"}
	codeext := []string{"cpp", "go", "py"}

	for _, file := range files {
		if !file.IsDir() {
			str := strings.Split(file.Name(), ".")
			switch  {
			case strings.ToLower(str[len(str)-1]) == "mp3":
				musicdir := "./music"
				musicdir = strings.Join([]string{musicdir, "/",file.Name()}, "")
				os.Rename(file.Name(), musicdir)
			case strings.ToLower(str[len(str)-1]) == "mp4":
				videodir := "./vid"
				videodir = strings.Join([]string{videodir, "/",file.Name()}, "")
				log.Printf(videodir)
				os.Rename(file.Name(), videodir)
			case contains_slice(photoext, strings.ToLower(str[len(str)-1])):
				photodir := "./img"
				photodir = strings.Join([]string{photodir, "/", file.Name()}, "")
				os.Rename(file.Name(), photodir)
			case str[len(str)-1] == "ttf":
				fontdir := "./font"
				fontdir = strings.Join([]string{fontdir, "/",file.Name()}, "")
				os.Rename(file.Name(), fontdir)
			case str[len(str)-1] == "drawio":
				drawiodir := "./drawio"
				drawiodir = strings.Join([]string{drawiodir, "/",file.Name()}, "")
				os.Rename(file.Name(), drawiodir)
			case contains_slice(docext, strings.ToLower(str[len(str)-1])):
				docdir := "./docs"
				docdir = strings.Join([]string{docdir, "/", file.Name()}, "")
				os.Rename(file.Name(), docdir)
			case contains_slice(installerext, strings.ToLower(str[len(str)-1])):
				installerdir := "./installer"
				installerdir = strings.Join([]string{installerdir, "/", file.Name()}, "")
				os.Rename(file.Name(), installerdir)
			case contains_slice(codeext, strings.ToLower(str[len(str)-1])):
				codedir := "./codesample"
				codedir = strings.Join([]string{codedir, "/", file.Name()}, "")
				os.Rename(file.Name(), codedir)
			default:
				log.Println(str[1])
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

func contains_slice(extensions []string, name string) bool {
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
		log.Fatal(err, msg)
	}
}