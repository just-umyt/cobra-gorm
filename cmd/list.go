/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/just-umyt/todo/pkg/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "to show todo list",
	Long:  ``,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	task := todo{}
	tasks := []todo{}

	d, _ := cmd.Flags().GetBool("done")

	if d {
		task.Done = true
	}

	db.DB.Where(&task).Find(&tasks)

	tableView(tasks)

}

func tableView(todos []todo) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created at"},
			{Align: simpletable.AlignCenter, Text: "Completed at"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, task := range todos {
		i++
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: task.Task},
			{Text: fmt.Sprintf("%t", task.Done)},
			{Text: task.Created.Format(time.DateTime)},
			{Text: task.Completed.Format(time.DateTime)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d todos", len(todos))},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("done", "d", false, "Get list with Done")
}
