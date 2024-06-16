/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/just-umyt/todo/pkg/db"
	"github.com/just-umyt/todo/pkg/models"
	"github.com/spf13/cobra"
)

type todo models.Todo

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "to add a new TODO",
	Long:  ``,
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	task, _ := cmd.Flags().GetString("task")

	if len(task) == 0 {
		log.Fatal("New task name is empty")
	}

	completed, err := time.Parse(time.DateTime, "1000-01-01 00:00:00")
	if err != nil {
		panic(err)
	}

	newtodo := todo{
		Task:      task,
		Created:   time.Now(),
		Completed: completed,
	}

	db.DB.Create(&newtodo)

	fmt.Println("Task created")
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("task", "t", "", "New task's name")
}
