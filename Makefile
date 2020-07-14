#Go Paramters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
    
all: clean app
app: 
	$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
	 $(GOCLEAN)
	 rm -f $(BINARY_NAME)
app-run: 
	./$(BINARY_NAME)