package main

import (
                "html/template"
        //  "io/ioutil"
        //        "net/http"
        "fmt"
              "os"
            "bufio"
       )

func main() {

    type News struct {
        Level int
        Title string
        Content string
        Html string
    }

    var sn [1]News

        sn[0].Level = 0
        sn[0].Title = "TITLE 0"
        sn[0].Content = "CONTENT 0"
        sn[0].Html = "<b>HTML 0</b>"

        fmt.Println(sn[0])
        t, err := template.ParseFiles("index.html")
        if (err != nil){
            fmt.Println(err)
        }
        // Create writer to file
        f, _ := os.Create("made.html")
        defer f.Close()

        w := bufio.NewWriter(f)
        t.Execute(w,sn)
        w.Flush()
}
