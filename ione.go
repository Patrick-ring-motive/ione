package ione

import (
	"fmt"
	. "github.com/Patrick-ring-motive/traigo"
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

func IoReadAll(ioReadCloser *IoReadCloser, options ...IoOptions) []byte {
	ioNoNil := true
	ioNoError := true
	ioNoPanic := true
	if len(options) > 0 {
		ioNoNil = options[0].IoNoNil
		ioNoError = options[0].IoNoError
		ioNoPanic = options[0].IoNoPanic
	}
	var bytes []byte
	Try(func() {
		bytes, err := io.ReadAll(*ioReadCloser.Value)
		if (bytes == nil) && (err == nil) && ioNoNil {
			fmt.Println("IoReadAll nil")
			bytes = []byte("IoReadAll nil")
		}
		if (err != nil) && ioNoError {
			fmt.Println("IoReadAll error: ", err.Error())
			bytes = []byte(fmt.Sprint("IoReadAll error: ", err.Error()))
		}
	}, Catch(func(e interface{}) {
		if ioNoPanic {
			fmt.Println("IoReadAll panic: ", e)
			bytes = []byte(fmt.Sprint("IoReadAll panic: ", e))
		} else {
			panic(e)
		}
	}))
	return bytes
}
