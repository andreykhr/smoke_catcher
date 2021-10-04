package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type addTest struct {
	sourceName, sourceExpected string
}

var addTests = []addTest{
	{"someName", "someName"},
	{"anotherName", "anotherName"},
}

func Test_checker_check(T *testing.T)  {
	for _, test := range addTests {
		c := Checker{}

		c.Init(test.sourceName)

		res, err := c.Check()
		if err != nil {
			assert.Failf(T, "failureMessage", "msg", err)
		}

		assert.Equal(T, test.sourceExpected, res)
	}
}