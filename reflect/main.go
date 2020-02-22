package main

import (
    "fmt"
    "reflect"
)

type User struct {
  Name string
  Age int
  Gender string
}
func reflectType(i interface{}){
  t := reflect.TypeOf(i)
  fmt.Println(t.Name(),t.Kind())
}
func main() {
    u := User{
      "kilric",
      24,
      "男",
    }
    
    reflectType(u)
    
    fmt.Println(reflect.TypeOf(18).Name())
}
