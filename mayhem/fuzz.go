package fuzz

import "github.com/m7shapan/njson"

func mayhemit(bytes []byte) int {

    type test interface {}
    var mayhem test
    njson.Unmarshal(bytes, mayhem)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}