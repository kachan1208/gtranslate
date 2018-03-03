package gtranslate

import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "regexp"
)

const (
    //Fix for RU-UA/UA-RU translation(wrong encoding)
    BaseUrl   = "https://translate.googleapis.com/translate_a/single"
    UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:57.0) Gecko/20100101 Firefox/57.0" 
)

var (
    client = &http.Client{}
    resultExp = regexp.MustCompile(`\[\[\["([^,]+)",`)
)

func Translate(text string, sourceLang string, targetLang string) (string, error) {
    api, _ := url.Parse(BaseUrl)

    params := url.Values{}
    params.Set("client", "gtx")
    params.Set("dt", "t")
    params.Set("sl", sourceLang)
    params.Set("tl", targetLang)
    params.Set("q", text)

    api.RawQuery = params.Encode()
    fmt.Println(api.String())    

    req, err := http.NewRequest("GET", api.String(), nil)
    if err != nil {
        return "", err
    }

    req.Header.Add("User-Agent", UserAgent)    
    response, err := client.Do(req)
    if err != nil || response.StatusCode != http.StatusOK {
        return "", err
    }
    
    result, err := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    if err != nil {
        return "", err
    }

    word := resultExp.FindStringSubmatch(string(result[:]))[1]

    return word, nil
}
