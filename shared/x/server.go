package x

import (
	"context"
	"github.com/apex/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Start(handler http.Handler, port int) {
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	log.Info("Server exiting")
}
