package tripofgo

import "regexp"

func SimpleRegex(str string) string {
	// all characters is not include alphabet, number and space
	re := regexp.MustCompile(`[^a-zA-Z0-9" "]+`)

	// match regex wtih string and change with ""
	str = re.ReplaceAllString(str, "")

	// return new string
	return str

	// input : I'm Golang Developer!
	// output: Im Golang Developer
}
