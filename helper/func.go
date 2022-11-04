package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/donkeylichao/jindouyun/jdyError"
)

func FormatOutPut(re []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, re, "", "   ")
	jdyError.CheckError(err, false)
	fmt.Printf("%+v\n",out.String())
}