package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"errors"
	"sync"
	"time"
)

var (
	m = make(map[string] int)
	lock sync.Mutex
) 

type JsonReponse  struct{
    Response []bool `json:"Response"`
}

type JsonRequest struct{
	Request []string `json:"Request"`
}

func BcjClient(input []string) (output []bool, err error){
	if len(input) <= 0 {
		err = errors.New("the lenth of input []string <= 0!")
		return
	}
	//output = append(output, true)
	err = nil
	var request JsonRequest
	request.Request = input
	msg, e := json.Marshal(request)
	if e == nil {
		pool := x509.NewCertPool()
		caCertPath := "server.pem"
		caCrt, er := ioutil.ReadFile(caCertPath)
		if er != nil {
			fmt.Println("ReadFile err:", er)
			err = er
			return
		}
		pool.AppendCertsFromPEM(caCrt)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		}

		req := bytes.NewBuffer([]byte(msg))
		reqh, _ := http.NewRequest("POST", "https://127.0.0.1:8081", req)
		reqh.Header.Set("Content-type", "application/json")		 
		client := &http.Client{Transport: tr}
		resp, er := client.Do(reqh)
		if er != nil {
			fmt.Println("Get error:", er)
			err = er
			return
		}
		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		var ress JsonReponse
		if err := json.Unmarshal(body, &ress); err == nil {
			output = append(output, ress.Response...)
		}
	}	
	return
}

func BcjServer(){
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.pem", "server.key", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var reqs JsonRequest
	var ress JsonReponse
	if err := json.Unmarshal(body, &reqs); err == nil{
		fmt.Println("http.HandleFunc handler", reqs)
		lock.Lock()
		for _, r := range reqs.Request {			
			_, ok := m[r]
			if ok{
				ress.Response = append(ress.Response, true)
			}else{
				m[r] = 0
				ress.Response = append(ress.Response, false)
			}
		}
		lock.Unlock()
		ret, _ := json.Marshal(ress)
		fmt.Fprintf(w, string(ret))
	}else{		
		s := "Can not recongnized the Params: "
		s = s + string(body)  
		fmt.Fprintf(w, s)
	}	
}

func main(){
	ProductCerts()
	go BcjServer()
	time.Sleep(3*time.Second)
	var sa  = []string{"abc", "bdc", "dc", "ddc", "fbc", "ec"}
	var output[]bool
	output, _ = BcjClient(sa)
	fmt.Println("exe result: ", output)
	output, _ = BcjClient(sa)
	fmt.Println("exe result: ", output)
	var se []string
	if _, err := BcjClient(se); err != nil {
		fmt.Println(err)
	}
	for{
		// go func(){
			// var sa  = []string{"abc", "bdc", "dc", "ddc", "fbc", "ec"}
			// var output []bool
			// output, _ = BcjClient(sa)
			// fmt.Println("exe result: ", output)
		// }()		
	}
}