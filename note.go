package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: note {filename}")
		return
	}

	// Get the filename from the command-line argument and add .md extension
	filename := os.Args[1] + ".md"

	// Set the directory where notes should be generated
	baseDir := filepath.Join(os.Getenv("HOME"), "Documents", "Notes")

	// Create the full path for the Markdown file
	fullPath := filepath.Join(baseDir, filename)

	// Check if the file exists
	_, err := os.Stat(fullPath)
	if err == nil {
		// File exists, so open and edit it
		err = editExistingMarkdownFile(fullPath)
		if err != nil {
			fmt.Println("Error editing the existing Markdown file:", err)
			return
		}
	} else {
		// File does not exist, so create a new file and edit it
		err = generateMarkdownFile(fullPath)
		if err != nil {
			fmt.Println("Error generating the Markdown file:", err)
			return
		}
		err = editExistingMarkdownFile(fullPath)
		if err != nil {
			fmt.Println("Error editing the new Markdown file:", err)
			return
		}
	}

	fmt.Println("Markdown file generated/opened and edited with Neovim successfully.")
}

func generateMarkdownFile(filename string) error {
	content := []byte("# My Markdown File\n\nWelcome to my Markdown file generated using Go!")

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}

func editExistingMarkdownFile(filename string) error {
	cmd := exec.Command("nvim", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

