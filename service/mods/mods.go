package mods

import "strings"
import "math/rand"
import "errors"

/**
	Metoda vymění náhodný znak v řetězci za token
	@param word řetězec, který bude modifikován
	@param token token, kterým bude nahrazen náhodný znak v řetězci (maximální velikost: 1 znak)
	@return string modifikovaný řetězec
*/
func MissingChar(word string, token string) (string, error) {
	// kontrola délky vstupního slova
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// kontrola délky vstupního tokenu
	if ok, err := checkToken(token); !ok {
		return "", err
	}

	// rozdělení řetězce najednotlivé znaky
	arr := strings.Split(word, "")
	// nahrazení jednoho znaku tokenem
	arr[rand.Intn(len(arr))] = token

	// spojení znaků do řetězce
	return strings.Join(arr, ""), nil
}

func SwitchChars(word string) (string, error) {
	// kontrola délky vstupního řetězce
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// vygenerování dvou náhodných pozic v řetězci
	var posA, posB int
	posA = rand.Intn(len(word))
	posB = posA
	for posA == posB {
		posB = rand.Intn(len(word))
	}

	// rozdělení řetězce na jednotlivé znaky
	arr := strings.Split(word, "")
	// prohození znaků na náhodných pozicích
	temp := arr[posA]
	arr[posA] = arr[posB]
	arr[posB] = temp

	// spojení znaků do řetězce
	return strings.Join(arr, ""), nil
}

func CharsToNumbers(word string) (string, error) {
	// kontrola délky vstupního řetězce
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// pravidla pro výměnu znaků
	rules := "i1z2e3a4s5g6t7b8o0"
	rules = rules + strings.ToUpper(rules)
	arr := strings.Split(rules, "")

	// procházení jednotlivých pravidel a následná výměna znaků
	for i := 0; i < len(rules); i += 2 {
		word = strings.ReplaceAll(word, arr[i], arr[i + 1])
	}

	return word, nil
}

func Reverse(word string) (string, error) {
	// kontrola vstupního řetězce
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// alokace pole pro opačný řetězec
	arr := make([]string, len(word))

	// rozdělení původního řetězce na znaky
	wrdArr := strings.Split(word, "")

	// převrácení řetězce
	for i := 0; i < len(word); i++ {
		arr[len(arr) - i - 1] = wrdArr[i]
	}

	// spojení znků do řetězce
	return strings.Join(arr, ""), nil
}


func checkWord(word string) (bool, error) {
	if (len(word) == 0) {
		return false, errors.New("Vstupní slovo je prázdný řetězec!")
	}
	return true, nil
}

func checkToken(token string) (bool, error) {
	if (len(token) == 0) {
		return false, errors.New("Token je prázdný řetězec!")
	}

	if (len(token) > 1) {
		return false, errors.New("Token je větší než jeden znak!")
	}

	return true, nil
}