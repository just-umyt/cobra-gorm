/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/just-umyt/todo/cmd"
	"github.com/just-umyt/todo/pkg/db"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(db.InitDB)

	cmd.Execute()
}
