package config

import(
  "os"
  "fmt"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

var Config config

func LoadConfig(filePath string) error{
  Config = config{}
  if _, err := os.Stat(filePath); os.IsNotExist(err) {
    return fmt.Errorf("Config file not found")
  }

  configContent, err := ioutil.ReadFile(filePath)
  if err != nil {
    return err
  }

  err = yaml.Unmarshal(configContent, &Config)
  if err != nil {
    return fmt.Errorf("Error parsing config file: %s", err)
  }

  return nil
}
