package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		// fmt.Println(servidor, "No esta disponible")
		canal <- servidor + "No esta disponible"
	} else {
		// fmt.Println(servidor, "Esta funcionando")
		canal <- servidor + "Esta funcionando"
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string)
	urls := []string{
		"http://www.udemy.com/",
		"http://www.google.com/",
		"http://www.facebook.com/",
		"http://www.youtube.com/",
	}

	for _, url := range urls {
		// revisarServidor(url)
		// go manda en concurrencia cada request
		//tenemos que comunicarnos por canales
		go revisarServidor(url, canal)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-canal)
	}

	tiempoEjecucion := time.Since(inicio)
	fmt.Println("Tiempo de ejecución: ", tiempoEjecucion)
}
