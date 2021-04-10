PACKAGE := github.com/mikemadden42/stats 
OUTFILE := stats

stats: main.go
	GOOS=darwin GOARCH=amd64 go build -o $(OUTFILE)-amd64 $(PACKAGE)
	GOOS=darwin GOARCH=arm64 go build -o $(OUTFILE)-arm64 $(PACKAGE)
	lipo -create -output $(OUTFILE) $(OUTFILE)-amd64 $(OUTFILE)-arm64
	codesign -s - $(OUTFILE)
	rm -f $(OUTFILE)-amd64 $(OUTFILE)-arm64

run: stats
	@./$(OUTFILE)
