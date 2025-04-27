package internal

import (
	"github.com/unum-cloud/usearch/golang"
)

func BuildIndex(embeddings [][]float32) (*usearch.Index, error) {
	conf := usearch.DefaultConfig(uint(len(embeddings[0]))) // dimension from embedding
	index, err := usearch.NewIndex(conf)
	if err != nil {
		return nil, err
	}

	err = index.Reserve(uint(len(embeddings)))
	if err != nil {
		return nil, err
	}

	for i, emb := range embeddings {
		err := index.Add(usearch.Key(i), emb)
		if err != nil {
			return nil, err
		}
	}

	return index, nil
}
