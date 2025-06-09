package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
	rootCmd.MarkFlagRequired("Task")

	Logs, err := os.OpenFile("TaskListLogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("The Programm had problems opening the Log file. \n Have you ran init yet?")
	}

	InfoLogger = log.New(Logs, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(Logs, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(Logs, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

var add = &cobra.Command{
	Use:   "add [Task Name]",
	Short: "adds a task to your task list",
	Long:  "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		if Task == "Task" {
			ErrorLogger.Fatalln("Cant add Task to List because it is the Heading")
			panic("Tried to edit Forbidden Row")
		}
		InfoLogger.Println("adding:", Task)

		InfoLogger.Println("Opening Task List File")
		TaskList, err := os.OpenFile("TaskList.csv", os.O_RDWR|os.O_APPEND, os.ModeAppend)
		if err != nil {
			defer TaskList.Close()
			ErrorLogger.Fatal(err)
			panic(err)
		}

		InfoLogger.Println("Closing File")
		defer TaskList.Close()

		Date := time.Now().Format(time.DateTime)

		InfoLogger.Println("Assembling Data")
		Status := "Open"
		TaskRow := []string{Task, Date, Status}
		InfoLogger.Println(TaskRow)

		InfoLogger.Println("Creating Writer")
		writer := csv.NewWriter(TaskList)

		InfoLogger.Println("Flushing Writer")
		defer writer.Flush()

		InfoLogger.Println("Writing Data")
		writer.Write([]string{Task, Date, Status})

	},
}
