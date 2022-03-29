package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server            `mapstructure:"server"`
	Service           `mapstructure:"service"`
	Mongodb           `mapstructure:"mongodb"`
	KafkaBlocketDB    `mapstructure:"kafka_blocketdb"`
	KafkaAffiliate    `mapstructure:"kafka_affiliate"`
	AffiliateCampaign `mapstructure:"affiliate_campaign"`
	APIKey            `mapstructure:"api_key"`
	Encryption        `mapstructure:"encryption"`
}

type Server struct {
	HTTPPort string `mapstructure:"HTTP_PORT" default:""`
	LogLevel string `mapstructure:"LOG_LEVEL" default:""`
	RunEnv   string `mapstructure:"RUN_ENV" default:""`
}

type Service struct {
	UserAdsDomain        string `mapstructure:"USER_ADS_DOMAIN" default:""`
	AdListingDomain      string `mapstructure:"AD_LISTING_DOMAIN" default:""`
	SpineDomain          string `mapstructure:"SPINE_DOMAIN" default:""`
	AccessTradeDomain    string `mapstructure:"ACCESS_TRADE_DOMAIN" default:""`
	SchemaRegistryDomain string `mapstructure:"SCHEMA_REGISTRY_DOMAIN" default:""`
}

type Mongodb struct {
	ConnectionString string `mapstructure:"CONNECTION_STRING" default:""`
	PoolSize         string `mapstructure:"POOL_SIZE" default:""`
}

type KafkaBlocketDB struct {
	Brokers           string `mapstructure:"BROKERS" default:"127.0.0.1"`
	ConsumerGroup     string `mapstructure:"CONSUMER_GROUP" default:""`
	TopicAds          string `mapstructure:"TOPIC_ADS" default:""`
	TopicActionStates string `mapstructure:"TOPIC_ACTION_STATES" default:""`
	InitOffset        string `mapstructure:"INIT_OFFSET" default:""`
}

type KafkaAffiliate struct {
	Brokers       string `mapstructure:"BROKERS" default:"127.0.0.1"`
	ConsumerGroup string `mapstructure:"CONSUMER_GROUP" default:""`
	TopicAds      string `mapstructure:"TOPIC_ADS" default:""`
	InitOffset    string `mapstructure:"INIT_OFFSET" default:""`
}

type AffiliateCampaign struct {
	Enable   bool   `mapstructure:"ENABLE" default:""`
	Category string `mapstructure:"CATEGORY" default:""`
}

type APIKey struct {
	Web    string `mapstructure:"WEB" default:""`
	Mobile string `mapstructure:"MOBILE" default:""`
}

type Encryption struct {
	JWTSecret string `mapstructure:"JWT_SECRET" default:"secret"`
}

var configs Config

const (
	LastOffset = "LastOffset"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error while reading config file, err: ", err)
	}
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Println("error while Unmarshal config, err: ", err)
	}
}

func GetConfiguration() Config {
	return configs
}
