package main

import (
        "strings"
        "fmt"
       )

func main(){
    origin_str := "space1     space2          space3"
    
    //to become "space1 space2 space3"
    new_str := strings.Join(strings.Fields(origin_str)," ")

    fmt.Println(new_str)
}
