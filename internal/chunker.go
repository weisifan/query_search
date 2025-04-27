package internal

import (
	"fmt"
	"os"
	"strings"
)

var commonAbbreviations = []string{
	"Mr.", "Mrs.", "Ms.", "Dr.", "St.", "Jr.", "Sr.", "e.g.", "i.e.", "etc.", "vs.",
}

// isAbbreviation checks if the given word matches a common abbreviation
func isAbbreviation(word string) bool {
	for _, abbr := range commonAbbreviations {
		if word == abbr {
			return true
		}
	}
	return false
}

func ChunkText(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	text := string(data)
	words := strings.Fields(text)

	var sentences []string
	var currentSentence strings.Builder

	for _, word := range words {
		currentSentence.WriteString(word + " ")

		// Check if the word ends a sentence
		if strings.HasSuffix(word, ".") || strings.HasSuffix(word, "!") || strings.HasSuffix(word, "?") {
			if !isAbbreviation(word) {
				sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
				currentSentence.Reset()
			}
		}
	}

	// Add any leftover sentence
	if currentSentence.Len() > 0 {
		sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
	}

	// Now pack sentences into chunks
	var chunks []string
	var currentChunk strings.Builder
	maxChunkSize := 200

	for _, sentence := range sentences {
		if currentChunk.Len()+len(sentence) < maxChunkSize {
			currentChunk.WriteString(sentence + " ")
		} else {
			chunks = append(chunks, strings.TrimSpace(currentChunk.String()))
			currentChunk.Reset()
			currentChunk.WriteString(sentence + " ")
		}
	}

	if currentChunk.Len() > 0 {
		chunks = append(chunks, strings.TrimSpace(currentChunk.String()))
	}

	fmt.Println("Total chunks:", len(chunks))
	return chunks, nil
}

