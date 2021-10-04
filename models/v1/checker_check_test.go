package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checker_check(T *testing.T)  {
	c := Checker{}

	c.Init("someName")

	res, err := c.Check()
	if err != nil {
		assert.Failf(T, "failureMessage", "msg", err)
	}

	assert.Equal(T, "someName", res )



//	T.Errorf("Ошибка")
//	return
}