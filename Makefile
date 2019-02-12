

.PHONY: data
data:
	go get -u github.com/a-urth/go-bindata
	go-bindata -pkg data -o data/data.go ./data/...
