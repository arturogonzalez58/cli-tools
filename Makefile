wordCounter:
	go build -o bin/wc cmd/wordCounter/main.go

clean:
	rm bin/*