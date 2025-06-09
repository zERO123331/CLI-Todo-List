package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)

	Logs, err := os.OpenFile("TaskListLogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("The Programm had problems opening the Log file. \n Have you ran init yet?")
	}

	InfoLogger = log.New(Logs, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(Logs, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(Logs, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

var list = &cobra.Command{
	Use:   "list all Taks",
	Short: "list all Taks",
	Long:  "list all Taks",
	Run: func(cmd *cobra.Command, args []string) {
		InfoLogger.Println("listing Tasks")

		InfoLogger.Println("Opening Task List File")
		TaskList, err := os.OpenFile("TaskList.csv", os.O_RDONLY|os.O_RDONLY, os.ModeAppend)
		if err != nil {
			defer TaskList.Close()
			log.Fatal(err)
			panic(err)
		}
		defer TaskList.Close()
		reader := csv.NewReader(TaskList)

		tablewriter := tabwriter.NewWriter(os.Stdout, 10, 10, 10, ' ', tabwriter.AlignRight|tabwriter.Debug)

		var i int
		for {
			records, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			var Line string
			if i == 0 {
				Line = "ID"
			} else {
				Line = strconv.Itoa(i)
			}

			defer tablewriter.Flush()
			fmt.Fprintln(tablewriter, Line+"\t"+records[0]+"\t"+records[1]+"\t"+records[2])
			i++
		}

	},
}
