package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	autoexe "github.com/mhdiiilham/autoexec"
)

const SLEEP = 10 * time.Second

func main() {
	fmt.Println("auto exec CSGO")

	var (
		source         string = ""
		steamID32      string = ""
		steamDirectory string = ""
		configName     string = ""
	)

	reader := bufio.NewReader(os.Stdin)

	for source == "" || steamID32 == "" || steamDirectory == "" {

		fmt.Println("Please input your argument")
		fmt.Print("Steam ID32: ")
		steamID32, _ = reader.ReadString('\n')
		steamID32 = strings.ReplaceAll(steamID32, "\r\n", "")

		fmt.Print("config source: ")
		source, _ = reader.ReadString('\n')
		source = strings.ReplaceAll(source, "\r\n", "")

		fmt.Print("steam directory: ")
		steamDirectory, _ = reader.ReadString('\n')
		steamDirectory = strings.ReplaceAll(steamDirectory, "\r\n", "")

		fmt.Print("config name (optional): ")
		configName, _ = reader.ReadString('\n')
		configName = strings.ReplaceAll(configName, "\r\n", "")

	}

	fmt.Println("creating config directory your config...")
	path, err := autoexe.CreateCFGDirectory(steamDirectory, steamID32)
	if err != nil {
		fmt.Printf("failed creating usedata directory: %v\n", err)
		time.Sleep(SLEEP)
		os.Exit(1)
	}

	fmt.Println("start downloading config file...")
	resp, err := autoexe.DownloadConfig(source)
	if err != nil {
		fmt.Printf("failed downloading csgo config: %v\n", err)
		time.Sleep(SLEEP)
		os.Exit(1)
	}

	if err := autoexe.SaveConfig(resp, path, configName); err != nil {
		fmt.Printf("failed saving config file to %s due to an error: %v\n", path, err)
		time.Sleep(SLEEP)
		os.Exit(1)
	}

	fmt.Printf("success copied csgo config from %s to %s\n", source, path)
	time.Sleep(SLEEP)
	os.Exit(0)
}
