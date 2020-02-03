NAME=projector
GO=go
DEP=dep
BIN=bin

.PHONY: default clean
default: $(BIN)/$(NAME)
clean:
	rm -rf $(BIN)

$(BIN):
	@mkdir -p $(BIN)

$(BIN)/$(NAME): main.go | $(BIN)
	$(DEP) ensure
	$(GO) build -o $(BIN)/$(NAME)
