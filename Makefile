wordCounter:
	go build -o bin/wc cmd/wordCounter.go

clean:
	rm bin/*