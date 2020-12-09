package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	serverErr := make(chan error, 1)
	sigC := make(chan os.Signal, 1)

	s := http.Server{Addr: ":8080"}

	eg.Go(func() error {
		go func() {
			serverErr <- s.ListenAndServe()
		}()
		select {
		case err := <-serverErr:
			close(sigC)
			close(serverErr)
			return err
		}
	})

	eg.Go(func() error {
		signal.Notify(sigC,
			syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sigC
		return s.Shutdown(context.TODO())
	})
	fmt.Println("test success......")
	log.Println(eg.Wait())
}
