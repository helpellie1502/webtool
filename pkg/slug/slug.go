package slug

import "strings"

func Slugify(str string) string {
	strLower := strings.ToLower(str)
	words := strings.Fields(strLower)
	strClean := strings.Join(words, "-")

	translit := strings.NewReplacer(
		"а", "a", "б", "b", "в", "v", "г", "g", "д", "d",
		"е", "e", "ё", "e", "ж", "zh", "з", "z", "и", "i",
		"й", "y", "к", "k", "л", "l", "м", "m", "н", "n",
		"о", "o", "п", "p", "р", "r", "с", "s", "т", "t",
		"у", "u", "ф", "f", "х", "kh", "ц", "ts", "ч", "ch",
		"ш", "sh", "щ", "shch", "ъ", "", "ы", "y", "ь", "",
		"э", "e", "ю", "yu", "я", "ya",
	)

	result := translit.Replace(strClean)
	result = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, result)

	return result
}
