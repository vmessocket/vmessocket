package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/vmessocket/vmessocket/common/cmdarg"
	"github.com/vmessocket/vmessocket/common/platform"
	"github.com/vmessocket/vmessocket/core"
	_ "github.com/vmessocket/vmessocket/main/all"
	"github.com/vmessocket/vmessocket/main/commands/base"
)

var CmdRun = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} run [-c config.json] [-d dir]",
	Short:       "running with config",
	Long: `running`,
	Run: executeRun,
}

var CmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "print version",
	Long: `Prints the build information`,
	Run: executeVersion,
}

var (
	configFiles cmdarg.Arg
	configDir   string
	version     = flag.Bool("version", false, "Show current version of vmessocket.")
	format      = flag.String("format", "json", "Format of input file.")
)

func dirExists(file string) bool {
	if file == "" {
		return false
	}
	info, err := os.Stat(file)
	return err == nil && info.IsDir()
}

func executeRun(cmd *base.Command, args []string) {
	flag.Parse()
	printVersion()
	configFiles = getConfigFilePath()
	server, err := startVmessocket()
	if err != nil {
		base.Fatalf("Failed to start: %s", err)
	}
	if err := server.Start(); err != nil {
		base.Fatalf("Failed to start: %s", err)
	}
	defer server.Close()
	runtime.GC()
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

func getConfigFilePath() cmdarg.Arg {
	if dirExists(configDir) {
		log.Println("Using confdir from arg:", configDir)
		readConfDir(configDir)
	} else if envConfDir := platform.GetConfDirPath(); dirExists(envConfDir) {
		log.Println("Using confdir from env:", envConfDir)
		readConfDir(envConfDir)
	}
	if len(configFiles) > 0 {
		return configFiles
	}
	if workingDir, err := os.Getwd(); err == nil {
		configFile := filepath.Join(workingDir, "config.json")
		if fileExists(configFile) {
			log.Println("Using default config: ", configFile)
			return cmdarg.Arg{configFile}
		}
	}
	if configFile := platform.GetConfigurationPath(); fileExists(configFile) {
		log.Println("Using config from env: ", configFile)
		return cmdarg.Arg{configFile}
	}
	log.Println("Using config from STDIN")
	return cmdarg.Arg{"stdin:"}
}

func GetConfigFormat() string {
	switch strings.ToLower(*format) {
	case "pb", "protobuf":
		return "protobuf"
	default:
		return "json"
	}
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}

func readConfDir(dirPath string) {
	confs, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range confs {
		if strings.HasSuffix(f.Name(), ".json") {
			configFiles.Set(path.Join(dirPath, f.Name()))
		}
	}
}

func startVmessocket() (core.Server, error) {
	configFiles := getConfigFilePath()
	config, err := core.LoadConfig(GetConfigFormat(), configFiles[0], configFiles)
	if err != nil {
		return nil, newError("failed to read config files: [", configFiles.String(), "]").Base(err)
	}
	server, err := core.New(config)
	if err != nil {
		return nil, newError("failed to create server").Base(err)
	}
	return server, nil
}
