package main

import (
	"log"

	imageConvertion "github.com/Maxim-Durand/GoVistaImageConverter/utils"

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

func main() {
	new_format := bimg.PNG

	//inputs := chooseFilesToConvert()
	inputs := [1]string{"/home/maximedurand/Documents/GoVistaImageConverter/test/data/valid_example.heic"}
	for _, input_path := range inputs {
		result, err := imageConvertion.ConvertImage(input_path, "", new_format)
		if result != "" {
			log.Printf("Successful convertion of %v", input_path)
		} else {
			if err != nil {
				log.Fatalf("Failed to convert %v", input_path)
			}
		}
	}

}
