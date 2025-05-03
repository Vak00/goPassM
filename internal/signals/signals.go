package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Vak00/goPassM/internal/model"
	"github.com/Vak00/goPassM/internal/service"
	"github.com/Vak00/goPassM/internal/store"
)

// Start the listening for the system signals
func StartSignalListener() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		service.SaveVault(store.NewVaultStore([]model.Entry{}))
	}()
}
