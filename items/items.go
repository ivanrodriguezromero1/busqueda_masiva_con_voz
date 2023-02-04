package items

import (
	"strings"
	"sync"

	"../lectura"
	"../modelos"
)

func SetItem(fileName string, index_name *string, sem chan struct{}, wg *sync.WaitGroup) string {
	value, read := lectura.LecturaConcurrente(fileName, sem, wg)
	if !read {
		return ""
	}
	var item []string
	item = append(item, *index_name)
	item = append(item, value)
	return strings.Join(item, "\n")
}
func SetItems(c *modelos.Carpeta, datos *[]*string, index_name *string, sem chan struct{}, wg *sync.WaitGroup) {
	for _, archivo := range c.Archivos {
		wg.Add(1)
		fileName := archivo.Ruta
		var item string
		go func() {
			item = SetItem(fileName, index_name, sem, wg)
		}()
		*datos = append(*datos, &item)
	}
	for _, carpeta := range c.Subcarpetas {
		SetItems(&carpeta, datos, index_name, sem, wg)
	}
}
func GetItems(datos []*string) string {
	var value_items []string
	for i := 0; i < len(datos); i++ {
		value_items = append(value_items, *datos[i])
	}
	return strings.Join(value_items, "\n")
}
