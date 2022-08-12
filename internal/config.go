package internal

import (
	"flag"
	"fmt"
	"github.com/mannemsolutions/pgtester/pkg/pg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

/*
 * This module reads the config file and returns a config object with all entries from the config yaml file.
 */

const (
	envConfName     = "PGCBL_CONFIG"
	defaultConfFile = "/etc/pgcbl/config.yaml"
)

var (
	debug      bool
	version    bool
	server     bool
	configFile string
)

type PCTConfigHost struct {
	IPs      []net.IP
	RootCert string   `yaml:"rootCert"`
	Ports    []uint32 `yaml:"ports"`
}

type PCTConfigHosts map[string]PCTConfigHost

type PCTConfigDatabase struct {
	Tables []string `yaml:"tables"`
}

type PCTConfigDatabases map[string]PCTConfigDatabase

type PCTConfigUser struct {
	Password       string             `yaml:"password"`
	ClientCert     string             `yaml:"clientCert"`
	ClientKey      string             `yaml:"clientKey"`
	ClientPassword string             `yaml:"clientPassword"`
	Databases      PCTConfigDatabases `yaml:"databases"`
}

type PCTConfigUsers map[string]PCTConfigUser

type PCTConfig struct {
	Debug      bool           `yaml:"debug"`
	ServerMode bool           `yaml:"serverMode"`
	Hosts      PCTConfigHosts `yaml:"hosts"`
	Users      PCTConfigUsers `yaml:"users"`
}

func (pc PCTConfig) GetDSNs() (dsns []pg.Dsn) {
	for hostname, hostConfig := range pc.Hosts {
		for _, port := range hostConfig.Ports {
			for userName, userConfig := range pc.Users {
				if len(userConfig.Databases) == 0 {
					userConfig.Databases = PCTConfigDatabases{userName: PCTConfigDatabase{}}
				}
				for dbname := range userConfig.Databases {
					dsn := pg.Dsn{
						"host":   hostname,
						"port":   string(port),
						"user":   userName,
						"dbname": dbname,
					}
					if hostConfig.RootCert != "" {
						dsn["sslmode"] = "verify-full"
						dsn["sslrootcert"] = string2File(hostConfig.RootCert)
					}
					if userConfig.ClientCert != "" {
						dsn["sslcert"] = string2File(userConfig.ClientCert)
					}
					if userConfig.ClientKey != "" {
						dsn["sslkey"] = string2File(userConfig.ClientKey)
					}
					if userConfig.ClientPassword != "" {
						dsn["sslkey"] = userConfig.ClientPassword
					}
					dsns = append(dsns, dsn)
				}
			}
		}
	}
	return dsns
}

func ProcessFlags() (err error) {
	if configFile != "" {
		return
	}

	flag.BoolVar(&debug, "d", false, "Add debugging output")
	flag.BoolVar(&version, "v", false, "Show version information")
	flag.BoolVar(&version, "s", false, "Run as server (listening) instead of client (connecting)")

	flag.StringVar(&configFile, "c", os.Getenv(envConfName), "Path to configfile")

	flag.Parse()

	if version {
		//nolint
		fmt.Println(appVersion)
		os.Exit(0)
	}

	if configFile == "" {
		configFile = defaultConfFile
	}

	configFile, err = filepath.EvalSymlinks(configFile)
	return err
}

func NewConfig() (config PCTConfig, err error) {
	var yamlConfig []byte
	if err = ProcessFlags(); err != nil {
		return
	} else if yamlConfig, err = ioutil.ReadFile(configFile); err != nil { // #nosec
		return config, err
	} else if err = yaml.Unmarshal(yamlConfig, &config); err != nil {
		return config, err
	}
	config.Debug = debug
	config.ServerMode = server
	return config, err
}
