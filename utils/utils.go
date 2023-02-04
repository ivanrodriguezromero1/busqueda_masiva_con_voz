package utils

import (
	"fmt"
	"log"
	"strings"

	"../modelos"
)

func PrintError(err error) {
	if err != nil {
		fmt.Println("fatal")
		log.Fatal(err)
	}
}

func Estructurar(inputFile string, contenido string) modelos.Nodo {
	var nodo modelos.Nodo
	nodo.Ruta = inputFile
	nodo.Message_ID = GetParam(inputFile, &contenido, "Message-ID:")
	nodo.Date = GetParam(inputFile, &contenido, "Date:")
	contenido = strings.Replace(contenido, "X-From:", "X-FROM:", 1)
	contenido = strings.Replace(contenido, "X-To:", "X-TO:", 1)
	nodo.From = GetParam(inputFile, &contenido, "From:")
	nodo.To = GetParam(inputFile, &contenido, "To:")
	nodo.Subject = GetParam(inputFile, &contenido, "Subject:")
	nodo.Mime_Version = GetParam(inputFile, &contenido, "Mime-Version:")
	nodo.Content_Type = GetParam(inputFile, &contenido, "Content-Type:")
	nodo.Content_Transfer_Encoding = GetParam(inputFile, &contenido, "Content-Transfer-Encoding:")
	nodo.X_From = GetParam(inputFile, &contenido, "X-FROM:")
	nodo.X_To = GetParam(inputFile, &contenido, "X-TO:")
	nodo.X_cc = GetParam(inputFile, &contenido, "X-cc:")
	nodo.X_bcc = GetParam(inputFile, &contenido, "X-bcc:")
	nodo.X_Folder = GetParam(inputFile, &contenido, "X-Folder:")
	nodo.X_Origin = GetParam(inputFile, &contenido, "X-Origin:")
	nodo.X_FileName = GetParam(inputFile, &contenido, "X-FileName:")
	nodo.Content = contenido
	return nodo
}
func GetParam(input string, texto *string, find string) string {
	f := strings.Index(*texto, find)
	if f == -1 {
		return ""
	}
	i := f + len(find)
	*texto = (*texto)[i:]
	j := strings.Index(*texto, "\n")
	if j == -1 {
		return *texto
	}
	result := (*texto)[:j]
	*texto = (*texto)[j+1:]
	return result
}
