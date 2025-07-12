package imageformatter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/h2non/bimg"
)

type ConvertOptions struct {
	Images          []string
	Format          string
	DeleteOriginals bool
	UseOutputFolder bool
	ReEncode        bool
	Cwd             string
	FolderName      string
	Quality         int
	ShowTime        bool
}

func convertFormatToBimgType(format string) bimg.ImageType {
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		return bimg.JPEG
	case "png":
		return bimg.PNG
	case "webp":
		return bimg.WEBP
	default:
		return bimg.UNKNOWN
	}
}

func cleanFileName(name, ext string) string {
	base := strings.TrimSuffix(name, filepath.Ext(name))
	base = strings.TrimSuffix(base, ".cli-pix")
	return base + ".cli-pix." + ext
}

func ConvertImages(opts ConvertOptions) error {
	outputDir := filepath.Join(opts.Cwd, opts.FolderName)
	if opts.UseOutputFolder {
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			if err := os.Mkdir(outputDir, 0755); err != nil {
				return fmt.Errorf("failed to create output folder: %v", err)
			}
		}
	}

	// Validate & sanitize quality value
	quality := opts.Quality
	if quality <= 0 || quality > 100 {
		quality = 80
	}

	var (
		wg             sync.WaitGroup
		mu             sync.Mutex
		convertedCount int
		skippedCount   int
		deletedCount   int
	)

	sem := make(chan struct{}, 16) // concurrency limit

	for _, image := range opts.Images {
		wg.Add(1)
		go func(image string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			start := time.Now()
			inputPath := filepath.Join(opts.Cwd, image)

			buffer, err := os.ReadFile(inputPath)
			if err != nil {
				fmt.Printf("❌ Failed to read %s\n", image)
				return
			}

			img := bimg.NewImage(buffer)
			meta, err := img.Metadata()
			if err != nil {
				fmt.Printf("❌ Could not read metadata for %s\n", image)
				return
			}

			currentFormat := strings.ToLower(meta.Type)
			targetFormat := strings.ToLower(opts.Format)

			if !opts.ReEncode && currentFormat == targetFormat {
				mu.Lock()
				skippedCount++
				mu.Unlock()
				fmt.Printf("⏭️  Skipped %s — already in .%s\n", image, currentFormat)
				return
			}

			newName := cleanFileName(image, targetFormat)
			outputPath := filepath.Join(opts.Cwd, newName)
			if opts.UseOutputFolder {
				outputPath = filepath.Join(outputDir, newName)
			}

			options := bimg.Options{
				Type:    convertFormatToBimgType(targetFormat),
				Quality: quality,
			}

			newImage, err := img.Process(options)
			if err != nil {
				fmt.Printf("❌ Conversion failed for %s\n", image)
				return
			}

			if err := os.WriteFile(outputPath, newImage, 0644); err != nil {
				fmt.Printf("❌ Failed to save %s\n", newName)
				return
			}

			if opts.DeleteOriginals {
				if err := os.Remove(inputPath); err == nil {
					mu.Lock()
					deletedCount++
					mu.Unlock()
				}
			}

			duration := time.Since(start).Seconds()
			mu.Lock()
			convertedCount++
			mu.Unlock()

			if opts.ShowTime {
				fmt.Printf("✔ Converted: %s → %s (%.2fs)\n", image, filepath.Base(outputPath), duration)
			} else {
				fmt.Printf("✔ Converted: %s → %s\n", image, filepath.Base(outputPath))
			}
		}(image)
	}

	wg.Wait()

	// 🧾 Summary
	fmt.Println("\n📦 Summary:")
	fmt.Printf("   ✔ Converted: %d\n", convertedCount)
	fmt.Printf("   ⏭️  Skipped:   %d\n", skippedCount)
	fmt.Printf("   🧪 Quality:   %d\n", quality)
	if opts.DeleteOriginals {
		fmt.Printf("   🧹 Deleted:   %d\n", deletedCount)
	}
	if opts.UseOutputFolder {
		fmt.Printf("   📂 Output:    %s\n", outputDir)
	}

	return nil
}
