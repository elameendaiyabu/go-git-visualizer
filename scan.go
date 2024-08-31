package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

func Scan(folder string) {
	fmt.Printf("Found Folders \n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addRepositoriesToFile(filepath, repositories)
	fmt.Printf("\n\nSuccessfully Added\n\n")
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, folder string) []string {
	// trim the last /

	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

func addRepositoriesToFile(path string, newRepositories []string) {
	existingRepos := parseFilePathToSlice(path)
	repos := joinSlices(newRepositories, existingRepos)
	dumpStringsSliceToFile(repos, path)
}

func parseFilePathToSlice(path string) []string {
	f := openFile(path)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func dumpStringsSliceToFile(repos []string, path string) {
	content := strings.Join(repos, "\n")
	os.WriteFile(path, []byte(content), 0755)
}
