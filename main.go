package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var (
	filename    = flag.String("f", "", "the chrome password csv file")
	domainRegex = regexp.MustCompile("[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+.?")
	ipRegex     = regexp.MustCompile("^((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}$")
)

func main() {
	flag.Parse()
	if *filename == "" {
		printUsage()
		os.Exit(1)
	}
	file, e := os.Open(*filename)
	defer file.Close()
	if e != nil {
		log.Fatalf("can't open the file : %s, the error is %+v", *filename, e)
	}
	reader := csv.NewReader(file)
	for {
		//csv文逐行读取
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("read record from file error, error is %+v\n", e)
		}
		if len(record) != 4 {
			continue
		}
		domain, _, account, password := record[0], record[1], record[2], record[3]
		// 非ip的域名进行校验
		domainNoraml := !ipRegex.MatchString(domain) && domainRegex.MatchString(domain) && domain != "localhost"
		credentialsNormal := account != "" && password != ""
		if domainNoraml && credentialsNormal {
			storeKeychain(domain, account, password)
		}
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `csv2keychain
Usage: csv2keychain [-f filename] 
Options:
`)
	flag.PrintDefaults()
}

func storeKeychain(domain string, account string, password string) {

	command := exec.Command("security", "add-internet-password",
		fmt.Sprintf("-s%s", domain),
		fmt.Sprintf("-a%s", account),
		fmt.Sprintf("-w%s", password))

	_, e := command.CombinedOutput()
	if e != nil {
		log.Printf("exec.command error : %+v \n", e)
	} else {
		log.Printf("store keychain =>  domain : %s,  account: %s, password: %s\n", domain, account, password)
	}

}
