package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func EmbedChunks(apiKey string, chunks []string) ([][]float32, error) {
	var embeddings [][]float32
	for _, chunk := range chunks {
		emb, err := GetEmbedding(apiKey, chunk)
		if err != nil {
			return nil, err
		}
		embeddings = append(embeddings, emb)
	}
	return embeddings, nil
}

func GetEmbedding(apiKey, input string) ([]float32, error) {
    url := "https://api.openai.com/v1/embeddings"

    body, _ := json.Marshal(map[string]interface{}{
        "input": input,
        "model": "text-embedding-ada-002",
    })

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Debug print
    bodyBytes, _ := io.ReadAll(resp.Body)
    // fmt.Println("Raw OpenAI response:", string(bodyBytes))

    var result struct {
        Data []struct {
            Embedding []float32 `json:"embedding"`
        } `json:"data"`
    }

    if err := json.Unmarshal(bodyBytes, &result); err != nil {
        return nil, err
    }

    if len(result.Data) == 0 {
        return nil, fmt.Errorf("no embedding returned")
    }

    return result.Data[0].Embedding, nil
}