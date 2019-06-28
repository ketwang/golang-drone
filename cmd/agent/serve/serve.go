package serve

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"util/cmd/agent/httpd"
	"util/pkg/singal"
)

var (
	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "a http server and other server",
		RunE:  serve,
	}
)

func serve(cmd *cobra.Command, args []string) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile("/backend/golang-drone/cmd/agent/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	fmt.Println(viper.Get("listen.port"))

	ctx := singal.WithSignalsContext(context.Background())

	errChan := make(chan error)

	go func() {
		server := httpd.NewRestfulServer(":9998", logger)
		errChan <- server.ListenAndServe()
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
	}

	return nil
}
