package softmem

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func getImagePathByString(number string) string {
	length := len(strings.Trim(number, " "))
	if length == 1 {
		return getSingleImagePath(number)
	} else {
		return getDoubleImagePath(number)
	}
}

func getSingleImagePath(number string) string {
	return fmt.Sprintf("assets/images/single/%s.png", strings.Trim(number, " "))
}

func getDoubleImagePath(number string) string {
	return fmt.Sprintf("assets/images/double/%s.png", strings.Trim(number, " "))
}

func getMnemonicsForNumber(number int) string {
	switch number {
	case 0:
		return "S"
	case 1:
		return "T"
	case 2:
		return "N"
	case 3:
		return "M"
	case 4:
		return "R"
	case 5:
		return "L"
	case 6:
		return "Sh"
	case 7:
		return "K"
	case 8:
		return "F"
	case 9:
		return "P"
	default:
		return ""
	}
}

func getLongHint() string {
	var result string

	for i := 0; i < 10; i++ {
		if result != "" {
			result += " : "
		}
		result += strconv.Itoa(i) + ":" + getMnemonicsForNumber(i)
	}

	return "(" + result + ")"
}

func getHint(number string) string {
	var result string
	num := strings.Trim(number, " ")

	for i := 0; i < len(num); i++ {
		if result != "" {
			result += " : "
		}
		num, _ := strconv.Atoi(num[i : i+1])
		result += getMnemonicsForNumber(num)
	}

	return "(" + result + ")"
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getResourcePath(fileName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := path.Dir(exePath)

	gladePath := path.Join(exeDir, fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	gladePath = path.Join(exeDir, "assets", fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	gladePath = path.Join(exeDir, "../assets", fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	return gladePath, nil
}
