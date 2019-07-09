Elasticlogger is a wrapper for oliveres elastic package

The reason of this package is to enable elasticlogger as an part of the io.writer interface.

Why io.Writer? Because then we can use elasticsearch for a TON of other features, example logging.

It is very easy to use, simple start by creating your elasticlogger
var elasticlogger, _ = elastic.NewElasticLogger($ip, $port, $system)

System variable is the actual index that will be used. Index will be appended with the date or any other flags added
to the stdlib logger.

It is recommended to add date to the Index via flags, since Elasticsearch does not handle 
static indexes very well. A common bug is that upon restarting the cluster, it wont accept updates to an index since it goes into
lock state. Added the date to the index will prevent this, since updates will be on new indexes.

Take your elasticlogger and put it into any wanteed io.Writer, below is a example of stdlib logger
var flags = log.Ldate | log.Lshortfile
var logger = log.New(elasticlogger, "main.go", flags)
logger.Println("erhmagerd")