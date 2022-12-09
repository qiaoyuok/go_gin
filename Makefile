## run: go run
.PHONY: run
run:
	go run main.go

.PHONY: air
air:
	./air -c air.conf

## run: go gen
.PHONY: gen
gen:
	go run ./cmd/generate.go