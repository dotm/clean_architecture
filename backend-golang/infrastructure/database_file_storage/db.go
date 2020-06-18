package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//FileStorageDatabase ..
type FileStorageDatabase struct {
	StorageDirectory string
}

//NewFileStorageDatabase ..
func NewFileStorageDatabase(storageDirectory string) FileStorageDatabase {
	directoryWithSuffix := filepath.Join(storageDirectory, "db")
	os.MkdirAll(directoryWithSuffix, os.ModePerm)
	return FileStorageDatabase{StorageDirectory: directoryWithSuffix}
}

//CreateNewOrOverwriteExistingFileWithString ..
func (database FileStorageDatabase) CreateNewOrOverwriteExistingFileWithString(filename, str string) {
	database.CreateNewOrOverwriteExistingFileWithBytes(filename, []byte(str))
}

//CreateNewOrOverwriteExistingFileWithBytes ..
func (database FileStorageDatabase) CreateNewOrOverwriteExistingFileWithBytes(filename string, bytes []byte) {
	path := filepath.Join(database.StorageDirectory, filename)
	err := ioutil.WriteFile(path, bytes, 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

//AppendStringToFile ..
func (database FileStorageDatabase) AppendStringToFile(filename, stringToBeAdded string) {
	path := filepath.Join(database.StorageDirectory, filename)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		f = database.CreateFile(path)
	}
	_, err = fmt.Fprintln(f, stringToBeAdded)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//CreateFile ..
func (database FileStorageDatabase) CreateFile(name string) *os.File {
	path := filepath.Join(database.StorageDirectory, name)
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}

//ReadFromFile ..
func (database FileStorageDatabase) ReadFromFile(name string) []byte {
	path := filepath.Join(database.StorageDirectory, name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error reading from file:", err)
	}

	return bytes
}

//CheckIfFileExistWithName ..
func (database FileStorageDatabase) CheckIfFileExistWithName(filename string) bool {
	dir := database.StorageDirectory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("error reading directory:", err)
	}
	for _, f := range files {
		if f.Name() == filename {
			return true
		}
	}

	return false
}

//GetAllFilePathWithPrefix ..
func (database FileStorageDatabase) GetAllFilePathWithPrefix(prefix string) ([]string, error) {
	dir := database.StorageDirectory
	filesWithPrefix := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return filesWithPrefix, err
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), prefix) {
			filesWithPrefix = append(filesWithPrefix, f.Name())
		}
	}
	return filesWithPrefix, nil
}
