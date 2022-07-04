package testing

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
} 

func (es *ExampleSuite) TestTrue() {
	es.T().Log("....Running TestTrue")
	es.True(true)
}

func (es *ExampleSuite) TestFalse() {
	es.T().Log("....Running TestFalse")
	es.False(false)
}

func (es *ExampleSuite) SetupSuite() {
	es.T().Log("SetupSuite")
}

func (es *ExampleSuite) TeardownSuite() {
	es.T().Log("TeardownSuite")
}
func (es *ExampleSuite) SetupTest() {
	es.T().Log("SetupTest")
}
func (es *ExampleSuite) TeardownTest() {
	es.T().Log("TeardownTest")
}

func (es *ExampleSuite) BeforeTest(suiteName, testName string) {
	es.T().Log("BeforeTest")
}

func (es *ExampleSuite) AfterTest(suiteName, testName string) {
	es.T().Log("AfterTest")
}