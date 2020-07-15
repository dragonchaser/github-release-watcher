all:
	go build .

clean:
	rm github-release-watcher

tests:
	go test
