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


	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing your Memo and create database.",
	Run: func(cmd *cobra.Command, args []string){
		makeDB()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func makeDB(){
	var err error
	Db,err = sql.Open("sqlite3","memo.sql")
		if err != nil{
			log.Println(err)
	}
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS Memo( command STRING, description STRING);`)
	_ , err = Db.Exec(q)
	if err != nil {
		log.Printf("err!!%s",err)
	}
}