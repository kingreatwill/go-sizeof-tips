Golang sizeof tips
==================

**Consider alternatives:**
- https://github.com/dominikh/go-tools/tree/master/cmd/structlayout

Web tool for interactive playing with Golang struct sizes.

Try online version [here](https://golang-sizeof.wcoder.com/).

## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing

```bash
go get github.com/kingreawill/go-sizeof-tips
cd github.com/kingreawill/go-sizeof-tips
go mod tidy
go build -o ./server
```
You may also install via simple `go get` by your own risk.


## Usage
```bash
./server -http=:7777 start
./server stop
./server restart
```

> https://github.com/tyranron/golang-sizeof.tips
