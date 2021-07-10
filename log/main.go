// package log

// import (
// 	logrus "github.com/sirupsen/logrus"
// 	"os"
// )

// var log *logrus.Logger

// func Init() error {
// 	// Log as JSON instead of the default ASCII formatter.
// 	log.SetFormatter(&log.JSONFormatter{})

// 	// Output to stdout instead of the default stderr
// 	// Can be any io.Writer, see below for File example
// 	log.SetOutput(os.Stdout)

// 	// Only log the warning severity or above.
// 	log.SetLevel(log.WarnLevel)
// 	return nil
// }

// func Log() *logrus.Logger {
// 	return log
// }
