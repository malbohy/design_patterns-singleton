package Testing

import (
	"../Singleton"
	"testing"

)


func TestGetInstance(t *testing.T) {
	firstCounter := Singleton.GetInstance()

	if firstCounter == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), but found nil")
	}


	expectedCounter := firstCounter
	secondCounter := Singleton.GetInstance()
	if secondCounter != expectedCounter {
		t.Error("Expected same counter in secondCounter but got different instance")
	}
}


func TestAddOneToCountAcrossAllInstances(t *testing.T){
	firstCounter := Singleton.GetInstance()
	secondCounter := Singleton.GetInstance()

	count := firstCounter.AddOne()

	if count != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", count)
	}

	count2 := secondCounter.AddOne()

	if count2 != 2 {
		t.Errorf("After calling AddOne in second counter, count must be 2 but was %d \n ",count2)
	}

}
