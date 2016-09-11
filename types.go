package goht

/*
Supported yaml format:

test_name: <string>
pre_test:
  url: <string>
  method: <string>
  payload:
    - var1: <interface>
    - var2: <interface>
    - varn: <interface>
  payload_type: (url_encoded, form_encoded, json)
  code: <int>
tests:
  - url: <string>
    method: <string>
    payload:
      - var1 : <interface>
      - varn : <interface>
    payload_type: (url_encoded, form_encoded, json)
    code: <int>
    workers: <int>
*/

//Payload used to mapping between key string and their value (interface)
type Payload map[string]interface{}

//StatusCode used to indicate http status code
type StatusCode int

//Workers used to indicate how many concurrent workers should be working together
type Workers int

//Endpoint used to set test data
type Endpoint struct {
	URL         string     `yaml:"url"`
	Method      string     `yaml:"method"`
	Payload     []Payload  `yaml:"payload,omitempty,flow"`
	PayloadType string     `yaml:"payload_type,omitempty"`
	Code        StatusCode `yaml:"code"`
}

//Tests contains many of endpoints and their workers
type Tests struct {
	Tests   Endpoint `yaml:"tests,inline"`
	Workers Workers  `yaml:"workers"`
}

//Data is a valid test data format parsed from yaml
type Data struct {
	Name    string   `yaml:"test_name"`
	PreTest Endpoint `yaml:"pre_test"`
	Tests   []Tests  `yaml:",flow"`
}
