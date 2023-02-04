package modelos

type Carpeta struct {
	Ruta        string    `json:"ruta"`
	Archivos    []Archivo `json:"archivos"`
	Subcarpetas []Carpeta `json:"subcarpetas"`
}
