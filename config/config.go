package config

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ARGOeu/argo-messaging/Godeps/_workspace/src/github.com/spf13/viper"
)

// APICfg holds kafka configuration
type APICfg struct {
	// values
	BindIP      string
	Port        int
	BrokerHosts []string
	StoreHost   string
	StoreDB     string
	Authen      bool
	Author      bool
	Ack         bool
	Cert        string
	CertKey     string
}

// NewAPICfg creates a new kafka configuration object
func NewAPICfg(params ...string) *APICfg {
	cfg := APICfg{}

	// If NewKafkaCfg is called with argument "LOAD" automatically load config
	for _, param := range params {
		if param == "LOAD" {
			cfg.Load()
			return &cfg
		}
	}

	return &cfg
}

// Load the configuration
func (cfg *APICfg) Load() {

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/argo-messaging")
	viper.AddConfigPath(".")

	// Find and read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Errod trying to read the configuration file: %s \n", err))
	}

	// Load Kafka configuration
	cfg.BindIP = viper.GetString("bind_ip")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - bind_ip", cfg.BindIP)
	cfg.Port = viper.GetInt("port")
	log.Printf("%s\t%s\t%s:%d", "INFO", "CONFIG", "Parameter Loaded - port", cfg.Port)
	cfg.BrokerHosts = viper.GetStringSlice("broker_hosts")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - broker_host", cfg.BrokerHosts)
	cfg.StoreHost = viper.GetString("store_host")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - store_host", cfg.StoreHost)
	cfg.StoreDB = viper.GetString("store_db")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - store_db", cfg.StoreDB)
	cfg.Authen = viper.GetBool("use_authentication")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_authentication", cfg.Authen)
	cfg.Author = viper.GetBool("use_authorization")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_authorization", cfg.Author)
	cfg.Ack = viper.GetBool("use_ack")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_ack", cfg.Ack)
	cfg.Cert = viper.GetString("certificate")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - certificate", cfg.Cert)
	cfg.CertKey = viper.GetString("certificate_key")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - certificate_key", cfg.CertKey)

}

// LoadStrJSON Loads configuration from a JSON string
func (cfg *APICfg) LoadStrJSON(input string) {
	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer([]byte(input)))
	// Load Kafka configuration
	cfg.BindIP = viper.GetString("bind_ip")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - bind_ip", cfg.BindIP)
	cfg.Port = viper.GetInt("port")
	log.Printf("%s\t%s\t%s:%d", "INFO", "CONFIG", "Parameter Loaded - port", cfg.Port)
	cfg.BrokerHosts = viper.GetStringSlice("broker_hosts")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - broker_host", cfg.BrokerHosts)
	cfg.StoreHost = viper.GetString("store_host")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - store_host", cfg.StoreHost)
	cfg.StoreDB = viper.GetString("store_db")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - store_db", cfg.StoreDB)
	cfg.Authen = viper.GetBool("use_authentication")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_authentication", cfg.Authen)
	cfg.Author = viper.GetBool("use_authorization")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_authorization", cfg.Author)
	cfg.Ack = viper.GetBool("use_ack")
	log.Printf("%s\t%s\t%s:%t", "INFO", "CONFIG", "Parameter Loaded - use_ack", cfg.Ack)
	cfg.Cert = viper.GetString("certificate")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - certificate", cfg.Cert)
	cfg.CertKey = viper.GetString("certificate_key")
	log.Printf("%s\t%s\t%s:%s", "INFO", "CONFIG", "Parameter Loaded - certificate_key", cfg.CertKey)

}
