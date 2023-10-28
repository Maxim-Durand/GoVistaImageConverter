package test

import (
	"testing"

	"golang.org/x/exp/maps"

	imageConvertion "github.com/Maxim-Durand/GoVistaImageConverter/utils"
)

func TestHeicToAllSupportedImageTypeConvertion(t *testing.T) {
	test_heic_image_path := "data/inputs/valid_example.heic"
	supportedImageType := maps.Keys(imageConvertion.SupportedImageTypeMapping)
	for _, format := range supportedImageType {
		format_str := imageConvertion.ImageTypeToStr(format)
		result, err := imageConvertion.ConvertImage(test_heic_image_path, "data/outputs/heic_converted_to_"+format_str+"."+format_str, format)
		if result == "" || err != nil {
			t.Fatalf(`TestHeicToAllSupportedImageTypeConvertion failed : %v`, err)
		}
	}

}
