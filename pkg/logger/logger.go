package logger

import (
	"log"
	"os"
)

// Init sets up a basic logger with a custom prefix
// @Description Initializes the application logger with a predefined configuration.
// Sets the output to standard output, adds date, time, and file information to log entries,
// and logs an initialization message to indicate that the logger is ready to use.
// @Success Logs a message "Logger initialized" to standard output.
func Init() {
	log.SetOutput(os.Stdout)                             // Log output set to standard output
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Log format includes date, time, and short file name
	log.Println("Logger initialized")                    // Log initialization message
}
