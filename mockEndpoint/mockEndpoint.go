package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.PUT("/", hello)

	e.Start(":8080")

}

func hello(c echo.Context) error {

	req := c.Request()
	log.Printf("[hello] - proto:%s host:%s remoteAddr:%s method:%s path:%s", req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path)

	payload := make(map[string]interface{})
	err := json.NewDecoder(req.Body).Decode(&payload)
	defer req.Body.Close()

	log.Printf("Body: %v\n", payload)
	log.Printf("Header: %v\n", req)
	log.Printf("Length: %v\n", req.ContentLength)


	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	} else {
		return c.String(http.StatusOK, "Ok\n")
	}

}
