/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/mercanil/simple-todo-app/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"sort"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new todo",
	Long:  `add a new todo`,
	Run: func(cmd *cobra.Command, args []string) {
		addTodo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTodo(title string) {
	var todoIndexes []int
	for _, todo := range TodosStored.Todo {
		todoIndexes = append(todoIndexes, todo.Id)
	}
	sort.Ints(todoIndexes)
	var newIndex = todoIndexes[len(todoIndexes)-1] + 1
	TodosStored.Todo = append(TodosStored.Todo, model.Todo{
		Id:        newIndex,
		Title:     title,
		Completed: false,
	})

	writeToFile()
}

func writeToFile() {
	bytes, err := json.Marshal(TodosStored)
	if err != nil {
		os.Exit(1)
	}
	err = ioutil.WriteFile(StorageFile, bytes, 0644)
	if err != nil {
		os.Exit(1)
	}
}
