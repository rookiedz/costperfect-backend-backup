package convert

import (
	"fmt"
	"strings"
)

//ArrayInt64ToString ...
func ArrayInt64ToString(arr []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}
