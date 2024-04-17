package ione

import (
	"fmt"
//	. "github.com/Patrick-ring-motive/traigo"
	"io"
)

type IoOptions struct {
	IoNoNil   bool
	IoNoError bool
	IoNoPanic bool
}

type IoReadCloser struct {
	Value *io.ReadCloser
}

func IoReadAll(ioReadCloser IoReadCloser, options ...IoOptions) []byte {

	var bytes []byte
  var b [1][]byte = [1][]byte{}
  ioReadAll(b, ioReadCloser, options...)
  bytes = b[0]
	return bytes
}

func ioReadAll(b [1][]byte, ioReadCloser IoReadCloser, options ...IoOptions){
  ioNoNil := true
  ioNoError := true
  ioNoPanic := true
  if len(options) > 0 {
    ioNoNil = options[0].IoNoNil
    ioNoError = options[0].IoNoError
  	ioNoPanic = options[0].IoNoPanic
  }
  if ioNoPanic {
    defer func() {
      if r := recover(); r != nil {
        b[0] = []byte(fmt.Sprint("IoReadAll panic: ", r))
      }
    }()
  }
  bytes, err := io.ReadAll(*ioReadCloser.Value)
  if (bytes == nil) && (err == nil) && ioNoNil {
    fmt.Println("IoReadAll nil")
    bytes = []byte("IoReadAll nil")
  }
  if (err != nil) && ioNoError {
    fmt.Println("IoReadAll error: ", err.Error())
    bytes = []byte(fmt.Sprint("IoReadAll error: ", err.Error()))
  }
  b[0]=bytes
}