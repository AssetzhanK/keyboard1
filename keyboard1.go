package keyboard1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Чтение текста и конвертация в float64 21:50
func GetFloat1() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	number += 10
	return number, nil
}
