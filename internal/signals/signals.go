package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Vak00/goPassM/internal/cli"
	"github.com/Vak00/goPassM/internal/model"
	"github.com/Vak00/goPassM/internal/service"
)

// Start the listening for the system signals
func StartSignalListener() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		cli.CommandQuit(service.NewVaultService([]model.Entry{}))
	}()
}
