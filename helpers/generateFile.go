package helpers

import (
	"fmt"
	"os"
	// "github.com/thanhpk/randstr"
)

func WriteFile(result string) {
	// f, err := os.Open("a.txt")
	// HandleError(err)

	// defer f.Close()

	// _, err2 := f.WriteString(result)
	// HandleError(err2)
}

func GenerateFile() {
	// token := randstr.String(10)
	// filename := token + ".txt"
	f, err := os.Create("a.txt")
	HandleError(err)

	f.Chmod(0666)

	// err3 := os.Chmod("a.txt", 0755)
	// HandleError(err3)

	defer f.Close()

	file, err2 := os.Open("a.txt")
	HandleError(err2)

	fmt.Println(file)

}
