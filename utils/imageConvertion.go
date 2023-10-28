package imageConvertion

import (
	"errors"
	"log"
	"path"
	"strings"

	"github.com/h2non/bimg"
)

var SupportedImageTypeMapping = map[bimg.ImageType]string{
	bimg.JPEG: "jpeg",
	//bimg.WEBP: "webp",
	bimg.PNG: "png",
	//bimg.TIFF: "tiff",
	//bimg.GIF:  "gif",
	//bimg.PDF:  "pdf",
	//bimg.SVG: "svg",
	// HEIF represents the HEIC/HEIF/HVEC image type
	bimg.HEIF: "heic",
	bimg.AVIF: "avif",
}

func ImageTypeToStr(imageType bimg.ImageType) string {

	elem, present := SupportedImageTypeMapping[imageType]
	if present {
		return elem
	} else {
		log.Fatal("Image format not supported")
		return ""
	}
}

func ConvertImage(input_image_path string, optionnal_output_image_path string, new_format bimg.ImageType) (string, error) {
	old_format_str := path.Ext(input_image_path)
	new_format_str := ImageTypeToStr(new_format)

	output_path := strings.Replace(input_image_path, old_format_str, "."+new_format_str, 1)
	if optionnal_output_image_path != "" {
		output_path = optionnal_output_image_path
	}

	buffer, err := bimg.Read(input_image_path)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Failed to read input image")
	}

	newImage, err := bimg.NewImage(buffer).Convert(new_format)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Failed to convert image read")
	}

	log.Printf("The image was converted into " + new_format_str + " at " + output_path)
	bimg.Write(output_path, newImage)
	return output_path, nil

}
