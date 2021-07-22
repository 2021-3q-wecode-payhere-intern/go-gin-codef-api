package request

import "strings"

// 파라미터 empty 체크
func IsEmptyByParam(param string) bool {
	return strings.ReplaceAll(param, " ", "") == ""
}
