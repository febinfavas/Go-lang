package main  
import (  
   "fmt"  
)  
type user struct {  
   name string  
   email string  
   password string  
}  
func main() {  
   u := user{name: "John", email: "John@gmail.com", password: "password", }  
   fmt.Println(u)  
   fmt.Println("Name is :: ", u.name)
   fmt.Println("Email is :: ", u.email)
   fmt.Println("Password is :: ", u.password)  
}  