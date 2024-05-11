package main

import "fmt"

type Log struct {
	ipAddr     string
	username   string
	datetime   string
	timezone   string
	method     string
	path       string
	version    string
	statusCode string
	respLength int
	referer    string
	userAgent  string
}

func (l *Log) String() string {
	return fmt.Sprintf("%s - %s [%s %s] \"%s %s HTTP/%s\" %s %d \"%s\" \"%s\"\n",
		l.ipAddr, l.username, l.datetime, l.timezone, l.method, l.path, l.version, l.statusCode, l.respLength, l.referer, l.userAgent)
}
