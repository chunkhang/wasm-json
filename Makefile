SOURCE := main.go
BINARY := main.wasm

all: $(BINARY)

install:
	@npm install

serve:
	@./node_modules/.bin/serve

$(BINARY): $(SOURCE)
	GOOS=js GOARCH=wasm go build -o $@

clean: 
	@$(RM) $(BINARY)

uninstall:
	@$(RM) -rf node_modules

.PHONY: all install serve clean uninstall
