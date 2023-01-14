package testutil

import (
	context "context"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func AppFx(t *testing.T, ops fx.Option, fn func(*fxtest.App)) {
	if err := fx.ValidateApp(ops); err != nil {
		log.Error().Err(err).Msg("Failed FX validation")
		t.FailNow()
	}

	app := fxtest.New(t, ops, fx.StartTimeout(5*time.Second), fx.StopTimeout(5*time.Second))

	app = app.RequireStart()

	fn(app)

	app.RequireStop()
}

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
