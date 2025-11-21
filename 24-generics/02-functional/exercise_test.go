package functional

import (

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenericFunction(t *testing.T) {
	
	
	
	
	valuesInt := [][]int{[]int{2,5,8,88,12},[]int{3,4}}
	valuesStr := [][]string{[]string{"Boldog","karacsonyt"},[]string{"te","mocskos","allat"}}
	assert.Equal(t, []int{2,5,8,88,12,3,4},flatten(valuesInt))
	assert.Equal(t, []string{"Boldog","karacsonyt","te","mocskos","allat"},flatten(valuesStr))
	
}
