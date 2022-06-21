package main

import (
    "os"
    "net/http"

    zlog "github.com/rs/zerolog/log"
)

func main() {

    resp, err := http.Get("http://localhost:3000/")

    if err != nil {
      zlog.Print("Got error during healthcheck", err)
      os.Exit(-1)
    }

    if resp.StatusCode == 200 {
      zlog.Print("Health Check response:", resp.StatusCode)
      os.Exit(0)
    }

    zlog.Print("Unable to Fetch healthcheck ", resp, err)
    os.Exit(1)
}
