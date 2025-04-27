build:
go build -o query_search ./cmd/main.go
# CGO_CFLAGS="-I/home/weisifan/usr/local/include" \
# CGO_LDFLAGS="-L/home/weisifan/usr/local/lib -lusearch_c" \
	

run:
	go run ./cmd/main.go $(args)

clean:
	rm -f query_search
