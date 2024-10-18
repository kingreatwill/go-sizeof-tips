Golang sizeof tips
==================

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


## Other

- https://github.com/dominikh/go-tools/tree/master/cmd/structlayout

