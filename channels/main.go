package main

import ("fmt"
    "net/http"
    "time"
)

func main() {
    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://golang.org",
        "http://amazon.com",
}

//with channels theory

    c := make(chan string)// that how we create a new channel of type string

// witout routines theorys
    //for _, link := range(links){
     //   checkLink(link)

    //}

// with go routines theori
    for _, link := range(links){
        // go reserved word say Golang "create a new thread go routine for this"
        go checkLink(link, c)
    }

//    fmt.Println(<-c)
    // if we use just one the run time of main will kill the prompt before
    // other operation finish **check screenshots**
//    fmt.Println(<-c)
//    fmt.Println(<-c)
//    fmt.Println(<-c)
//    fmt.Println(<-c)
//    fmt.Println(<-c)
//    fmt.Println(<-c)
    // if the number of "waiting" channels exced the number of links it will
    // be wating forever
    // fmt.Prinln just wakeup when we recive a new value via channel

//    for i :=0; i< len(links); i++{
//        fmt.Prinln(<-c)
        //that will make that every time the routine finish it will start again
        // more like a status checket that ping 
    //for{
    //    go checkLink(<-c, c)
    //}
    // we will make a little change for stylin reasons
    for l := range c {
        //now we will make the routine "wait"
        //time.Sleep(5* time.Scond)
        //time.Sleep will sleep the current go routine
        //but if we put this function here will sleep the main routine
        //and thats a not desired behavior
        //go checkLink(l, c)
        // we will fix it with Function literal
        go func(link string) {
            time.Sleep(5* time.Second)
            checkLink(link, c)
        }(l)// in Literl functions we have to put another pair of parenthesis
        //at the end to invoque it
    }
    // using range on a channel means wait for the valur on channel
    // it menas that l := range c is equivalen to <-c

}
// in go Lmabda function is called FuctionLiteral
func checkLink(link string, c chan string) {// we have to modify the function
    // to acept the channel as argument
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link + " might be down!")
        //c <- "Might be down I think"
        //that will make that every time the routine finish it will start again
        // more like a status checket that ping continuous
        c <- link
        return
    }


    fmt.Println(link, "is up!")
//    c <- "Yep its up"
        //that will make that every time the routine finish it will start again
        // more like a status checket that ping continuous
    c <- link

}
// concurrency -> we can have multiple threads execting code. If one thread
// blocks another one is picked up and worked on / not necesary using multple
// cores

// parallelism -> Multople threads executed at the exac same time.
// / requieres multiple CPU's

// channels comunicate between routines and they must be of type they will
// comunicate (strings, ints, etc)

// channel <- 5 / Send the value 5 into this channel
// myNumber <- channel / Wait for a value to be sent into the channel. When
// we get opne assign the value to myNumber
// fmt.Println(<-channel) Wait for a value to be sent into the canner. When
// we get one, log it our inmediately (use it as argument)
