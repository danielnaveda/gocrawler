package conf

// Conf is the configuration structure
type Conf struct {
	Domains                  []string `yaml:"domains"`
	API                      string   `yaml:"api"`
	WorkersPerDomain         int      `yaml:"workers_per_domain"`
	MaxPagesCrawledPerDomain int      `yaml:"max_pages_crawled_per_domain"`
	SaveIntoFiles            bool     `yaml:"save_into_files"`
}

// GetConf reads the configuration file
func (c *Conf) GetConf(fileName string, reader func(string) ([]byte, error), unmarshaller func([]byte, interface{}) error) (*Conf, error) {
	yamlFile, err := reader(fileName)

	if err != nil {
		return nil, err
	}

	err = unmarshaller(yamlFile, c)

	return c, err
}
