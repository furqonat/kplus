package utils

import (
	"regexp"
	"strconv"

	"go.uber.org/fx"
	"golang.org/x/exp/rand"
)

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewJwt),
	fx.Provide(NewRandomIntGenerator),
	fx.Provide(NewSqlDB),
)

func Int64Pointer(value int64) *int64 {
	return &value
}

func IsPhoneNumber(input string) bool {
	phoneRegex := regexp.MustCompile(`^\d{10,15}$`)
	return phoneRegex.MatchString(input)
}

func IsGmailAddress(input string) bool {
	gmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail\.com$`)
	return gmailRegex.MatchString(input)
}

func StringPtr(value string) *string {
	return &value
}

func IntPtr(value int) *int {
	return &value
}

func StringToInt(s string) int {
	if v, err := strconv.Atoi(s); err != nil {
		return 0
	} else {
		return v
	}
}

type RandomIntGenerator interface {
	RandomInt(min, max int) int
}

type DefaultRandomIntGenerator struct{}

func (d DefaultRandomIntGenerator) RandomInt(min, max int) int {
	return randomInt(min, max)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func NewRandomIntGenerator() RandomIntGenerator {
	return DefaultRandomIntGenerator{}
}

func StringToFloat(s string) float64 {
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		return 0
	} else {
		return v
	}
}
