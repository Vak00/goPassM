package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Start the listening for the system signals
func StartSignalListener() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigchan
		fmt.Println()
		fmt.Printf("Signal %s received. TODO: manage the save of the current modifications befroe interrupt\n", sig)

		os.Exit(0)
	}()
}
