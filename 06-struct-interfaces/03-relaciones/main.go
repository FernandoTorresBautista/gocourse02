package main

import "fmt"

// relacion 1 a 1
type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// relacion uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	// aux := User{
	// 	nombre: "Fer",
	// 	email:  "fer@test.com",
	// 	activo: true,
	// }
	// aux2 := User{
	// 	nombre: "Fer2",
	// 	email:  "fer2@test.com",
	// 	activo: true,
	// }
	// auxStudent := Student{
	// 	user:   aux,
	// 	codigo: "001",
	// }
	// fmt.Println(aux2)
	// fmt.Println(auxStudent.user.nombre)
	// fmt.Println(auxStudent)

	v1 := Video{titulo: "01-introduccion"}
	v2 := Video{titulo: "02-Instalacion"}
	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{v1, v2},
	}
	v1.curso = curso
	v2.curso = curso

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
