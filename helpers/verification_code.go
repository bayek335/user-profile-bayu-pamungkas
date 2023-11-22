package helpers

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func GenerateCode() string {
	verifyCode := []string{"", "", "", "", "", ""}

	for i := 0; i < 6; i++ {
		randInt := int(math.Round(rand.Float64() * 9))
		randStr := strconv.Itoa(randInt)
		verifyCode[i] = randStr
	}

	newCode := strings.Join(verifyCode, "")
	return newCode
}
