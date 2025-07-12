package imageformatter

import (
	"fmt"
	"os"
	"strings"
)

func RunWithConfig(quality int, defaults bool, formatFlag string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("❌ Could not get working directory: %v", err)
	}

	files, err := GetImageFiles(cwd)
	if err != nil {
		return fmt.Errorf("❌ Failed to read images: %v", err)
	}
	if len(files) == 0 {
		return fmt.Errorf("❌ No image files found in current directory")
	}

	var selected []string
	if defaults {
		selected = files
	} else {
		selected, err = AskImageSelection(files)
		if err != nil {
			return err
		}
	}

	var format string
	if formatFlag != "" {
		formatFlag = strings.ToLower(formatFlag)
		if formatFlag != "webp" && formatFlag != "png" && formatFlag != "jpeg" {
			return fmt.Errorf("❌ Invalid format '%s'. Supported formats: webp, png, jpeg", formatFlag)
		}
		format = formatFlag
	} else {
		format, err = AskFormat()
		if err != nil {
			return err
		}
	}

	var reencode bool
	if defaults {
		reencode = false
	} else {
		reencode, err = AskReEncode(format)
		if err != nil {
			return err
		}
	}

	var useFolder bool
	var folderName string
	if defaults {
		useFolder = true
		folderName = "cli-pix"
	} else {
		useFolder, folderName, err = AskUseOutputFolder()
		if err != nil {
			return err
		}
	}

	var deleteOriginals bool
	if defaults {
		deleteOriginals = false
	} else {
		deleteOriginals, err = AskDeleteOriginals()
		if err != nil {
			return err
		}
	}

	return ConvertImages(ConvertOptions{
		Images:          selected,
		Format:          format,
		DeleteOriginals: deleteOriginals,
		UseOutputFolder: useFolder,
		ReEncode:        reencode,
		Cwd:             cwd,
		FolderName:      folderName,
		Quality:         quality,
		ShowTime:        true,
	})
}


