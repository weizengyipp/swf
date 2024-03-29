package main

import (
	"io"
	"net/http"
	"net/url"
	"testing"
)

var host = "127.0.0.1"
var port = "8080"

func TestSWFIndex(t *testing.T) {
	url := "http://" + host + ":" + port
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "xxxx")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.StatusCode)
	if response.StatusCode != 200 {
		t.Fatal(response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}

func TestSWFHello(t *testing.T) {
	url := "http://" + host + ":" + port + "/v1/hello"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "xxxx")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.StatusCode)
	if response.StatusCode != 200 {
		t.Fatal(response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}

func TestSWFHelloV2(t *testing.T) {
	url := "http://" + host + ":" + port + "/v2/hello/zengyi"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "xxxx")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.StatusCode)
	if response.StatusCode != 200 {
		t.Fatal(response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}

func TestSWFLogin(t *testing.T) {
	Posturl := "http://" + host + ":" + port + "/v2/login"
	payload := url.Values{"username": {"admin"}, "password": {"admin"}}
	response, err := http.PostForm(Posturl, payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.StatusCode)
	if response.StatusCode != 200 {
		t.Fatal(response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}

func TestSWF404(t *testing.T) {
	Posturl := "http://" + host + ":" + port + "/404"
	payload := url.Values{"username": {"admin"}, "password": {"admin"}}
	response, err := http.PostForm(Posturl, payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.StatusCode)
	if response.StatusCode != 404 {
		t.Fatal(response.StatusCode)
	}
	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}
