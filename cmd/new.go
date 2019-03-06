//Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitatio
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"

	"log"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create Memo with description",
	Run: func(c *cobra.Command, args []string) {
		cmd,err := c.Flags().GetString("cmd")
		if err != nil{
			fmt.Println(err)
		}
		dsc,err := c.Flags().GetString("dsc")
		if err != nil{
			fmt.Println(err)
		}
		err1 := Create(cmd,dsc)
		if err1 != nil{
			fmt.Println(err1)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("cmd","c","","command")
	newCmd.Flags().StringP("dsc","d","","command's description")
}

func Create(cmd,dsc string)error{
	var err error
	Db,err = sql.Open("sqlite3","memo.sql")
	if err != nil{
		log.Println(err)
		return err
	}
	q := fmt.Sprintf("INSERT INTO Memo (command,description) VALUES (?,?)")
	stmt,err := Db.Prepare(q)
	if err != nil{
		log.Print(err)
		return err
	}
	_,err = stmt.Exec(cmd,dsc)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}