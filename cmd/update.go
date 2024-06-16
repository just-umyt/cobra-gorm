/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/just-umyt/todo/pkg/db"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "To change name of task",
	Long:  ``,
	Run:   update,
}

func update(cmd *cobra.Command, args []string) {
	task := todo{}
	id, _ := cmd.Flags().GetInt("update")
	newName, _ := cmd.Flags().GetString("newName")
	done, _ := cmd.Flags().GetBool("done")

	db.DB.First(&task, id)

	if done {
		task.Done = true
	}

	if len(newName) > 0 {
		task.Task = newName
	}

	db.DB.Save(&task)

	fmt.Printf("Task with id: %d, update name: %s, done: %t", id, newName, done)
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("newName", "n", "", "New task name")
	updateCmd.Flags().IntP("update", "i", 0, "to update task ID")
	updateCmd.Flags().BoolP("done", "d", false, "to done task")
}
