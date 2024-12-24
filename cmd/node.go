package cmd

import (
	"fmt"
	"log"
	"rune/db"

	"github.com/spf13/cobra"
)

var createNodeCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new node",
	Run: func(cmd *cobra.Command, args []string) {
		db.Client.Create(&db.Node{})
		record := db.Node{}
		db.Client.Find(&record)
		fmt.Println(record)
	},
}

func init() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(createNodeCmd)
}
