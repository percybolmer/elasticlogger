package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	elastic "github.com/percybolmer/elasticlogger"
	filewatcher "github.com/percybolmer/fileWatcher"
)

var elasticlogger, _ = elastic.NewElasticLogger("localhost", 9200, "Users")
var flags = log.Ldate | log.Lshortfile
var logger = log.New(elasticlogger, "main.go", flags)

func main() {

	// Print the lines, view kibana and be amazed
	logger.Println("Erhmagerd")
}

func testFileWatcher() {
	// Create a channel with string
	filechannel := make(chan string)
	// Create a file watcher
	watcher := filewatcher.NewFileWatcher()
	watcher.ChangeExecutionTime(1)
	watcher.ChangeTTL(500)
	go watcher.WatchDirectory(filechannel, "./fileWatcher/")

	for {
		select {
		case newFile := <-filechannel:
			parseXML(newFile)
		}
	}
}

func parseXML(filename string) {
	xmlFile, err := os.Open("./fileWatcher/" + filename)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	buf, _ := ioutil.ReadAll(xmlFile)

	var users Users

	xml.Unmarshal(buf, &users)
	for i := 0; i < len(users.Users); i++ {
		logger.Println(users.Users[i])
	}
}

// Users stores users
type Users struct {
	XMLName xml.Name `xml:"users" json:"users"`
	Users   []User   `xml:"user" json:"user"`
}

// User represents a user of social media
type User struct {
	XMLName xml.Name `xml:"user" json:"user"`
	Type    string   `xml:"type,attr" json:"type"`
	Name    string   `xml:"name" json:"name"`
}
