package main

import(
	"fmt"
	"testing"
	// "net/http/httptest"
	// "net/http"
	// "encoding/json"
)

func TestClient(t *testing.T){
	var se []string
	output, err := BcjClient(se)
	if err != nil {
		if err.Error() != "the lenth of input []string <= 0!" {		
			t.Errorf("the error is not expacted and the err is :%s", err.Error())	
		}else{
			fmt.Printf("the length of output is actual[%d] expacted[0]", len(output))
		}
	}
	go BcjServer()
	var sa = []string{"abc", "bcd", "cde", "def", "abc"}
	output, err = BcjClient(sa)
	var expacted = []bool{false, false, false, false, true}
	var rc = true 
	if len(output) != 5{
		rc = false
	}
	for i, val := range output{
		if val != expacted[i]{
			rc = false
			break
		}
	}
	if !rc {
		t.Errorf("Error the expacted value is: false, false, false, true but the actual value is: ")
		fmt.Println(output)
	}
}
