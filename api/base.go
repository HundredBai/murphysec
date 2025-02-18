package api

import (
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var CliServerAddressOverride string

// serverAddress returns API logger URL
func serverAddress() string {
	var envServer string
	envServer = strings.Trim(strings.TrimSpace(CliServerAddressOverride), "/")
	if envServer == "" {
		envServer = strings.Trim(strings.TrimSpace(os.Getenv("MPS_CLI_SERVER")), "/")
	}
	if len(envServer) == 0 {
		return "https://www.murphysec.com"
	}
	return envServer
}

var client *http.Client

func init() {
	c := new(http.Client)
	c.Timeout = time.Second * 300
	i, e := strconv.Atoi(os.Getenv("API_TIMEOUT"))
	if e == nil && i > 0 {
		c.Timeout = time.Duration(int64(time.Second) * int64(i))
	}
	client = c
}

var ErrTokenInvalid = errors.New("Token invalid")
var ErrApiTimeout = errors.New("API timeout")
