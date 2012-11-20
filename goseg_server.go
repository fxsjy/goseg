package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "./goseg"
    "time"
)

func cutAction(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    sentence := r.Form["content"]
    if sentence==nil || len(sentence)==0{
    	fmt.Fprintf(w, "Failed.") 
    }else{
        t1 := time.Now()
    	lines := strings.Split(sentence[0],"\n")
    	result := make([]string,0)
    	for _,line := range lines{
    		words := goseg.Cut([]rune(line))
    		for _,w := range words{
    			result = append(result,w)
    		}
    		result = append(result,"\n")
    	}
    	fmt.Fprintf(w,strings.Join(result,"/ "))
        log.Printf("cut %v bytes, time costs: %v", len(sentence[0]), time.Now().Sub(t1))
    }
}

func indexPage(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form1.html")
        t.Execute(w, nil)
    } else {
        fmt.Fprintf(w, "No such method")
    }
}

func main() {
    http.HandleFunc("/cut", cutAction)       //设置访问的路由
    http.HandleFunc("/", indexPage)         //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    log.Println("listening at :9090 port")
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
