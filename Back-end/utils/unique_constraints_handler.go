package utils

import (
	
	"regexp"

    "strings"

)

func ExtractColumn(err string) string {
	r, _ := regexp.Compile(`uni_[^"]+`)
	result :=r.FindString(err)
	match:=strings.Split(result,"_")
	match= match[2:]
	result=strings.Join(match, " ")
	return result
}