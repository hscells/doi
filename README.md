# doi

[![documentation](https://img.shields.io/badge/GoDoc-doi-blue.svg)](https://godoc.org/github.com/hscells/doi)

_dealing with dois in go_

## Usage

```go
d, err := doi.Parse("11.1038/123456")
if err != nil {
    println(d.ToString())
}
if d.IsValid() {
    println("We are happy!")
}
```

