package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

type Log struct {
	ipAddr     string
	username   string
	datetime   string
	timezone   string
	method     string
	path       string
	version    string
	statusCode int
	respLength int
	referer    string
	userAgent  string
}

func (l *Log) String() string {
	return fmt.Sprintf("%s - %s [%s %s] \"%s %s HTTP/%s\" %d %d \"%s\" \"%s\"\n",
		l.ipAddr, l.username, l.datetime, l.timezone, l.method, l.path, l.version, l.statusCode, l.respLength, l.referer, l.userAgent)
}

func main() {
	var (
		outputFile    = flag.String("output", fmt.Sprintf("access_log_%s.log", time.Now().Format("02_01_06_15h04")), "the output file name")
		numLines      = flag.Int("num", 100, "number of lines to generate")
		sleepInterval = flag.Int("sleep", 0, "number of seconds between two log entries")
		size          = flag.Int("size", 0, "size of the output file in MB")
		statusCodes   = []int{200, 201, 202, 204, 301, 302, 304, 400, 401, 403, 404, 500}
		methods       = []string{"GET", "POST", "PUT", "DELETE"}
		HTTPVersions  = []string{"1.0", "1.1", "2"}

		userAgentList []string
		pathList      []string
	)

	flag.Parse()

	fmt.Println("[*] initializing data...")
	// prepare data
	file, err := os.Open("./data/ua.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		userAgentList = append(userAgentList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range []string{"./data/filenames.txt", "./data/directories.txt"} {
		file, err = os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			pathList = append(pathList, "/"+scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("[+] loaded ua completed, loaded %d user-agents\n", len(userAgentList))
	fmt.Printf("[+] loaded path completed, loaded %d paths\n", len(pathList))

	fmt.Println("[!] outputFile:", *outputFile)
	timeNow := time.Now()

	outFile, err := os.Create("./output/" + *outputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer outFile.Close()

	if *size == 0 {
		for i := 0; i < *numLines; i++ {
			var increment time.Duration
			if *sleepInterval != 0 {
				increment = time.Second * time.Duration(*sleepInterval)
			} else {
				randomSeconds := rand.Intn(30)
				increment = time.Second * time.Duration(randomSeconds)
			}
			timeNow = timeNow.Add(increment)
			log := Log{
				ipAddr:     genIPv4(),
				username:   "-",
				datetime:   timeNow.Format("02/Jan/2006:15:04:05"),
				timezone:   timeNow.Format("-0700"),
				method:     methods[rand.Intn(len(methods))],
				path:       pathList[rand.Intn(len(pathList))],
				version:    HTTPVersions[rand.Intn(len(HTTPVersions))],
				statusCode: statusCodes[rand.Intn(len(statusCodes))],
				respLength: rand.Intn(4000),
				referer:    "",
				userAgent:  userAgentList[rand.Intn(len(userAgentList))],
			}

			_, err := outFile.WriteString(log.String())
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
		}
	} else {
		var totalWritten uint64 = 0
		for totalWritten < uint64(*size*1024*1024) {
			var increment time.Duration
			if *sleepInterval == 4 {
				increment = time.Second * time.Duration(*sleepInterval)
			} else {
				randomSeconds := rand.Intn(271) + 30
				increment = time.Second * time.Duration(randomSeconds)
			}
			timeNow = timeNow.Add(increment)
			log := Log{
				ipAddr:     genIPv4(),
				username:   "-",
				datetime:   timeNow.Format("02/Jan/2006:15:04:05"),
				timezone:   timeNow.Format("-0700"),
				method:     methods[rand.Intn(len(methods))],
				path:       pathList[rand.Intn(len(pathList))],
				version:    HTTPVersions[rand.Intn(len(HTTPVersions))],
				statusCode: statusCodes[rand.Intn(len(statusCodes))],
				respLength: rand.Intn(4000),
				referer:    "",
				userAgent:  userAgentList[rand.Intn(len(userAgentList))],
			}

			data := []byte(log.String())
			byteWritten, err := outFile.Write(data)
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
			totalWritten += uint64(byteWritten)

			if totalWritten >= uint64(*size*1024*1024) {
				break
			}
		}
	}
	fmt.Println("[+] finished writing to file")
}

func genIPv4() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()
	binary.LittleEndian.PutUint32(buf, ip)
	return net.IP(buf).String()
}
