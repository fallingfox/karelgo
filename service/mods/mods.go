package mods

import "strings"
import "math/rand"
import "errors"

/**
	~ Missing Character ~
	> Funkce vymění náhodný znak v řetězci za token

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

/**
	~ Switch Characters ~
	> Funkce prohodí pozice náhodných dvou znaků v řetězci

	@param word řetězec, který bude modifikován
	@return string modifikovaný řetězec
*/
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

/**
	~ Characters to Numbers ~
	> Funkce vymění znaky v řetězci za čísla podle kódovací tabulky

	@param word řetězec, který bude modifikován
	@return string modifikovaný řetězec
*/
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

/**
	~ Reverse Word ~
	> Fuknce zrcadlově prohodí pořadí všech zanků v řetězci

	@param word řetězec, který bude modifikován
	@return string modifikovaný řatězec
*/
func Reverse(word string) (string, error) {
	// kontrola vstupního řetězce
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// alokace pole pro opačný řetězec
	//arr := make([]string, len(word))
	arr := strings.Split(word, "")

	// rozdělení původního řetězce na znaky
	//wrdArr := strings.Split(word, "")

	// převrácení řetězce
	/*
	for i := 0; i < len(word); i++ {
		arr[len(arr) - i - 1] = wrdArr[i]
	}
	*/
	for i, j := 0, len(word) - 1; i < j; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	// spojení znků do řetězce
	return strings.Join(arr, ""), nil
}

/**
	~ Random Mix ~
	> Funkce náhodně prohází všechny znaky v řetězci

	@param word řetězec, který bude modifikován; nesmí být prázdný
	@return string modifokovaný řetězec
*/

func Random(word string) (string, error) {
	// kontrola vstupního řetězce
	if ok, err := checkWord(word); !ok {
		return "", err
	}

	// rozdělení vstupního řetězce na jednotlivé znaky
	arr := strings.Split(word, "")
	// alokace pole pro výstupní řetězec
	arrOut := make([]string, len(word))

	// náhodné prohození všech zanků v řatězci
	for i := 0; i < len(word); i++ {
		p := rand.Intn(len(arr)) // náhodná pozice
		arrOut[i] = arr[p] // vložení znaku na náhodné pozicic do výstupního řetězce
		arr[p] = arr[len(arr) - 1] // přesunutí znaku z konce pole na pozici
		arr = arr[:len(arr) - 1] // odříznutí posledního znaku pole
	}

	// spojení pole do řetězce
	return strings.Join(arrOut, ""), nil
}

/**
	> Pomocná funkce zkontroluje zdali vstupní řetězec není prázdný

	@param word kontrolovaný řetězec
	@return bool false pokud je řetězec prázný, jinak true
	@return error pokud je řetězec prázdný, jinak nil
*/
func checkWord(word string) (bool, error) {
	if (len(word) == 0) {
		return false, errors.New("Vstupní slovo je prázdný řetězec!")
	}
	return true, nil
}

/**
	> Pomocná funkce zkontroluje zdali token není prázný řetězec nebo že neobsahuje víc jak 1 znak

	@param token token
	@return bool false pokud je token prázný řetězec nebo obsahuje víc jak 1 znak, jinak true
	@return error pokud je token chybný, jinak nil
*/
func checkToken(token string) (bool, error) {
	if (len(token) == 0) {
		return false, errors.New("Token je prázdný řetězec!")
	}

	if (len(token) > 1) {
		return false, errors.New("Token je větší než jeden znak!")
	}

	return true, nil
}