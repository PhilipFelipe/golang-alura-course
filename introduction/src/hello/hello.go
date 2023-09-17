package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {
	showIntroduction()
	for { // é o "while" do Go
		showOptions()
		command := getCommandInput()

		switch command { // Não precisa do break
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command unrecognized.")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Felipe"
	version := 1.1
	fmt.Println("Hello", name)
	fmt.Println("Current version", version)
}

func showOptions() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func getCommandInput() int {
	var command int
	// fmt.Scanf("%d", &command)
	// fmt.Println("The memory address of the command variable is:", &command)
	fmt.Scan(&command) // Sem necessidade de formatar com tipo
	fmt.Println("Select command:", command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	// var urls [3]string // Array
	// urls[0] = "https://www.alura.com.br"
	// urls[1] = "https://www.google.com.br"
	// urls[2] = "https://www.caelum.com.br"

	urls := readSitesFromFile()

	for i := 0; i < monitorings; i++ {
		for i, url := range urls {
			fmt.Println("Testing site", i, ":", url)
			testSite(url)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func showNames() {
	names := []string{"Felipe", "Ana", "Alan", "Breno"} // Slice
	names = append(names, "Brenda")
	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
}

func testSite(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("An error ocurred:", err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(url, "is offline. Code:", resp.StatusCode)
		registerLog(url, false)
		return
	}
	fmt.Println(url, "is online")
	registerLog(url, true)
}

func readSitesFromFile() []string {
	var sites []string

	// file, err := ioutil.ReadFile("sites.txt")
	file, err := os.Open("sites.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("An error ocurred:", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}
	}

	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("An error ocurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
}

func showLogs() {
	file, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error ocorred:", err)
	}
	fmt.Println(string(file))
}
