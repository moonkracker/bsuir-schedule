/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"time"

	"github.com/spf13/cobra"

	"bsuir-schedule/database"
	"bsuir-schedule/helpers"
)

// testDbCmd represents the testDb command
var testDbCmd = &cobra.Command{
	Use:   "testDb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewDatabase("bsuir-schedule.db")
		helpers.ErrHandler(err)
		defer db.CloseDatabase()
		err = db.CreateTable("groupSchedule", map[string]string{"id": "INTEGER PRIMARY KEY AUTOINCREMENT", "last_update": "DATE"})
		helpers.ErrHandler(err)
		err = db.CreateTable("teacherSchedule", map[string]string{"id": "INTEGER PRIMARY KEY AUTOINCREMENT", "last_update": "DATE"})
		helpers.ErrHandler(err)
		err = db.InsertInTable("groupSchedule", map[string]string{"last_update": time.Now().Format("2006-01-02")})
		helpers.ErrHandler(err)
	},
}

func init() {
	rootCmd.AddCommand(testDbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
