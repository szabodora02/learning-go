package closures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	
	
	
	balancedOne := proxy(func()int{return 1},func() int { return 2 })
	balancedTwo := proxy(func()int{return 1},func() int { return 2 })
	assert.Equal(t, 1, balancedOne())
	assert.Equal(t, 1, balancedTwo())
	assert.Equal(t, 2, balancedOne())
	assert.Equal(t, 2, balancedTwo())
	
	
}