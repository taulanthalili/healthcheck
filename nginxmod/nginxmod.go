package nginxmod

//Check nginx Logs files
// func checkNginxLogs() {}

//TEST
func TestNginxmod(s string) string {

	nginxLogsExist := NginxLogsExist()
	nginxAccessLogs := NginxAccessLogs()
	nginxErrorLogs := NginxAccessLogs()

	nginx := string("Nginx: " + s + "\n" + nginxLogsExist + "\n" + nginxAccessLogs + "\n" + nginxErrorLogs)
	return string(nginx)
}

//Check if the log files access.log and error.log exist and are not empty
func NginxLogsExist() string {
	//get location
	//get access.log
	//get error.log
	return string("Nginx Logs Exist and are not empty?")
}

//Analyze the requests
func NginxAccessLogs() string {
	nginxAccessLogs := "Nginx Access Logs\n"
	badRequests := BadRequests()
	top10IPs := Top10IPs()
	top10UserAgentsOrBots := Top10UserAgentsOrBots()
	top10RequestedPages := Top10RequestedPages()
	top10PageResponse5XX := Top10PageResponse5XX()
	top10PageResponse404 := Top10PageResponse404()

	nginxAccessLogs += string(badRequests + "\n" + top10IPs + "\n" + top10UserAgentsOrBots + "\n")
	nginxAccessLogs += string(top10RequestedPages + "\n" + top10PageResponse5XX + "\n" + top10PageResponse404)
	return string(nginxAccessLogs)
}

func NginxErrorLogs() string {
	nginxErrorLog := "Nginx Error Logs\n"
	nginxErrorLogPHP := NginxErrorLogPHP()
	nginxErrorLogVarnish := NginxErrorLogVarnish()

	nginxErrorLog += string(nginxErrorLogPHP + "\n" + nginxErrorLogVarnish)
	return string(nginxErrorLog)
}

// ***** Nginx Access Logs *****

//Check for bad Strings on URL like (char, select, sleep, from etc..)
func BadRequests() string{
	return string("Nginx Access Logs - Bad requests")
}

//Top 10 request IPs
func Top10IPs() string {
	return string("Nginx Access Logs - Top 10 IPs")
}

//Top 10 request User Agent/Bots
func Top10UserAgentsOrBots() string {
	return string("Nginx Access Logs - Top 10 Request User Agent/Bots")
}

//Top 10 Request Pages
func Top10RequestedPages() string {
	return string("Nginx Access Logs - Top 10 Requested Pages")
}

//Top 10: 5XX(500,502,503) Page Responses
func Top10PageResponse5XX() string {
	return string("Nginx Access Logs - Top 10 5XX Page Responses")
}

//Top 10: 404 Page Responses
func Top10PageResponse404() string {
return string("Nginx Access Logs - Top 10 404 Page Responses")
}

// ***** Nginx Error Logs *****

//Check any Nginx Error related to PHP-FPM (php pool port)
func NginxErrorLogPHP() string {
	return string("Nginx Error Log - PHP-FPM pool port")
}

//Check any Nginx Error related Varnish ( 6081 varnish port)
func NginxErrorLogVarnish() string {
	return string("Nginx Error Log - Varnish")
}