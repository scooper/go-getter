package serve

import (
	"errors"
	"fmt"
	"net/http"
)

// NOTE: use 'context' to capture ^C interrupt for more graceful shutdown

func Serve() {
	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	}
}
