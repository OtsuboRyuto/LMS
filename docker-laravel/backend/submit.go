package main
import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
)

type RequestBody struct {
    id int
    group int
    contents int
}


func main(c web.C, w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)

    var posted RequestBody
    json.Unmarshal(body, &posted)
    id := posted.id
    group := posted.group
    contents := posted.contents
    fmt.Println(id)
    fmt.Println(group)
    fmt.Println(contents)
    return;
}