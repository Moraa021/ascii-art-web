package ascii

import (
	"os"
	"strings"
)

// GenerateASCII handles character conversions matching the precise audit templates.
func GenerateASCII(text string, banner string) (string, int) {
	// Try standard root path first
	filePath := "banners/" + banner + ".txt"
	
	// If it doesn't exist, check parent directory (fixes 'go test' context)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = "../banners/" + banner + ".txt"
	}
	
	// Read full file bytes
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", 404 // 404 Not Found if banner template is missing
	}
	
	// ... (Keep the exact rest of your ascii.go code completely identical) ...
	content := strings.ReplaceAll(string(contentBytes), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	if len(lines) < 855 {
		return "", 500 
	}

	text = strings.ReplaceAll(text, "\r\n", "\n")
	
	if text == "" {
		return "", 200
	}
	
	words := strings.Split(text, "\n")
	allEmpty := true
	for _, word := range words {
		if word != "" {
			allEmpty = false
			break
		}
	}
	if allEmpty && len(words) > 1 {
		return strings.Repeat("\n", len(words)-1), 200
	}

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					return "", 400 
				}
				
				startIdx := int(char-32)*9 + 1 + i
				if startIdx < len(lines) {
					result.WriteString(lines[startIdx])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), 200
}