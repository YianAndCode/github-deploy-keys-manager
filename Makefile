BIN_DIR = ./bin
PROJ_NAME = github-deploy-keys-manager

all: gdkm

gdkm:
	go build -o $(BIN_DIR)/$(PROJ_NAME) .

clean:
	rm $(BIN_DIR)/$(PROJ_NAME)

.PHONY: clean