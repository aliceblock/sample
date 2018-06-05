# Sample Golang Project
## Install
1. Clone this repo
```bash
git clone https://github.com/aliceblock/sample
```

---

## Usage

1. Get livereload
```bash
go get github.com/codegangsta/gin
```
2. Get go-assets-builder for pre-build html assets
```bash
go get github.com/jessevdk/go-assets-builder
go-assets-builder web/templates -o cmd/server/assets.go
```
3. Run with
```bash
gin --path ./cmd/server --port 8080 main.go
```

---

## Framework
- Gin [ [Here](https://github.com/gin-gonic/gin) ]

---

## Optional
- Dep [ [Here](https://github.com/golang/dep) ]
- Gomove to rename package name [ [Here](https://github.com/KSubedi/gomove) ]