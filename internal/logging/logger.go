package logging

import (
	"fmt"
	"log"
	"os"
)
var debug bool = false
/*
SetOutputFile: opens a file for writing logs to in lieu of stdout
 */
func SetOutputFile (filename *string){
	file, err := os.OpenFile(*filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		message := fmt.Sprintf("Unable to open %s for logging will log to $STDOUT",*filename)
		Info(&message)
		return
	}
	log.SetOutput(file)
}
/*
Debug: generates a log message with the DEBUG prefix
 */
func Debug(message *string){
	if debug{
		output := "DEBUG: " + *message
		log.Println(output)
	}
}
/*
Info: generates a log message with the INFO prefix
 */
func Info(message *string){
	output := "INFO: " + *message
	log.Println(output)
}
/*
Error: generates a log message with the ERROR prefix and exits the program
 */
func Error(message *string){
	output := "ERROR: " + *message
	//basically exit program on error
	log.Fatalln(output)
}
/*
SetDebug: if useDebug is true turns on debug logging globally
 */
func SetDebug(useDebug bool){
	debug=useDebug
}
/*
GetDebug: returns the value of debug : used to prevent generation of DEBUG messages if they won't be logged
 */
func GetDebug()bool{
	return debug
}