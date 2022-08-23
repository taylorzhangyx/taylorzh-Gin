package async_task

import (
	"fmt"
	"strings"
)

func echo(voice ...string) (string, error) {
	if len(voice) == 0 {
		return "", fmt.Errorf("empty voice error: you may forget to speak")
	}
	return strings.Join(voice, ","), nil
}

func add(nums ...int64) (int64, error) {
	if len(nums) == 0 {
		return 0, nil
	}
	sum := int64(0)
	for _, n := range nums {
		sum += n
	}
	return sum, nil
}
