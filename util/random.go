package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomSample picks a random element from array arr
func RandomPickArrayStr(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFlaot generates a random integer between min and max
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomString generates a random alphabet string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// randomName generates a random firstname or lastname of school
func randomName() string {
	return RandomString(15)
}

// randomEmail generates a random email
func randomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(20))
}
