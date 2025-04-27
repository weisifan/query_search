// package internal

// import (
// 	"fmt"
// 	"os"
// )

// func ChunkText(filePath string) ([]string, error) {
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	text := string(data)
// 	var chunks []string
// 	chunkSize := 200 // you can tune this later

// 	for i := 0; i < len(text); i += chunkSize {
// 		end := i + chunkSize
// 		if end > len(text) {
// 			end = len(text)
// 		}
// 		chunks = append(chunks, text[i:end])
// 	}
// 	fmt.Println("Number of chunks:", len(chunks))
// 	for i, chunk := range chunks {
// 		fmt.Printf("Chunk %d length: %d\n", i, len(chunk))
// 	}
// 	return chunks, nil
// }


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

