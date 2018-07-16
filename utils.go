/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-七月-15
 *
*/

package main

import (
    "github.com/parnurzeal/gorequest"
    "log"
    "os"
)

const(
    dbQueryURL = `https://api.douban.com/v2/movie/search`
    dbMovieURL = `https://api.douban.com/v2/movie/subject/`
)

type Movies []MovieInfo

type MovieInfo struct {
    Title               string
    OriginalTitle       string `json:"original_title"`
    Year                string
    Images              ImageInfo
    Id                  string
    Rating              Rates
}

type ImageInfo struct {
    Small       string
    Large       string
    Medium      string
}

type SearchInfo struct {
    Count       int
    Start       int
    Total       int
    Subjects    Movies
}

type Rates struct{
    Max         int
    Average     float32
    Stars       string
    Min         int
}

type ConnectionClient struct {
    req         *gorequest.SuperAgent
    path        string
}

func (cli ConnectionClient) Init(){
    cli.req = gorequest.New()
}

func (cli ConnectionClient) Search(title string, limits int){
    cond := "q="+title
    var res = SearchInfo{}
    if cli.req == nil{
        cli.req = gorequest.New()
    }

    resp, body, err := cli.req.Get(dbQueryURL+"?"+cond).EndStruct(&res)
    log.Println(resp.Status)
    log.Println("body:"+string(body))
    if err != nil{
        log.Println(err)
    }else{
        //log.Println(res.Subjects)
        cli.searchMovies(res.Subjects,limits)
    }
}

func (cli ConnectionClient) searchMovies(movies Movies,limits int) {
    log.Println("Searching imgs...")
    if limits == 0 {
        for _, m := range movies {
            _, imageBytes, errs := cli.req.Get(m.Images.Large).EndBytes()
            if errs != nil {
                log.Println(errs)
            } else {
                log.Println("Write images..."+m.Title)
                image, e := os.Create(m.Title + "-" + m.Year + ".jpg")
                if e != nil {
                    log.Println(e)
                } else {
                    image.Write(imageBytes)
                }
            }
        }
        return
    }
    for index, m := range movies {
        log.Println(m)
        if index+1 <= limits {
            _, imageBytes, errs := cli.req.Get(m.Images.Large).EndBytes()
            if errs != nil {
                log.Println(errs)
            } else {
                image, e := os.Create(m.Title + "-" + m.Year + ".jpg")
                if e != nil {
                    log.Println(e)
                } else {
                    image.Write(imageBytes)
                }
            }
        } else {
            return
        }
    }
}