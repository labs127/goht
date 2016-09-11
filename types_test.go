package goht

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var fakeyaml = `
test_name: "testing"
pre_test:
  url: "http://testing.com"
  method: "POST"
  payload:
    - var1: "testvar1"
    - var2: "testvar2"
    - var3: "testvar3"
  payload_type: "form_encoded"
  code: 200
tests:
  -  url: "http://testing.com"
     method: "GET"
     payload:
       - var1: "testvar1"
       - var2: "testvar2"
       - var3: "testvar3"
     payload_type: "url_encoded"
     code: 200
     workers: 10
`

var fakeyaml2 = `
test_name: "testing"
pre_test:
  url: "http://testing.com"
  method: "POST"
  payload:
    - var1: "testvar1"
    - var2: "testvar2"
    - var3: "testvar3"
  payload_type: "form_encoded"
  code: 200
`

var fakeyaml3 = `
test_name: "testing"
pre_test:
  url: "http://testing.com"
  method: "GET"
  code: 200
`

func TestPayloadMap(t *testing.T) {
	p := make(Payload)
	p["test"] = "testing"

	if len(p) < 1 {
		t.Fail()
		t.Log("Expected result, map should filled")
		t.Log("Given result : ", p)
	}
}

func TestParseData(t *testing.T) {
	d := Data{}
	err := yaml.Unmarshal([]byte(fakeyaml), &d)

	if err != nil {
		t.Error(err)
	}

	if d.Name != "testing" {
		t.Fail()
		t.Log("Expected result : testing")
		t.Log("Given result : ", d.Name)
	}

	if d.PreTest.URL != "http://testing.com" {
		t.Fail()
		t.Log("Expected result : http://testing.com")
		t.Log("Given result : ", d.PreTest.URL)
	}

	if d.PreTest.Method != "POST" {
		t.Fail()
		t.Log("Expected result : POST")
		t.Log("Given result : ", d.PreTest.Method)
	}

	if d.PreTest.PayloadType != "form_encoded" {
		t.Fail()
		t.Log("Expected result : form_encoded")
		t.Log("Given result : ", d.PreTest.PayloadType)
	}

	if d.PreTest.Code != 200 {
		t.Fail()
		t.Log("Expected result : 200")
		t.Log("Given result : ", d.PreTest.Code)
	}

	if len(d.PreTest.Payload) < 3 {
		t.Fail()
		t.Log("Expected result : 3")
		t.Log("Given result : ", len(d.PreTest.Payload))
	}

	payload := d.PreTest.Payload

	if payload[0]["var1"] != "testvar1" {
		t.Fail()
		t.Log("Expected result : testvar1")
		t.Log("Given result : ", payload[0]["var1"])
	}

	if payload[1]["var2"] != "testvar2" {
		t.Fail()
		t.Log("Expected result : testvar2")
		t.Log("Given result : ", payload[0]["var2"])
	}

	if payload[2]["var3"] != "testvar3" {
		t.Fail()
		t.Log("Expected result : testvar3")
		t.Log("Given result : ", payload[0]["var3"])
	}

	if len(d.Tests) < 1 {
		t.Fail()
		t.Log("Expected result : 1")
		t.Log("Given result : ", len(d.Tests))
	}

	if d.Tests[0].Workers != 10 {
		t.Fail()
		t.Log("Expected result : 10")
		t.Log("Given result : ", d.Tests[0].Workers)
	}

	if d.Tests[0].Tests.Method != "GET" {
		t.Fail()
		t.Log("Expected result : GET")
		t.Log("Given result : ", d.Tests[0].Tests.Method)
	}

	if d.Tests[0].Tests.PayloadType != "url_encoded" {
		t.Fail()
		t.Log("Expected result : url_encoded")
		t.Log("Given result : ", d.Tests[0].Tests.PayloadType)
	}
}

func TestParseDataWithoutTests(t *testing.T) {
	d := Data{}
	err := yaml.Unmarshal([]byte(fakeyaml2), &d)

	if err != nil {
		t.Error(err)
	}

	if len(d.Tests) >= 1 {
		t.Fail()
		t.Log("Expected result is empty slices")
		t.Log("Given result : ", len(d.Tests))
	}
}

func TestParseDataWithoutPayload(t *testing.T) {

	d := Data{}
	err := yaml.Unmarshal([]byte(fakeyaml3), &d)

	if err != nil {
		t.Error(err)
	}

	if len(d.PreTest.Payload) >= 1 {
		t.Fail()
		t.Log("Expected result is empty slices")
		t.Log("Given result : ", len(d.PreTest.Payload))
	}

	if d.PreTest.PayloadType != "" {
		t.Fail()
		t.Log("Expected result is empty string")
		t.Log("Given result : ", d.PreTest.PayloadType)
	}
}
