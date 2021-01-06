SOURCE := main.go
BINARY := main.wasm

all: $(BINARY)

$(BINARY): $(SOURCE)
	GOOS=js GOARCH=wasm go build -o $@

clean: 
	@$(RM) $(BINARY)

.PHONY: all clean
