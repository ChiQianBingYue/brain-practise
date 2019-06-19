package stringproblem

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	var reg = regexp.MustCompile(`^[\+\-]?\d+`)

	filterd, _ := strconv.Atoi(reg.FindString(strings.TrimSpace(str)))
	if filterd > math.MaxInt32 {
		filterd = math.MaxInt32
	} else if filterd < math.MinInt32 {
		filterd = math.MinInt32
	}
	return filterd
}
