package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	Logs, err := os.OpenFile("TaskListLogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Cant Find Logfile is probably because it doesnt exist yet. \n i have no idea how to work around this but it shouldnt cause any problems")
	}

	InfoLogger = log.New(Logs, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(Logs, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(Logs, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialises Task List",
	Long:  "initialises Task List",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("initialising Log File")
		LogFile, err := os.Create("TaskListLogger.log")
		if err != nil {
			defer LogFile.Close()
			log.Fatal("couldnt create Log File")
			log.Fatal(err)
			panic("An Error has occured")
		} else {
			log.Println("Created Log File")
			log.Println("Closing Log File")
			defer LogFile.Close()
		}

		log.Println("initialising Task list...")
		TaskList, err := os.Create("TaskList.csv")
		if err != nil {
			defer TaskList.Close()
			log.Fatal("couldnt create Task List")
			ErrorLogger.Fatal("couldnt create Task List:", err)
			panic("An Error has occured")
		} else {
			InfoLogger.Println("Created Task List")
			InfoLogger.Println("Closing Task File")
			defer TaskList.Close()
		}

		log.Println("Creating Top Row")
		TopRow := []string{"Task", "Date", "Status"}

		log.Println(TopRow)

		log.Println("Creating Writer")
		writer := csv.NewWriter(TaskList)

		log.Println("Flushing Writer")
		defer writer.Flush()

		log.Println("Writing Top Row")
		writer.Write(TopRow)

		log.Println("Finished init")

	},
}
