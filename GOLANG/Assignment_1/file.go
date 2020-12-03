package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("hello.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hello World !!!. Welcome in Go Programming World !!!")  
   file.Close()  
   stream, err:= ioutil.ReadFile("hello.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  