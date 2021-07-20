package utils

import "strconv"

// 월의 마지막일자 구하기
func GetLastDay(year, month string) string {
	lastDayMap := map[int]string{
		1: "31", 2: "28",
		3: "31", 4: "30",
		5: "31", 6: "30",
		7: "31", 8: "31",
		9: "30", 10: "31",
		11: "30", 12: "31",
	}

	tempYear, _ := strconv.Atoi(year)
	tempMonth, _ := strconv.Atoi(month)

	if tempMonth == 2 && ((tempYear%4 == 0 && tempYear%100 != 0) || tempYear%400 == 0) {
		lastDayMap[tempMonth] = "29"
	}

	return lastDayMap[tempMonth]
}

// 년 월 일 string값 합치기
func GetConcatDate(year, month, day string) string {
	return year + month + day
}
