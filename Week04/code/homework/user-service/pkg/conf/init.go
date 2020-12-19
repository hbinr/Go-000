package conf

import (
	"os"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/common/yaml"
	"github.com/spf13/pflag"
)

const defaultConfigFile = "../../conf/myConf.yml"

func Init() (config *MyConfig, err error) {
	var (
		b []byte
	)
	pflag.StringP("myConf", "c", "", "choose conf file.")
	pflag.Parse()
	configFile := os.Getenv("conf")
	if configFile == "" {
		configFile = defaultConfigFile
	}
	logger.Infof("conPath ------->:%v\n", configFile)
	if b, err = yaml.LoadYMLConfig(configFile); err != nil {
		logger.Error("conf:yaml.LoadYMLConfig failed", err)
		return
	}
	config = new(MyConfig)
	if err = yaml.UnmarshalYML(b, config); err != nil {
		logger.Error("conf:yaml.UnmarshalYML failed", err)
		return
	}
	return
}
