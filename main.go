package main

import httpserver "study/http_server"

func main() {
	if err := httpserver.StartHttpServer(); err != nil {
		panic(err)
	}
}
