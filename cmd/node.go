package cmd

import (
	"log"
	"rune/db"

	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node management",
}

var createNodeCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new node",
	Run: func(cmd *cobra.Command, args []string) {
		hostname, err := cmd.Flags().GetString("hostname")
		if err != nil {
			log.Fatal(err)
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatal(err)
		}
		db.Client.Create(&db.Node{Hostname: hostname, User: user})
		log.Println("Node created")
	},
}

var deleteNodeCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a node",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		db.Client.Delete(&db.Node{}, id)
		log.Println("Node deleted")
	},
}

func init() {
	nodeCmd.AddCommand(createNodeCmd)
	createNodeCmd.Flags().StringP("hostname", "n", "", "Hostname")
	createNodeCmd.Flags().StringP("user", "u", "", "User")
	nodeCmd.AddCommand(deleteNodeCmd)
	rootCmd.AddCommand(nodeCmd)
}
