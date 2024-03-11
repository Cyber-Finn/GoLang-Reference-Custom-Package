package custom_logging_package //custom package name -> not main

//Note: remember to edit the Go.mod, and declare the package as the actual github/gitlab URL that you want to use/reference in other projects
//	see my go.mod for more info

//I'm only using the "log" package here, but you can have complex dependencies as needed
import (
	"log"
)

// a very basic struct for what my message needs to look like
type message_object struct {
	message      string
	message_type int
	severity     int
}

// remember to make these publicly accessible (by having first letter CAPS)
func Log_info(message message_object) {
	log.Print(message.message, message.message_type, message.severity)
}

func Log_warning(message message_object) {
	log.Print(message.message, message.message_type, message.severity)
}

func Log_error(message message_object) {
	log.Print(message.message, message.message_type, message.severity)
}

func Log_fatal(message message_object) {
	log.Fatal(message.message, message.message_type, message.severity)
}
