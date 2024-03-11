package custom_logging_package //custom package name -> not main

//Note: remember to edit the Go.mod, and declare the package as the actual github/gitlab URL that you want to use/reference in other projects
//	see my go.mod for more info

//I'm only using the "log" package here, but you can have complex dependencies as needed
import (
	"log"
)

// a very basic struct for what my message needs to look like
type Message_object struct {
	Message      string
	Message_type int
	Severity     int
}

// remember to make these publicly accessible (by having first letter CAPS)
func Log_info(message Message_object) {
	log.Print(message.Message, message.Message_type, message.Severity)
}

func Log_warning(message Message_object) {
	log.Print(message.Message, message.Message_type, message.Severity)
}

func Log_error(message Message_object) {
	log.Print(message.Message, message.Message_type, message.Severity)
}

func Log_fatal(message Message_object) {
	log.Fatal(message.Message, message.Message_type, message.Severity)
}
