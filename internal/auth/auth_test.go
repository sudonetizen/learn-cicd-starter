package auth

import (
   "fmt"
   "testing"
   "net/http"
)

func TestGetAPIKey(t *testing.T) {

    req1, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {fmt.Println("error with creating request 1")}
    req1.Header.Set("Authorization", "ApiKey  895623")

    req2, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {fmt.Println("error with creating request 2")}

    req3, err := http.NewRequest("GET", "http://example.com", nil)
    if err != nil {fmt.Println("error with creating request 3")}
    req3.Header.Set("Authorization", "ApiKey895623")


    tests := []struct{
        req *http.Request
        err error 
    }{
        {
            req: req1,
            err: nil,
        },
        {
            req: req2,
            err: ErrNoAuthHeaderIncluded,
        },
        {
            req: req3,
            err: ErrMalformedAuthzHeader,
        },
    }

    for _, tst := range tests {
        _, err := GetAPIKey(tst.req.Header)
        if err != tst.err {t.Errorf("failed: %v", err)}
    }    

}
