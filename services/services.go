package services

//TEST
func TestServices(s string) string {
	TestMysql("Mysql")
	TestPhp("Php-fpm")
	TestElasticsearch("Elasticsearch")
	TestVarnish("Varnish")
	return string("Hello: " + s)
}

//Test that the Mysql service is running
func TestMysql(s string) string {

	mysqlStatus := MysqlStatus()
	mysqlErrorLogs := MysqlErrorLogs()
	mysqlSlowLogs := MysqlSlowLogs()

	mysql := string("Mysql: " + s + "\n" + mysqlStatus + "\n" + mysqlErrorLogs + "\n" + mysqlSlowLogs)
	return string(mysql)
}

func TestPhp(s string) string {
	phpSystemLogs := PhpSystemLogs()
	phpErrorLogs := PhpErrorLogs()

	php := string("PHP-FPM: " + s + "\n" + phpSystemLogs + "\n" + phpErrorLogs)
	return string(php)
}

func TestElasticsearch(s string) string {

	elasticsearchStatus := ElasticsearchStatus()
	elasticsearchDir := ElasticsearchDirAndDiskSpace()
	elasticsearchLogs := ElasticsearchLogs()

	elasticsearch := string("Elasticsearch: " + s + "\n" + elasticsearchStatus + "\n" + elasticsearchDir + "\n" + elasticsearchLogs)
	return string(elasticsearch)
}

func TestVarnish(s string) string {

	varnishStatus := VarnishStatus()

	varnish := string("Varnish: " + s + "\n" + varnishStatus)
	return string(varnish)
}


// ***** Mysql functions ******

//Check if Mysql service is running
func MysqlStatus() string {
	return string("Mysql Service Status")
}

//Check if Mysql error logs for any issue
func MysqlErrorLogs() string {
	return string("Mysql Error Logs")
}

//Check if Mysql slow logs
func MysqlSlowLogs() string {
	return string("Mysql Slow Logs")
}

// ***** PHP-FPMh functions ******

// Analize php system logs
func PhpSystemLogs() string {
	return string("PHP System Logs")
}

// Analyse php fatal error logs
func PhpErrorLogs() string {
	return string("PHP Error Logs")
}

// ***** Elastic Search functions ******

//Check if Elasticsearch service is running
func ElasticsearchStatus() string {
	return string("Elasticsearch Service Status")
}

//Check disk space and directories
func ElasticsearchDirAndDiskSpace() string {
	return string("Elasticsearch Error Logs")
}

//Check Elasticsearch logs for any issue
func ElasticsearchLogs() string {
	return string("Elasticsearch Slow Logs")
}

// ***** Varnish functions ******
func VarnishStatus() string {
	return string ("Varnish Service Status")
}