package reader

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type StubInputServicer struct {
}

func (s StubInputServicer) Data(url string) ([]byte, error) {
	if url != "" {
		return []byte(url), nil
	}
	return nil, fmt.Errorf("Error")
}

func TestFetchToGetByteArray(t *testing.T) {
	reader := Reader{
		Servicer: StubInputServicer{},
	}

	actualResult, err := reader.Fetch("getInput")

	require.NoError(t, err, "Expected no error")
	expectedResult := []byte("getInput")

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestFetchToReturnError(t *testing.T) {
	reader := Reader{
		Servicer: StubInputServicer{},
	}

	actualResult, err := reader.Fetch("")
	expectedResult := "Unable to fetch the data from url : Error"
	require.Equal(t, err.Error(), expectedResult)
	require.Nil(t, actualResult)
}
