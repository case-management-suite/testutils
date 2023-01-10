package testutil

import (
	context "context"
	"time"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func FxAppTest(app *fx.App, test func()) {

	// In a typical application, we could just use app.Run() here. Since we
	// don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal().Err(err)
	}

	test()

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal().Err(err)
	}
	// os.Remove("./test_cases.db")
}
