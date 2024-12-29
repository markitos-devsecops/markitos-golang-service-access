package domain

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func RandomString(n ...int) string {
	length := 10
	if len(n) > 0 {
		length = n[0]
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(), RandomString())
}

func IsUUIDv4(uuid string) bool {
	uuidRegex := `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	matched, err := regexp.MatchString(uuidRegex, uuid)

	return err == nil && matched
}

func UUIDv4() string {
	var uuid [16]byte

	for i := range uuid {
		uuid[i] = byte(rand.Intn(256))
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func PersonalName() string {
	minWords, maxWords := 1, 6
	wordCount := rand.Intn(maxWords-minWords+1) + minWords
	minLength, maxLength := 3, 150

	var result []string
	currentLength := 0

	for currentLength < minLength || currentLength+wordCount-1 > maxLength {
		result = nil
		currentLength = 0

		for i := 0; i < wordCount; i++ {
			word := RandomString(rand.Intn(8) + 3)
			word = strings.ToLower(word)
			result = append(result, word)
			currentLength += len(word)
		}
	}

	return strings.Join(result, " ")
}

func Slug() string {
	minLength, maxLength := 3, 75

	var result []string
	currentLength := 0

	for currentLength < minLength || currentLength+len(result)-1 > maxLength {
		result = nil
		currentLength = 0

		wordCount := rand.Intn(6) + 1
		for i := 0; i < wordCount; i++ {
			word := RandomString(rand.Intn(8) + 3)
			word = strings.ToLower(word)
			result = append(result, word)
			currentLength += len(word)
		}
	}

	return strings.Join(result, "-")
}
