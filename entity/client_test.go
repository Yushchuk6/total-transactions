package entity_test

import (
	"net/url"
	"testing"

	"github.com/Yushchuk6/total-transactions/entity"
)

func TestNewClient(t *testing.T) {
	c := entity.NewClient("https", "test.com", "TestKey")
	u, _ := url.Parse("https://test.com/api?apikey=TestKey")
	want := entity.Client{URL: *u}
	if *c != want {
		t.Errorf("Error creating request URL, want %v got %v", want, *c)
	}
}

func TestAddParams(t *testing.T) {
	c := entity.NewClient("https", "test.com", "TestKey")
	p := map[string]string{
		"param1": "123",
		"param2": "test",
		"param3": "098",
	}
	c.AddParams(p)
	want, _ := url.Parse("https://test.com/api?apikey=TestKey&param1=123&param2=test&param3=098")
	if c.URL != *want {
		t.Errorf("Error creating request URL, want %v got %v", *want, c.URL)
	}
}
