BIN_DIR = ./bin
PROJ_NAME = gdkm

all: gdkm

gdkm:
	go build -o $(BIN_DIR)/$(PROJ_NAME) .

clean:
	rm $(BIN_DIR)/$(PROJ_NAME)

.PHONY: clean