package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/unum-cloud/usearch/golang"
)

func StartREPL(apiKey string, idx *usearch.Index, chunks []string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter search query (or type 'quit'): ")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		if query == "quit" {
			break
		}

		queryEmb, err := GetEmbedding(apiKey, query)
		if err != nil {
			fmt.Println("Failed to embed query:", err)
			continue
		}

		keys, distances, err := idx.Search(queryEmb, 3)
		if err != nil {
			fmt.Println("Search failed:", err)
			continue
		}

		fmt.Println("Top Results:")
		for i, key := range keys {
			if int(key) < len(chunks) {
				fmt.Printf("Chunk #%d (distance %.4f):\n%s\n\n", key, distances[i], chunks[key])
			}
		}
	}
}
