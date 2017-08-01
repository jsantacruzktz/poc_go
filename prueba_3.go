package main
/**
Go - Multiplexing Routines

Channel(chan) en Go son "first-class" es decir que son nativos como int o string.
 */
import (
    "fmt"
    "time"
    "math/rand"
)


func main() {
    /*fmt.Println(" channelGenerator ------------------------------------- ")

    //Primero vemos como se comporta el patron de channel generator.
    channelGenerator()

    fmt.Println(" notMultiplexed -------------------------------------- ")

    //Ejecutamos tarea concurrente sin "multiplexar"
    notMultiplexed()

    fmt.Println(" multiplexed -------------------------------------- ")

    multiplexed()

    fmt.Println(" selectMultiplexed -------------------------------------- ")

    selectMultiplexed()

    fmt.Println(" ThePowerOfRoutines -------------------------------------- ")
    */

    blowThoseRoutinesChain()

}

/**
Un channel generator es un "pattern" de en la programacion concurrente en go.
Permite generar metodos contenidos en chnnels que facilitan el manejo de la comunicacion
de datos refresco de estados.
 */
func channelGenerator(){
    c := generateChannel("Channel Generated.")
    for i := 0; i<5; i++ {
        fmt.Printf("Cahnnel messaged: %q \n", <-c)
    }
}

/**
Channel generator implementation.
 */
func generateChannel(msg string) <- chan string{
    c := make(chan string)
    go func(){
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            sleepRandomlyTime( 1e3 )
        }
    }()

    return c
}

/**
Delays the current thread the time expressed in miliseconds.
 */
func sleepRandomlyTime( napTime int ){
    var randomSleepTime = time.Duration(rand.Intn( napTime )) * time.Millisecond
    time.Sleep( randomSleepTime )
}

/**
Aca vemos como podemos ejecutar distintas actividades multi-tareas manteniendo un channel vivo

NOTA: Si bien esto se esta ejecutando en multiple Rutinas podemos ver como la ejecucion es pausada.
Esta pausa la genera el for cundo quiere hacer el print por que al querer hacer el print tenemos
que esperar que cada posteo del channel se ejecute "secuencialmente" para poder imprimirlo.
 */
func notMultiplexed() {
    joel := generateChannel("Joel messages:")
    ana := generateChannel("Ana messages:")

    for i:= 0; i<5; i++ {
        fmt.Println( <-joel )
        fmt.Println( <-ana )
    }
}

/**
The multiplexer works like a single channel listeng a bunch of other channels and every channel post
their news to the multiplexer channel wichs  dispatchs each new "news" from the other channels.


0----| \
     |  \
0----|   > ---> O -> post every new from each channel.
     |  /
0----| /
 */
func multiplexed() {
    c := multiplexer(
        generateChannel("Joel multi message:"),
        generateChannel("Ana multi message:") )

    for i:= 0; i<10; i++ {
        fmt.Println( <-c )
    }

    fmt.Println(" Lo mismo pero mas mejor -------------------- ")

    secondC := massiveMultiplexer(
        generateChannel("Manuel multi message:"),
        generateChannel("Fede multi message:"),
        generateChannel("Maria multi message:"),
        generateChannel("Camila multi message:"))

    for i:= 0; i<10; i++ {
        fmt.Println( <-secondC )
    }
}

/**
* Better than multiplexed()
 */
func selectMultiplexed() {
    c := selectMultiplexer(
        generateChannel("Joel multi message:"),
        generateChannel("Ana multi message:") )

    for i:= 0; i<10; i++ {
        fmt.Println( <-c )
    }
}

/**
* Straight fowrd multiplexed
 */
func multiplexer(input1, input2 <- chan string) <-chan string{
    c := make(chan string)

    go func(){ for { c <- <-input1 } }()
    go func(){ for { c <- <-input2 } }()

    return c
}

func massiveMultiplexer(inputs ... <- chan string) <-chan string{
    c := make(chan string)

    for _, input := range inputs {
        queued := input
        go func(){ for { c <- <-queued } }()
    }

    return c
}


/**
Smart multiplexer
 */
func selectMultiplexer(input1, input2 <- chan string) <-chan string{
    c := make(chan string)

    go func(){
        for {
            select {
                case s := <-input1: c <- s
                case s := <-input2: c <- s
            }
        }
    }()

    return c
}

func blowThoseRoutinesChain(){
    const n = 1e5;

    chainBegin := make(chan int)

    chainNext := chainBegin
    chainOwner := chainBegin

    start := time.Now()

    for i := 0; i < n; i++ {
        chainNext = make(chan int)
        go sendTheMessageChain(chainOwner, chainNext)
        chainOwner = chainNext
    }

    go func(c chan int){ c <- 1 }(chainNext)
    fmt.Println( <-chainBegin )
    elapsed := time.Since(start)

    fmt.Println(elapsed)
}

func sendTheMessageChain(next, owner chan int){
    next <- 1 + <-owner
}

