/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/just-umyt/todo/pkg/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete task",
	Long:  ``,
	Run:   delete,
}

func delete(cmd *cobra.Command, args []string) {

	id, _ := cmd.Flags().GetInt("delete")

	db.DB.Delete(&todo{}, id)

	fmt.Printf("Task is deleted, id:%d", id)
}
func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("delete", "i", 0, "Delete task id")
}
