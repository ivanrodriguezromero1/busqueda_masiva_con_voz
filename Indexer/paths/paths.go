package paths

import (
	"io/fs"
	"io/ioutil"

	"../modelos"
	"../utils"
)

func SetPaths(pathDir string, files []fs.FileInfo, c *modelos.Carpeta) {
	for _, file := range files {
		fileName := file.Name()
		isDir := file.IsDir()
		if !isDir {
			newFile := new(modelos.Archivo)
			inputFile := c.Ruta + "/" + fileName
			newFile.Ruta = inputFile
			c.Archivos = append(c.Archivos, *newFile)
		} else {
			newPathDir := pathDir + "/" + fileName
			newDir, newError_lectura := ioutil.ReadDir(newPathDir)
			utils.PrintError(newError_lectura)
			newFolder := new(modelos.Carpeta)
			newFolder.Ruta = newPathDir
			c.Subcarpetas = append(c.Subcarpetas, *newFolder)
			positionNewFolder := len(c.Subcarpetas) - 1
			SetPaths(newPathDir, newDir, &c.Subcarpetas[positionNewFolder])
		}
	}
}
