# Búsqueda Masiva por texto y voz
El portal web desarrollado tiene la capacidad de realizar una búsqueda en conjunto masivo de datos, de una manera muy rápida utilizando simplemente la voz o una entrada de texto normal, todo esto gracias a tecnologías como zincsearch , golang, chi, vue , tailwind y SpeechRecognition.

Para la carga de la base de datos, puede utilizar el siguiente comando:

<b> - go run Indexer\indexer.go (nombre de la carpeta que contiene la base de datos)</b>
  
En el ejemplo:
  
<b> - go run Indexer\indexer.go enron_mail_20110402</b>
  
Para levantar la interfaz gráfica, puede utilizar el siguiente comando:

<b> - go run mamuro.go -port (número del puerto)</b>

En el ejemplo:

<b> - go run mamuro.go -port 3000</b>

![mamuroEmail](https://user-images.githubusercontent.com/100105456/216748743-c71dfdfa-c0cf-4b72-a69c-5b82014b6d5b.png)
