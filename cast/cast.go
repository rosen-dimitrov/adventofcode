package cast

//lib for casting types

import (
	"fmt"
	"strconv"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}
