package helpers

import (
	"fmt"
	"github.com/Igor-Koniukhov/bookings/internal/config"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var (
	app *config.AppConfig
	Purple      = "\033[35m"
	Red         = "\033[31m"
	Reset       = "\033[0m"
)


var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
var ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println(Red,"Client error with status of", Purple, status, Reset)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
