# godice

Go version of the
[Sørensen–Dice coefficient](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient).
It can be used to compare two strings or return the best match in a slice of strings.

**Note:** The code was part of an exercise in learning Go. Pull requests are welcome.

## Install

```shell
go get github.com/imjasonmiller/godice
```

## Usage

#### godice.compareString(a, b str) float64
Compare string `a` to string `b` and return the score.

```go
godice.CompareString("gopher", "golang") 
```

The above would return `0.2`.

#### godice.compareStrings(a str, b []str) Matches
Compare string `a` to a `slice` of strings. Strings are sorted by their score.

```go
godice.CompareStrings("golang", []string{"gopher", "gerbil", "giraffe"})
```

The above would return:

```go
Matches{
  BestMatch:{ Text: "gopher", Score: 0.2 },
  Candidates:[
    { Text: "gopher", Score: 0.2 },
    { Text: "gerbil", Score: 0.0 },
    { Text: "grison", Score: 0.0 },
  ],
}
```
