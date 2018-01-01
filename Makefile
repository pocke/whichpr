dep:
	go get github.com/goreleaser/goreleaser

release:
	goreleaser

snapshot:
	goreleaser --snapshot

clean:
	rm -r dist
