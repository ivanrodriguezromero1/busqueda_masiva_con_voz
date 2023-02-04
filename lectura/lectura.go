package lectura

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"../utils"
)

func LecturaConcurrente(inputFile string, sem chan struct{}, wg *sync.WaitGroup) (string, bool) {
	sem <- struct{}{}
	defer func() { <-sem }()
	defer wg.Done()
	return Lectura(inputFile)
}
func Lectura(inputFile string) (string, bool) {
	file, errOpen := os.Open(inputFile)
	utils.PrintError(errOpen)
	defer file.Close()
	maxSz := 10
	b := make([]byte, maxSz)
	var lineas []string
	for {
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		lineas = append(lineas, string(b[:readTotal]))
	}
	contenido := strings.Join(lineas, "")
	if len(contenido) > 1000000 ||
		strings.Index(contenido, "Message-ID:") == -1 {
		return "", false
	}
	item, errJson := json.Marshal(utils.Estructurar(inputFile, contenido))
	utils.PrintError(errJson)
	return string(item), true
}
