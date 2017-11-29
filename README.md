# godice

Go version of the
[Sørensen–Dice coefficient.](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient).
It can be used to compare two strings or return the best match out of an array of strings.

**Note:** The code is part of my exercises in learning Go.

## Installation

```shell
go get github.com/imjasonmiller/godsc
```

## Usage

Compare two strings and return a score as float64

```go
package main

import (
  "github.com/imjasonmiller/godsc"
  "log"
)

func main() {
  godsc.CompareTwoStrings("gopher", "golang")
}
```

Find and sort the highest scoring matches out of array of strings score

```go
package main

import (
  "github.com/imjasonmiller/godsc"
  "log"
)

func main() {
  matches := godsc.FindBestMatch("golang", [5]string{
    "gopher", "gecko", "giraffe", "grizzly", "great dane",
  })

  // gopher has a score of 0.2
  log.Println(matches.BestMatch.text, "has a score of", matches.BestMatch.Score)

  // [ Match{ gecko, 0.0 }, ... ]
  log.Println(matches.Candidates)
}
```
