package main

import (
  "net/http"
  "./urls"
)

func main () {

  http.HandleFunc ("/", urls.Index)
  http.ListenAndServe (":8000", nil)

}
