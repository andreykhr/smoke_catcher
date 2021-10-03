package main

type Scenario struct {
	Smokeapi float64 `yaml:"smokeapi"`
	Tests    []struct {
		Test struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Request     struct {
				Server  string `yaml:"server"`
				Headers struct {
					ContentType string `yaml:"content_type"`
					Accept      string `yaml:"accept"`
				} `yaml:"headers"`
				Body string `yaml:"body"`
			} `yaml:"request"`
			Response struct {
				HTTPCode int `yaml:"http_code"`
				Headers  []struct {
					Key string `yaml:"key"`
				} `yaml:"headers"`
				Body struct {
					ValidateStruct string `yaml:"validate_struct"`
				} `yaml:"body"`
			} `yaml:"response"`
			Output struct {
				Code int `yaml:"code"`
			} `yaml:"output"`
		} `yaml:"test"`
	} `yaml:"tests"`
}