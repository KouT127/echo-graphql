package develop

import (
	. "github.com/labstack/echo/v4/middleware"
	"os"
)

var (
	DebugLoggerConfig = LoggerConfig{
		Skipper: DefaultSkipper,
		Format: `[${method}:${status}] ${uri} [${latency_human}] bytes_in:${bytes_in} bytes_out:${bytes_out}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           os.Stdout,
	}
)

