package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/labstack/echo"

	"httpstatus/statuscodes"
)

func handleStatusCodeRequest(c echo.Context) error {
	allowedCodes := statuscodes.GetAllowedCodes()
	statusCode, _ := strconv.Atoi(c.Param("statusCode"))

	if allowedCodes.IsInList(statusCode) {
		return c.String(statusCode, "")
	}

	return c.String(500, fmt.Sprintf("Error: status code %s is not available.", c.Param("statusCode")))
}

func main() {
	tlsPtr := flag.Bool("tls", false, "use https (ssl/tls) encryption")
	addrPtr := flag.String("address", "localhost:8080", "server address as hostname:port e.g. 1337.my.org:443")
	certFilePtr := flag.String("certFile", "mycert1.cer", "certificate file (try https://github.com/deckarep/EasyCert)")
	keyFilePtr := flag.String("keyFile", "mycert1.key", "key file (try https://github.com/deckarep/EasyCert)")
	flag.Parse()

	e := echo.New()
	e.File("/", "public/index.html")

	e.GET("/:statusCode", handleStatusCodeRequest)
	e.GET("/:statusCode/*", handleStatusCodeRequest)
	e.POST("/:statusCode", handleStatusCodeRequest)
	e.POST("/:statusCode/*", handleStatusCodeRequest)
	e.PUT("/:statusCode", handleStatusCodeRequest)
	e.PUT("/:statusCode/*", handleStatusCodeRequest)
	e.PATCH("/:statusCode", handleStatusCodeRequest)
	e.PATCH("/:statusCode/*", handleStatusCodeRequest)
	e.DELETE("/:statusCode", handleStatusCodeRequest)
	e.DELETE("/:statusCode/*", handleStatusCodeRequest)
	e.OPTIONS("/:statusCode", handleStatusCodeRequest)
	e.OPTIONS("/:statusCode/*", handleStatusCodeRequest)
	e.HEAD("/:statusCode", handleStatusCodeRequest)
	e.HEAD("/:statusCode/*", handleStatusCodeRequest)

	if *tlsPtr {
		e.Logger.Fatal(e.StartTLS(*addrPtr, *certFilePtr, *keyFilePtr))
	} else {
		e.Logger.Fatal(e.Start(*addrPtr))
	}
}
