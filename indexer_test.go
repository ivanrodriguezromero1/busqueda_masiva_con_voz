package main

import (
	"io/ioutil"
	"runtime"
	"sync"
	"testing"

	"./api"
	"./items"
	"./modelos"
	"./paths"
	"./utils"
)

func TestGetVal(t *testing.T) {
	var wg sync.WaitGroup
	var data []*string
	pathDir := "enron_mail_20110402"
	files, error_lectura := ioutil.ReadDir(pathDir)
	utils.PrintError(error_lectura)
	raiz := new(modelos.Carpeta)
	raiz.Ruta = pathDir
	paths.SetPaths(pathDir, files, raiz)
	index_name := `{"index" : { "_index" : "` + pathDir + `" } }`
	sem := make(chan struct{}, runtime.NumCPU())
	items.SetItems(raiz, &data, &index_name, sem, &wg)
	wg.Wait()
	DataBase := items.GetItems(data)
	api.ApiBulk(DataBase)
}
