package helper

import (
	"encoding/json"
	"bytes"
	"jindouyun/jdyError"
	"fmt"
)

func FormatOutPut(re []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, re, "", "   ")
	jdyError.CheckError(err, false)
	fmt.Printf("%+v\n",out.String())
}