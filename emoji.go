package nigari

func IsEmoji(c rune) bool {
	return emojilist[c]
}
