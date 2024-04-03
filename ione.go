package ione

import (
	"fmt"
//	utils "github.com/Patrick-ring-motive/utils"
  . "github.com/Patrick-ring-motive/traigo"
	"io"
)

var IoNoNil = true
var IoNoError = true
var IoNoPanic = true


type IoReadCloser struct{
  Value *io.ReadCloser
}

func IoReadAll(ioReadCloser ...IoReadCloser) []byte{
  var bytes []byte

  Try(func(){
    if len(ioReadCloser) == 0 {
      fmt.Println("IoReadAll called with no arguments")
      bytes = []byte("IoReadAll called with no arguments")
      return
    }
    bytes, err := io.ReadAll(*ioReadCloser[0].Value)
    if (bytes == nil) && (err == nil) && IoNoNil {
      fmt.Println("IoReadAll nil")
      bytes = []byte("IoReadAll nil")
    }
    if (err != nil) && IoNoError {
      fmt.Println("IoReadAll error: ", err.Error())
       bytes = []byte(fmt.Sprint("IoReadAll error: ", err.Error()))
    }
  }, Catch(func(e interface{}){
    if IoNoPanic {
        fmt.Println("IoReadAll panic: ", e)
        bytes = []byte(fmt.Sprint("IoReadAll panic: ", e))
    }else{
        panic(e)
    }
  }))

  return bytes
}