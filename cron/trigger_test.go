package cron

import (
	"io/ioutil"
	"encoding/json"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

func getJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

const testConfig string = `{
  "id": "cron1",
  "settings": {
    "expression": "* * * * * *"
  },
  "handlers": [
  ]
}`

func TestCreate(t *testing.T) {

	// New factory
	md := trigger.NewMetadata(getJsonMetadata())
	f := NewFactory(md)

	if f == nil {
		t.Fail()
	}

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	trg := f.New(&config)

	if trg == nil {
		t.Fail()
	}
}
