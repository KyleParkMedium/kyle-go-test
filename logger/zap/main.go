package main

import (
	"time"
	"zap/flogging"

	"go.uber.org/zap"
)

// var logger = flogging.MustGetLogger("common.tools.configtxgen")

func main() {

	defer func() {
		if r := recover(); r != nil {
			flogging.Info("Starting the application...")
		}
	}()

	flogging.Info("Starting the application...")

	// // For some users, the presets offered by the NewProduction, NewDevelopment,
	// // and NewExample constructors won't be appropriate. For most of those
	// // users, the bundled Config struct offers the right balance of flexibility
	// // and convenience. (For more complex needs, see the AdvancedConfiguration
	// // example.)
	// //
	// // See the documentation for Config and zapcore.EncoderConfig for all the
	// // available options.
	// rawJSON := []byte(`{
	// 	"level": "debug",
	// 	"encoding": "json",
	// 	"outputPaths": ["stdout", "/tmp/logs"],
	// 	"errorOutputPaths": ["stderr"],
	// 	"initialFields": {"foo": "bar"},
	// 	"encoderConfig": {
	// 	  "messageKey": "message",
	// 	  "levelKey": "level",
	// 	  "levelEncoder": "lowercase"
	// 	}
	//   }`)

	// var cfg zap.Config
	// if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(cfg)
	// logger := zap.Must(cfg.Build())
	// defer logger.Sync()

	// logger.Info("logger construction succeeded")

	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	// logger := zap.NewExample() // or NewProduction, or NewDevelopment
	logger, err := zap.NewProduction()
	if err != nil {
		err.Error()
	}
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// logger.Panic("panic")

	logger.Sugar().Panicf("pac")
}
