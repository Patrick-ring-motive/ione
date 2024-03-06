package ione

import (
	"fmt"
	utils "github.com/Patrick-ring-motive/utils"
	"io"
)

var IoNoNil = true
var IoNoError = true
var IoNoPanic = true

func let() { utils.AllowUnused([5]any{fmt.Print, io.Copy, IoNoNil, IoNoError, IoNoPanic}) }


type IoReadCloser struct{
  Value *io.ReadCloser
}



func IoReadAll(ioReadCloser IoReadCloser) []byte{
  bytes  := []byte(nil)
  if IoNoPanic {
    defer func() {
      if r := recover(); r != nil {
        fmt.Println("IoReadAll panic: ", r)
        bytes = []byte(fmt.Sprint("IoReadAll panic: ", r))
      }
    }()
  }
  bytes, err := io.ReadAll(*ioReadCloser.Value)
  if (bytes == nil) && (err == nil) && IoNoNil {
    fmt.Println("IoReadAll nil")
    bytes = []byte("IoReadAll nil")
  }
  if (err != nil) && IoNoError {
    fmt.Println("IoReadAll error: ", err.Error())
     bytes = []byte(fmt.Sprint("IoReadAll error: ", err.Error()))
  }
  return bytes
}