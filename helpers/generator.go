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

func GenerateRandomString(length int) string {
	strs := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "W", "X", "Y", "Z"}

	resSlice := []string{}
	for i := 0; i < length; i++ {
		ranInt := rand.Intn(49)
		resSlice = append(resSlice, strs[ranInt])
	}

	resStr := strings.Join(resSlice, "")

	return resStr
}
