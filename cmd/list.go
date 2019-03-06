// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
// limitations under the License.

package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"

	"log"


)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of memo you created.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetList(); err != nil{
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func GetList()error{
	var err error
	var (
		command string
		description string
	)
	Db,err = sql.Open("sqlite3","memo.sql")
	if err != nil{
		log.Println(err)
		return err
	}
	rows,err := Db.Query("SELECT * FROM Memo;")
	if err != nil{
		panic(err)
		return err
	}
	fmt.Println("command",":","description")
	for rows.Next() {
		err = rows.Scan(&command,&description)
		if err != nil{
			panic(err)
			return err
		}
		fmt.Println(command,":",description)
	}
	return nil
}