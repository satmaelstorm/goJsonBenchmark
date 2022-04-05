.PHONY: bench

bench:
	go test -bench=. -benchmem

gen:
	go generate ./...
