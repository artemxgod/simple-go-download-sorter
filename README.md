# simple-go-download-sorter

## Introduction

### The download sorter will help you organize your downloads
* It creates folders needed for uploaded files
* Arranges files in created folders by their extension

### Supported extensions
| Folder     | Extension                |
|------------|--------------------------|
| music      | mp3                      |
| vid        | mp4                      |
| img        | jpg jpeg psd ico bmp svg |
| font       | ttf                      |
| drawio     | drawio                   |
| docs       | pdf docx pptx pptm xlsx  |
| installer  | msi exe dmg              |
| codesample | cpp go py                |


## Example

![example_1](materials/sort_example.gif)

## Configuration

* If your download folder is not named `Downloads` change DOWNLOAD_PATH value in `.env` file
* if you use windows OS, go to branch windows - `git checkout windows`