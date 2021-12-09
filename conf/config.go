package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Port  string
	Url string `yaml:"url"`
}

var Config Conf

func init() {
	file, err := ioutil.ReadFile("./conf/application.yml")
	if err != nil {
		panic(err)
	}
	parseErr := yaml.Unmarshal(file, &Config)

	if parseErr != nil {
		panic(err)
	}
}
