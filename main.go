package main

import (
  "github.com/julienschmidt/httprouter"
  
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

func main() {
  router := httprouter.New()
  router.GET("/status", statusHandler)
  router.GET("/account/:accountName", accountHandler)
  router.GET("/", indexHandler)
  
  router.NotFound = http.HandlerFunc(notFound)
  router.MethodNotAllowed = http.HandlerFunc(notAllowed)
  router.PanicHandler = panicHandler
  
  log.Println("server START")
  log.Fatal(http.ListenAndServe(":8085", router))
  
}

// errors and stuffs

func notFound(responseWriter http.ResponseWriter, request *http.Request) {
  printJsonError(responseWriter, request, 404, "Not Found")
}

func notAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	printJsonError(responseWriter, request, 405, "Not allowed")
}

func panicHandler(responseWriter http.ResponseWriter, request *http.Request, _ interface{}) {
	printJsonError(responseWriter, request, 500, "Internal server error")
}

// regular router funcs

func indexHandler(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
  printJsonKeyValueResponse(responseWriter, request, "ETHKEDDIT", "YEET")
}

func statusHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
    // TODO: put a function here to check connection to ipc and bindings
    printJsonKeyValueResponse(responseWriter, request, "Status", "Ok")
}

func accountHandler(responseWriter http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
  // TODO: put function here to get latest post
  printJsonKeyValueResponse(responseWriter, request, "good shit", "Ok")
}


// responses

func printJsonResponse(responseWriter http.ResponseWriter, request *http.Request, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Println("500 " + request.URL.String())
		printJsonError(responseWriter, request, 500, "Json marshal failed")
	} else {
		log.Println("200 " + request.URL.String())
		fmt.Fprintf(responseWriter, "%s", json)
	}
}

func printJsonKeyValueResponse(responseWriter http.ResponseWriter, request *http.Request, key string, value string) {
	log.Println("200 " + request.URL.String())
	fmt.Fprintf(responseWriter, keyValueJson(key, value))
}

func printJsonError(responseWriter http.ResponseWriter, request *http.Request, statusCode int, errorMessage string) {
	log.Println(fmt.Sprintf("%d %s", statusCode, request.URL.String()))
	responseWriter.WriteHeader(statusCode)
	fmt.Fprintf(responseWriter, keyValueJson("Error", errorMessage))
}

func keyValueJson(key string, value string) string {
	return fmt.Sprintf(`{"` + key + `":"` + value + `"}`)
}
