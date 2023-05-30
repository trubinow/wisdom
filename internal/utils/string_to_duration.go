package utils

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func StringToDuration(durationStr string) (time.Duration, error) {
	var re = regexp.MustCompile(`(?m)^([\d]{1,3})(m|s)$`)
	var result time.Duration

	match := re.FindStringSubmatch(durationStr)
	if len(match) < 3 {
		return result, errors.New("invalid duration string")
	}

	multiplier, err := strconv.Atoi(match[1])
	if err != nil {
		return result, err
	}

	switch match[2] {
	case "m":
		result = time.Duration(multiplier) * time.Minute
	case "s":
		result = time.Duration(multiplier) * time.Second
	}

	return result, nil
}
