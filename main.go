package main



import (
	"time"
	"fmt"
	"bufio"
	"os"
)

func (s *slide) Borrar(id int) {
	var aux []*procesos
	for _, v := range s.Procesos {
		if v.Id == id {
			fmt.Println("ID : ", v.Id, " Borrado")
			v.Inicia = false

		} else {
			aux = append(aux, v)
		}
	}
	s.Procesos = aux
}


func main(){
	
	
	var salir = true
	var opc string
	var oproc int
	inicia := make(chan bool)
	var ID []*procesos	
	var addpro *procesos
	var cadena = &slide{Procesos:ID,}
	var id int
	var exit = bufio.NewScanner(os.Stdin)
	
	for sal := true; sal; sal = salir {
		fmt.Println("\n------ ADMINISTRADOR DE PROCESOS ------")
		fmt.Println("1 - AGREGAR PROCESO ")
		fmt.Println("2 - MOSTRAR PROCESOS")
		fmt.Println("3 - ELIMINAR PROCESOS")
		fmt.Println("4 - SALIR")
		fmt.Scan(&opc)

		switch opc {
		case "1":
			
			
			
			
			fmt.Println("Agregando Proceso", id)
			
			addpro = &procesos{Id: id}
			cadena.Inserta(addpro)
			go addpro.start()
			id = id + 1			
			break		
			
			
		case "2":
			
			
			go cadena.Muestra(inicia)
			exit.Scan()
			inicia <- true
			break

		case "3":
			
			fmt.Println("Numero ID por BORRAR? ")
			fmt.Scan(&oproc)
			go cadena.Borrar(oproc)
			break

		
		case "4":
				
			salir = false
		
		
		default:
			fmt.Println("OPCION INVALIDA")
		}
	}	
	
}


type procesos struct {
	Id int
	Inicia bool
	contador uint64
}
	
func (pro *procesos) start() {
	pro.Inicia = true
	pro.contador = 0
	for {
		pro.contador++
		time.Sleep(time.Millisecond * 500)
		if !pro.Inicia {
			break
		}
	}
}


type slide struct {
	Procesos []*procesos
}

func (s *slide) Inserta (pro *procesos) {
	s.Procesos = append(s.Procesos,pro)
}

func (s *slide) Muestra(inicia chan bool) {

	for {
		select {
		case <-inicia:
			return
		default:
			for _, v := range s.Procesos {
				fmt.Println("ID: ", v.Id, " -- ", v.contador)
				time.Sleep(time.Millisecond * 500)

			}

		}
	}
}

