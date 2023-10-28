package main

import (
	"log"
	"path"
	"strings"

	"github.com/h2non/bimg"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

func chooseFilesToConvert() []string {
	results, err := cfdutil.ShowOpenMultipleFilesDialog(cfd.DialogConfig{
		Title: "Open Multiple Files",
		Role:  "OpenFilesExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 2,
		FileName:                "file.txt",
		DefaultExtension:        "txt",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Printf("Chosen file(s): %s\n", results)
	return results
}

func imageTypeToStr(imageType bimg.ImageType) string {
	imageTypeMapping := map[bimg.ImageType]string{
		bimg.JPEG: "jpeg",
		bimg.WEBP: "webp",
		bimg.PNG:  "png",
		bimg.TIFF: "tiff",
		bimg.GIF:  "gif",
		bimg.PDF:  "pdf",
		bimg.SVG:  "svg",
		// HEIF represents the HEIC/HEIF/HVEC image type
		bimg.HEIF: "heic",
		bimg.AVIF: "avif",
	}
	elem, present := imageTypeMapping[imageType]
	if present {
		return elem
	} else {
		log.Fatal("Image format not supported")
		return ""
	}
}

func convertImage(image_path string, new_format bimg.ImageType) {
	old_format_str := path.Ext(image_path)
	new_format_str := "." + imageTypeToStr(new_format)
	output_path := strings.Replace(image_path, old_format_str, new_format_str, 1)

	buffer, err := bimg.Read(image_path)
	if err != nil {
		log.Fatal(err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(new_format)
	if err != nil {
		log.Fatal(err)
	}

	if bimg.NewImage(newImage).Type() == "png" {
		log.Printf("The image was converted into" + new_format_str + " at " + output_path)
		bimg.Write(output_path, newImage)
	}
}

func main() {
	new_format := bimg.PNG

	inputs := chooseFilesToConvert()
	for _, input_path := range inputs {
		convertImage(input_path, new_format)
	}

}
