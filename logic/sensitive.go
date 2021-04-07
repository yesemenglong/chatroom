package logic

import (
	"chatroom/global"
	"log"
	"strings"
)

func FilterSensitive(content string) string {

	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
		log.Println(content)
	}
	return content
}
