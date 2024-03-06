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
