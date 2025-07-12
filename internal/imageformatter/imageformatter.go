package imageformatter

import (
	"fmt"
	"os"
)

func Run() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("❌ Could not get working directory: %v", err)
	}

	// 🔍 Step 1: Read available image files
	files, err := GetImageFiles(cwd)
	if err != nil {
		return fmt.Errorf("❌ Failed to read images: %v", err)
	}
	if len(files) == 0 {
		return fmt.Errorf("❌ No image files found in current directory")
	}

	// 🖼️ Step 2: Ask user to select files
	selected, err := AskImageSelection(files)
	if err != nil {
		return err
	}

	// 🎨 Step 3: Ask output format
	format, err := AskFormat()
	if err != nil {
		return err
	}

	// 🧪 Step 4: Ask for output quality
	quality, err := AskQuality()
	if err != nil {
		return err
	}

	// 🔁 Step 5: Ask if we should re-encode same formats
	reencode, err := AskReEncode(format)
	if err != nil {
		return err
	}

	// ⏱️ Step 6: Ask if time per image should be shown
	showTime, err := AskShowTime()
	if err != nil {
		return err
	}

	// 📁 Step 7: Ask about output folder
	useFolder, folderName, err := AskUseOutputFolder()
	if err != nil {
		return err
	}

	// 🧹 Step 8: Ask about deleting originals
	deleteOriginals, err := AskDeleteOriginals()
	if err != nil {
		return err
	}

	// 🚀 Step 9: Kick off the conversion
	return ConvertImages(ConvertOptions{
		Images:          selected,
		Format:          format,
		DeleteOriginals: deleteOriginals,
		UseOutputFolder: useFolder,
		ReEncode:        reencode,
		Cwd:             cwd,
		FolderName:      folderName,
		Quality:         quality,
		ShowTime:        showTime,
	})
}
