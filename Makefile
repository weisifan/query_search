build:
	go build -o query_search ./cmd/main.go

run:
	go run ./cmd/main.go $(args)

clean:
	rm -f query_search
