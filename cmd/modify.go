package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(modify)
	rootCmd.MarkFlagRequired("Task")
	rootCmd.MarkFlagRequired("Status")

	Logs, err := os.OpenFile("TaskListLogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("The Programm had problems opening the Log file. \n Have you ran init yet?")
	}

	InfoLogger = log.New(Logs, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(Logs, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(Logs, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

var modify = &cobra.Command{
	Use:   "modify [Task Name] [New Status]",
	Short: "modifies a Tasks Status",
	Long:  "modify lets you change the Status of a Task to whatever you like it to be",
	Run: func(cmd *cobra.Command, args []string) {

		if Task == "Task" {
			ErrorLogger.Println("Cant modify Task in List because it is the Heading")
			log.Println("invalid Task Name:", Task)
			panic("Tried to edit Forbidden Row")
		}
		InfoLogger.Println("editing:", Task, "to be Status:", Status)

		InfoLogger.Println("Opening Task List File")
		TaskList, err := os.OpenFile("TaskList.csv", os.O_RDWR|os.O_APPEND, os.ModeAppend)
		if err != nil {
			defer TaskList.Close()
			ErrorLogger.Println("An Error Occured while opening the Task List File:", err)
			panic(err)
		}

		InfoLogger.Println("Closing File")
		defer TaskList.Close()

		InfoLogger.Println("Creating Reader")
		reader := csv.NewReader(TaskList)
		InfoLogger.Println("Creating Writer")
		writer := csv.NewWriter(TaskList)

		InfoLogger.Println("Flushing Writer")
		defer writer.Flush()

		InfoLogger.Println("Modifying Target(s) in Todo List")

		OutputBuffer := [][]string{}

		var Changecount int = 0

		for {
			records, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				ErrorLogger.Println("An Error Occured while parsing the Task List for the Task that should be removed:", err)
				panic(err)
			}

			if records[0] == Task {
				records[2] = Status
				Changecount++
			}
			OutputBuffer = append(OutputBuffer, records)

		}
		InfoLogger.Println(OutputBuffer)

		if err := os.Truncate("TaskList.csv", 0); err != nil {
			log.Println("An Error occured while deleting the Old data to replace it. \nThe Data will be written to the file and likely needs to be manually reviewed and fixed")
			writer.WriteAll(OutputBuffer)
			ErrorLogger.Println("couldnt truncate tasklist", err)
			panic(err)
		}

		InfoLogger.Println("Writing to Task List File")
		writer.WriteAll(OutputBuffer)

		if Changecount >= 2 {
			WarnLogger.Println("Modified", Changecount, "Tasks")
			fmt.Println("Modified", Changecount, "Tasks")
		} else if Changecount == 0 {
			WarnLogger.Println("No Tasks found to modify")
			fmt.Println("No Tasks found to modify")
		} else {
			WarnLogger.Println("succesfully modified Task")
			fmt.Println("succesfully modified Task")
		}

	},
}
