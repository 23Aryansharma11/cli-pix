package imageformatter

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskFormat() (string, error) {
	var format string
	prompt := &survey.Select{
		Message: "🎨 Choose target format:",
		Options: []string{"webp", "png", "jpeg"},
		Default: "webp",
	}
	err := survey.AskOne(prompt, &format)
	return format, err
}

func AskImageSelection(files []string) ([]string, error) {
	var selected []string
	choices := append([]string{"All"}, files...)

	prompt := &survey.MultiSelect{
		Message: "🖼️  Select images to convert:",
		Options: choices,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return nil, err
	}

	// If "All" is selected, return all files (excluding "All")
	for _, sel := range selected {
		if sel == "All" {
			return files, nil
		}
	}

	return selected, nil
}

func AskReEncode(format string) (bool, error) {
	var reencode bool
	prompt := &survey.Confirm{
		Message: "Re-encode files already in ." + format + "?",
		Default: false,
	}
	err := survey.AskOne(prompt, &reencode)
	return reencode, err
}

func AskUseOutputFolder() (bool, string, error) {
	var useFolder bool
	var folderName string

	err := survey.AskOne(&survey.Confirm{
		Message: "Place output images in a separate folder?",
		Default: true,
	}, &useFolder)
	if err != nil {
		return false, "", err
	}

	if useFolder {
		err = survey.AskOne(&survey.Input{
			Message: "Enter output folder name:",
			Default: "cli-pix",
		}, &folderName)
		if err != nil {
			return false, "", err
		}
	}

	return useFolder, folderName, nil
}

func AskDeleteOriginals() (bool, error) {
	var deleteOriginals bool
	err := survey.AskOne(&survey.Confirm{
		Message: "Delete original images after conversion?",
		Default: false,
	}, &deleteOriginals)
	return deleteOriginals, err
}

func AskQuality() (int, error) {
	var quality int
	prompt := &survey.Input{
		Message: "📉 Enter output quality (1–100):",
		Default: "80",
	}
	err := survey.AskOne(prompt, &quality, survey.WithValidator(survey.Required))
	if err != nil {
		return 0, err
	}

	if quality < 1 || quality > 100 {
		quality = 80 // fallback to safe default
	}

	return quality, nil
}

func AskShowTime() (bool, error) {
	var show bool
	prompt := &survey.Confirm{
		Message: "⏱️  Show time taken per image?",
		Default: true,
	}
	err := survey.AskOne(prompt, &show)
	return show, err
}
