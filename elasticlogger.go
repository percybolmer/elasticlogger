// Package elasticlogger is a "wrapper" class for elasticlogger, used to simplyfi usage of elastic
// @Author PercyBolmer
// @Version 1.0
package elasticlogger

import (
	"fmt"
	"strings"
	"time"

	elastic "gopkg.in/olivere/elastic.v3"
)

// ElasticLog is a logger that sends its requests to elastic
type ElasticLog struct {
	ServerIP string
	Port     int16
	System   string
	Client   *elastic.Client
}

//IndexExists checks if there is such a index
func (logger *ElasticLog) IndexExists(index string) (exist bool, err error) {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := logger.Client.IndexExists(index).Do()
	if err != nil {
		// Handle error
		return false, err
	}
	return exists, nil
}

// Write a new error to the elastic DB
// Index will be System-YYYY-MM-DD
// This write method makes the ElasticLog a part of the io.writer interface
func (logger *ElasticLog) Write(data []byte) (int, error) {
	t := time.Now()
	s := t.Format("2006-01-02")
	index := fmt.Sprintf("%s-%s", logger.System, s)
	exists, _ := logger.Client.IndexExists(index).Do()
	if !exists {
		// Create a new index.
		createIndex, err := logger.Client.CreateIndex(index).Do()
		if err != nil {
			// Handle error
			return 0, err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
	_, err := logger.Client.Index().
		Index(index).
		Type("error").
		// 		Id("4").
		BodyJson(string(data)).
		Do()
	if err != nil {
		// Handle error
		return 0, err
	}
	return len(data), nil
}

// NewElasticLogger returns a new logger
func NewElasticLogger(ip string, port int16, system string) (elasticLogger *ElasticLog, err error) {

	client, err := elastic.NewClient()
	if err != nil {
		return
	}

	// Ping the Elasticsearch server to get e.g. the version number
	url := fmt.Sprintf("http://%s:%d", ip, port)
	_, _, err = client.Ping(url).Do()
	if err != nil {
		// Handle error
		return
	}

	elasticLogger = &ElasticLog{
		ServerIP: ip,
		Port:     port,
		System:   strings.ToLower(system),
		Client:   client,
	}

	return

}
