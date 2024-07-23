package utils

import (
	"regexp"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewJwt),
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
