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

// ChunkText reads a file, splits text into sentences, and groups sentences into meaningful chunks.
func ChunkText(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	text := string(data)
	words := strings.Fields(text)

	var sentences []string
	var currentSentence strings.Builder

	// Step 1: Split into sentences
	for _, word := range words {
		currentSentence.WriteString(word + " ")
		if strings.HasSuffix(word, ".") || strings.HasSuffix(word, "!") || strings.HasSuffix(word, "?") {
			if !isAbbreviation(word) {
				sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
				currentSentence.Reset()
			}
		}
	}
	if currentSentence.Len() > 0 {
		sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
	}

	// Step 2: Group sentences into chunks
	var chunks []string
	maxSentencesPerChunk := 5
	maxCharactersPerChunk := 500

	currentChunkSentences := []string{}
	currentChunkSize := 0

	for i, sentence := range sentences {
		sentenceLength := len(sentence)
		currentChunkSentences = append(currentChunkSentences, sentence)
		currentChunkSize += sentenceLength

		// Conditions to close current chunk:
		if len(currentChunkSentences) >= maxSentencesPerChunk || currentChunkSize >= maxCharactersPerChunk || i == len(sentences)-1 {
			chunk := strings.Join(currentChunkSentences, " ")
			chunks = append(chunks, chunk)
			currentChunkSentences = []string{}
			currentChunkSize = 0
		}
	}

	// Print all chunks clearly
	fmt.Println("========== Chunks ==========")
	for idx, chunk := range chunks {
		fmt.Printf("\n--- Chunk %d ---\n", idx+1)
		fmt.Println(chunk)
	}
	fmt.Printf("\nTotal chunks: %d\n", len(chunks))
	fmt.Println("=============================")

	return chunks, nil
}
