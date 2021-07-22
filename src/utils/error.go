package utils

import (
	"strconv"
	"strings"
)

// codef 관련 에러메세지 생성
func GenerateCodefErrorMsg(errorMsg string) string {
	return "codef_error : " + errorMsg
}

// service 관련 에러메세지 생성
func GenerateErrorMsg(fileName string, line int) string {
	slice := strings.Split(fileName, "/")
	sLine := strconv.Itoa(line)

	msg := "error : " + slice[len(slice)-1] + ", " + sLine + "line"
	return msg
}
