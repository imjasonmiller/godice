# godice

Go version of the
[Sørensen–Dice coefficient](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient).
It can be used to compare two strings or return the best match out of an array of strings.

**Note:** The code is part of my exercises in learning Go.

## Installation

```shell
go get github.com/imjasonmiller/godsc
```

## Usage

#### compareString
Compare a string to one other string and return the score

```go
package main

import (
  "github.com/imjasonmiller/godice"
  "log"
)

func main() {
  godice.CompareString("gopher", "golang")
}
```
#### compareStrings
Variadic function and can take many strings. Given strings are sorted by their score.

```go
package main

import (
  "github.com/imjasonmiller/godice"
  "log"
)

func main() {
  matches := godice.CompareStrings("golang", "gopher", "gecko", "giraffe", "grizzly", "great dane" })

  // gopher has a score of 0.2
  log.Println(matches.BestMatch.text, "has a score of", matches.BestMatch.Score)

  // [ Match{ gecko, 0.0 }, ... ]
  log.Println(matches.Candidates)
}
```
