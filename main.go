/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-七月-15
 *
*/

package main

import (
    "flag"
    "log"
    "github.com/parnurzeal/gorequest"
)

var title = flag.String("t","","Set movie title")
var limits = flag.Int("l",0,"Set count for download images")
var path = flag.String("p",".","Set path to save images")
func main(){
    flag.Parse()
    if *title == ""{
        log.Println("Please set a title.")
    }else{
        cli := ConnectionClient{
            req:gorequest.New(),
            path:*path,
        }
        cli.Search(*title,*limits)
    }
}
