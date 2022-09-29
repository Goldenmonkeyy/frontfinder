package main

import (
    "net/http"
    "fmt"
    "io"
    "regexp"
    "strings"
    "os"
    "crypto/tls"
)

var function = os.Args[1]
var url = os.Args[2]
var keywords = []string{"password", "api", "email", "api-key", "secret", "key", "密码", "上传", "下载", "秘钥"} 

//function to remove duplicates

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

// function to GET http request

func getRequest(url string) string {

    //ignore certificate
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    resp, err := http.Get(url)
    if err != nil {
        panic (err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic (err)
    }

    responseBody := string(body)
    return responseBody
}

func secretfinder() {

    fmt.Println("\r\nModule 1: secretfinder\r\n\v")

    filter := func(keyword, response string) []string {
        reg := regexp.MustCompile(`(?mi).{20}` + keyword + `.{20}`)
        match := reg.FindAllString(response, -1)
        return match
    }
    
    responseBody := getRequest(url)
    for _, keyword := range keywords {
//    fmt.Println(keyword)
        for _, results := range filter(keyword, responseBody) {
            fmt.Printf("%s\r\n", results)
        }
    }
}

func pathfinder() {

    fmt.Println("\r\nModule 2: pathfinder\r\n\v")
    
    responseBody := getRequest(url)
    //define regex
    reg1 := regexp.MustCompile(`('\/.[^" ><+!@,"|=[\](){};,]+?')`)
    reg2 := regexp.MustCompile(`("\/.[^" ><@!,|"+=[\](){};,]+?")`)
    
    //regex matching
    match1 := reg1.FindAllString(responseBody, -1)
    match2 := reg2.FindAllString(responseBody, -1)
    
    // print the slice line by line      
    for _, path1 := range removeDuplicateStr(match1) {
        fmt.Printf("%s\r\n", strings.Trim(path1, "'"))
    }
    for _, path2 := range removeDuplicateStr(match2) {
        fmt.Printf("%s\r\n", strings.Trim(path2, "\""))
    }
}

func assetfinder() {
   
    fmt.Println("\r\nModule 3: assetfinder\r\n\v")
    responseBody := getRequest(url)
    //filter := func(responseBody string) []string {

    
        reg1 := regexp.MustCompile(`(?mi)`)
        //reg2 := regexp.MustCompile(`(?mi)`)
        //reg3 := regexp.MustCompile(`(?mi)`)
        //for _ := range []string("reg1","reg2","reg3") {}
        match := reg1.FindAllString(responseBody, -1)
    
    for _, asset := range removeDuplicateStr(match) {
    fmt.Printf("\r\n%s\r\n", asset)
    }
}

func main() {
    if function == "secret" {
        fmt.Println("\r\nhere goes the path finder.\r\n")
        secretfinder()
    } else if function == "secret" {
        fmt.Println("\r\nhere goes the secret finder.\r\n")
        pathfinder()
    } else if function == "asset" {
        fmt.Println("\r\nhere goes the asset finder.\n\r")
        assetfinder()
    } else if function == "all" {
        fmt.Println("\r\nhere goes all the functions.\r\n")
        secretfinder()
        pathfinder()
        assetfinder()
    } else {
        fmt.Println("Nope, input 'secret, path, asset, all' are expected.\r\n")
    }
}

