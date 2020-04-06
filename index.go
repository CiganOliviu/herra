package main

import (
  "net/http"
  "./urls"
  "./setup"
)

func main () {

  http.HandleFunc ("/", urls.Index)
  http.ListenAndServe (setup.PORT, nil)

}
