package main

import ("fmt"
    "net/http"
    "os"
    "io"
)

//https://www.airbnb.es/rooms/19676874?c=.pi80.pkbWVzc2FnaW5nL25ld19tZXNzYWdl&euid=e69390ac-2b0e-2cf8-aba3-8007fb1a2f47&source_impression_id=p3_1625843814_oVSvzWyMkCQx%2F9a9&guests=1&adults=1https://www.airbnb.es/rooms/19676874?c=.pi80.pkbWVzc2FnaW5nL25ld19tZXNzYWdl&euid=e69390ac-2b0e-2cf8-aba3-8007fb1a2f47&source_impression_id=p3_1625843814_oVSvzWyMkCQx%2F9a9&guests=1&adults=1func main() {
func main(){
    resp, err := http.Get("http://google.com")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
// if we asign a inferface as type on a struct it means that
//every value that success the condition of interfaces can be there
//we can embed one interface inside other ieg
// type x interface{
//  interface 1
//  interface 2
//}
// that means that things must succes the condition of interface 1 and interface 2

    //bs := make([]byte, 999999)
    //resp.Body.Read(bs)//on go we pass the byte slice and read will inject the values on it
    //fmt.Println(string(bs))
    lw := randomWriter{}
//normal style
    //io.Copy(io.Stdout, resp.Body)
//custom style
    io.Copy(lw, resp.Body)
    // the default structure of Copy function is 
    // func (Copy dst Writer, src Reader) (written int64, err error) 
}
// We will create a custom weriter

type randomWriter struct{}

func (randomWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Just wrote this many bytes:", len(bs))

    return len(bs), nil
}

