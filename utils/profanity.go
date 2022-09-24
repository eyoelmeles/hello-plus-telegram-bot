package utils

import (
	"github.com/finnbear/moderation"
)

func Profanity(text string) string {
	safeMessage,_ := moderation.Censor(text, moderation.Any)
	return safeMessage
}