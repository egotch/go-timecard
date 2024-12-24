package utils

import (
  "os"
  "log"
)

// inits logging for the application
// sets up logging to external log file
// returns pointer to Logger struct to be used throughout
func InitLog() (*log.Logger, *log.Logger) {

  logfile, _ := os.Create("./timecardlog.log")
  Logger := log.New(logfile, "", log.LstdFlags)
  LoggerDebug := log.New(logfile, "DEBUG:  ", log.Lshortfile)

  Logger.Println("Logger has been initialized")

  return Logger, LoggerDebug
}

