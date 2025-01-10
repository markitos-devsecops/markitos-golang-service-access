package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

func RandomString(howManyLetters ...int) string {
	length := 10
	if len(howManyLetters) > 0 {
		length = howManyLetters[0]
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)

	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
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
	rand.Read(uuid[:])

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func RandomPersonName() string {
	minWords, maxWords := 1, 6
	wordCount, _ := rand.Int(rand.Reader, big.NewInt(int64(maxWords-minWords+1)))
	wordCount = big.NewInt(wordCount.Int64() + int64(minWords))
	minLength, maxLength := 3, 150

	var result []string
	currentLength := 0

	for currentLength < minLength || currentLength+int(wordCount.Int64())-1 > maxLength {
		result = nil
		currentLength = 0
		for i := 0; i < int(wordCount.Int64()); i++ {
			num, _ := rand.Int(rand.Reader, big.NewInt(8))
			word := RandomString(int(num.Int64()) + 3)
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
		num, _ := rand.Int(rand.Reader, big.NewInt(6))
		wordCount := int(num.Int64()) + 1
		for i := 0; i < wordCount; i++ {
			num, _ := rand.Int(rand.Reader, big.NewInt(8))
			word := RandomString(int(num.Int64()) + 3)
			word = strings.ToLower(word)
			result = append(result, word)
			currentLength += len(word)
		}
	}

	return strings.Join(result, "-")
}

func RandomPassword(length int) string {
	if length < 6 {
		length = 10
	}

	const (
		lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
		uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numericChars   = "0123456789"
		specialChars   = "!@#$%^&*(),.?\":{}|<>"
	)

	randomUppercase, err := getRandomCharacter(uppercaseChars)
	if err != nil {
		return "abc123..ABC"
	}

	randomNumber, err := getRandomCharacter(numericChars)
	if err != nil {
		return "abc123..DEF"
	}

	randomSpecialChar, err := getRandomCharacter(specialChars)
	if err != nil {
		return "abc123..GHI"

	}

	remainingLength := length - 3

	remainingChars := lowercaseChars + uppercaseChars + numericChars
	var passwordBuilder strings.Builder
	for i := 0; i < remainingLength; i++ {
		char, err := getRandomCharacter(remainingChars)
		if err != nil {
			return "abc123..JKL"
		}
		passwordBuilder.WriteString(char)
	}

	password := randomUppercase + randomNumber + randomSpecialChar + passwordBuilder.String()

	return shuffleString(password)
}

func getRandomCharacter(chars string) (string, error) {
	index, err := cryptoRandInt(len(chars))
	if err != nil {
		return "", err
	}
	return string(chars[index]), nil
}

func cryptoRandInt(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("max must be greater than 0")
	}

	bytes := make([]byte, 1)
	_, err := rand.Read(bytes)
	if err != nil {
		return 0, err
	}

	return int(bytes[0]) % max, nil
}

func shuffleString(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		j, _ := cryptoRandInt(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
