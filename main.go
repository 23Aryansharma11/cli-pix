package main

import (
	"flag"
	"log"

	"cli-pix/internal/imageformatter"
)

func main() {

	quality := flag.Int("quality", 80, "Image quality (1–100)")
	defaults := flag.Bool("defaults", false, "Use default options without prompting")
	format := flag.String("format", "", "Image format (webp, png, jpeg)")

	flag.Parse()

	// Pass flags to Run()
	if err := imageformatter.RunWithConfig(*quality, *defaults, *format); err != nil {
		log.Fatal("❌", err)
	}
}
