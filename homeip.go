package main

import (
    "os"
    "net/http"
    "fmt"
    "io/ioutil"
    log "github.com/Sirupsen/logrus"
    "net"
)

func main() {

    log.Info("Starting homeip service")

    secret := os.Getenv("SECRET")
    if secret == "" {
        log.Fatal("No secret set. Set one with env var SECRET.")
    }

    port := os.Getenv("PORT")
    if port == "" {
        log.Warn("No port set. Using default of 8080. Set one with env var PORT.")
        port = "8080"
    }

    current_ip := "none"

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

        if r.Method == "POST" {
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Error(err)
                w.WriteHeader(500)
                return
            }
            if string(body) == secret {
                log.Warn("Correct secret received, updating")
                current_ip, _, err = net.SplitHostPort(r.RemoteAddr)
                if err != nil {
                    log.Error(err)
                    w.WriteHeader(500)
                    return
                }
                fmt.Fprint(w, current_ip)
                return
            } else {
                log.Warn("Wrong secret received")
                w.WriteHeader(403)
                return
            }
        } else {
            fmt.Fprint(w, current_ip)
        }

    })
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Error(err)
    }

    log.Info("Exiting homeip service")
}

