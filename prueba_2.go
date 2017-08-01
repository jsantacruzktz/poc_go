package main
/**
Go - Routines

GO Concurrency doesnt comunicate by Sharing memory.
Share memory by comunicating.

 */
import (
    "fmt"
    "time"
    "math/rand"
)


func main() {
    //Channel necesario para la comunicacion entre las distintas Go Routines.
    c := make(chan string)

    //Ejecutamos el metodo "messageSender" como una Go Routine. Es similar a poner el & en bash
    //Para luego generar comunicacion entre las distintas rutinas es necesario pasar un Channel.
    go messageSender("Mensaje! ", c)

    for i := 0; i < 5; i++ {
        //Aca printeamos los mensajes que vayamos allocando en el channel
        //Cuando haces <-c, Go espera que al Channel le pasen un valor para ejecutar.
        //Hasta que no se postea algo no sigue el procesamiento.
        fmt.Printf("Channel comunication: %q\n", <-c)
    }

    fmt.Println("hello, world")

}


func messageSender(msg string, c chan string){
    for i := 0; i<10; i++{
        var randomSleepTime = time.Duration(rand.Intn(1e3)) * time.Millisecond
        c <- fmt.Sprintf("Send this message: %s - %d (Then wait %v for next message.)", msg, i, randomSleepTime )
        time.Sleep( randomSleepTime )
    }
}
