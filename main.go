package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	outputFile    = flag.String("output", fmt.Sprintf("access_log_%s.log", time.Now().Format("02_01_06_15h04")), "the output file name")
	numLines      = flag.Int("num", 100, "number of lines to generate")
	sleepInterval = flag.Int("sleep", 0, "number of seconds between two log entries")
	size          = flag.Int("size", 0, "size of the output file in MB")
	enableIPv6    = flag.Bool("ipv6", false, "enable ipv6 in log")
	HTTPVersions  = []string{"1.0", "1.1", "2"}

	userAgentList []string
	pathList      []string

	statusCodes = []Choice{
		{Item: "200", Weight: 0.6},
		{Item: "301", Weight: 0.05},
		{Item: "400", Weight: 0.05},
		{Item: "401", Weight: 0.05},
		{Item: "403", Weight: 0.05},
		{Item: "404", Weight: 0.1},
		{Item: "500", Weight: 0.1},
	}
	methods = []Choice{
		{Item: "GET", Weight: 0.5},
		{Item: "POST", Weight: 0.4},
		{Item: "PUT", Weight: 0.05},
		{Item: "DELETE", Weight: 0.05},
	}
)

func main() {
	flag.Parse()

	if *enableIPv6 {
		fmt.Println("[!] Sorry, --ipv6 is not usable right now.")
		return
	}

	if *size != 0 && *numLines != 100 {
		fmt.Println("[!] You have specified both --size and --num, please specify only 1 option to proceed.")
		return
	}

	fmt.Println("[*] initializing data...")
	userAgentFile, err := os.Open("./data/ua.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer userAgentFile.Close()
	scanner := bufio.NewScanner(userAgentFile)
	for scanner.Scan() {
		userAgentList = append(userAgentList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range []string{"./data/filenames.txt", "./data/directories.txt"} {
		pathFile, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer pathFile.Close()
		scanner = bufio.NewScanner(pathFile)
		for scanner.Scan() {
			pathList = append(pathList, "/"+scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("[+] Loaded userAgent, total %d user-agents\n", len(userAgentList))
	fmt.Printf("[+] Loaded pathList, total %d paths\n", len(pathList))

	fmt.Println("\n[!] Your options:")
	fmt.Println("- outputFile:", *outputFile)
	if *size != 0 {
		fmt.Println("- filesize:", *size, "MB")
	} else {
		fmt.Println("- number of log entries:", *numLines)
	}
	fmt.Println("- IPv6 enabled?:", *enableIPv6)
	fmt.Println("- sleep time between 2 log entries (0 means randomized):", *sleepInterval)
	fmt.Println()

	timeNow := time.Now()

	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriterSize(outFile, 16384)
	defer writer.Flush()

	if *size == 0 {
		// --num: write up to n lines of log as specified
		for i := 0; i < *numLines; i++ {
			increment := getIncrement(*sleepInterval)
			timeNow = timeNow.Add(increment)

			log, err := generateLog(timeNow)
			if err != nil {
				fmt.Printf("Error generating log: %s\n", err)
				return
			}

			_, err = writer.WriteString(log.String())
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
		}
	} else {
		// --size: write up to n bytes as specified
		var totalWritten uint64 = 0
		for totalWritten < uint64(*size*1024*1024) {
			increment := getIncrement(*sleepInterval)
			timeNow = timeNow.Add(increment)

			log, err := generateLog(timeNow)
			if err != nil {
				fmt.Printf("Error generating log: %s\n", err)
				return
			}

			data := []byte(log.String())
			byteWritten, err := writer.Write(data)
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
			totalWritten += uint64(byteWritten)
		}
	}
	outfileInfo, _ := outFile.Stat()
	fmt.Printf("[+] Finished writing to file, with final size = %v\n", processFileSize(outfileInfo.Size()))
}

func getIncrement(sleepInterval int) time.Duration {
	if sleepInterval != 0 {
		return time.Second * time.Duration(sleepInterval)
	}
	return time.Second * time.Duration(rand.Intn(3))
}

func generateLog(timeNow time.Time) (Log, error) {
	randomMethod, err := weightedRandom(methods)
	if err != nil {
		return Log{}, err
	}
	randomStatusCode, err := weightedRandom(statusCodes)
	if err != nil {
		return Log{}, err
	}

	return Log{
		ipAddr:     genIPv4(),
		username:   "-",
		datetime:   timeNow.Format("02/Jan/2006:15:04:05"),
		timezone:   timeNow.Format("-0700"),
		method:     randomMethod,
		path:       pathList[rand.Intn(len(pathList))],
		version:    HTTPVersions[rand.Intn(len(HTTPVersions))],
		statusCode: randomStatusCode,
		respLength: rand.Intn(4000),
		referer:    "",
		userAgent:  userAgentList[rand.Intn(len(userAgentList))],
	}, nil
}
