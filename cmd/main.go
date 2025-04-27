package main

import (
	"fmt"
	"log"
	"os"

	"query_search/internal"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: query_search <OPENAI_API_KEY> <PATH_TO_TEXT_FILE>")
		os.Exit(1)
	}

	apiKey := os.Args[1]
	filePath := os.Args[2]

	chunks, err := internal.ChunkText(filePath)
	if err != nil {
		log.Fatal(err)
	}

	embeddings, err := internal.EmbedChunks(apiKey, chunks)
	if err != nil {
		log.Fatal(err)
	}

	idx, err := internal.BuildIndex(embeddings)
	if err != nil {
		log.Fatal(err)
	}

	internal.StartREPL(apiKey, idx, chunks)
}
