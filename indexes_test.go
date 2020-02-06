package main

import (
	"testing"
)

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		datas     []House
		predicate func(House) bool
		want      bool
	}{
		{name: "Testing any return true",
			datas: houses,
			predicate: func(h House) bool {
				if h.Name == "Center1" || h.Name == "Center2" || h.Name == "Center3" {
					return true
				}
				return false
			},
			want: true},
		{name: "Testing any return false",
			datas: houses,
			predicate: func(h House) bool {
				if h.Name == "Center5" {
					return true
				}
				return false
			},
			want: false},
	}

	for _, tt := range tests {
		if got := Any(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s got: %v but want %v", tt.name, got, tt.want)
		}
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		name      string
		datas     []House
		predicate func(House) bool
		want      bool
	}{
		{"TestNone must found item", houses, func(h House) bool {
			if h.Name == "Center1" {
				return true
			}
			return false
		}, false},
		{"TestNone don't found item", houses, func(h House) bool {
			if h.Name == "Center6" {
				return true
			}
			return false
		}, true},
	}

	for _, tt := range tests {
		if got := None(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s TestNone() got: %v but want %v", tt.name, got, tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name      string
		datas     []House
		predicate func(House) bool
		want      int
	}{
		{"TestIndex must found Center4", houses, func(h House) bool {
			if h.Name == "Center4" {
				return true
			}
			return false
		}, 3},
		{"TestIndex must return -1", houses, func(h House) bool {
			if h.Name == "Center5" {
				return true
			}
			return false
		}, -1},
	}

	for _, tt := range tests {
		if got := Index(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s TestIndex() got: %v and want %v", tt.name, got, tt.want)
		}
	}
}

type House struct {
	Name  string
	Price int64
}

var houses = []House{
	{Name: "Center1", Price: 200_000},
	{Name: "Center2", Price: 250_000},
	{Name: "Center3", Price: 300_000},
	{Name: "Center4", Price: 350_000},
}

func TestFind(t *testing.T) {
	want := House{Name: "Center1", Price: 200_000}
	notFind := House{}

	tests := []struct {
		name      string
		datas     []House
		predicate func(House) bool
		want      House
	}{
		{
			name:  "Test Find must return",
			datas: houses,
			predicate: func(h House) bool {
				if h == want {
					return true
				}
				return false
			},
			want: House{Name: "Center1", Price: 200_000},
		},
		{
			name:  "Test Find must return 2",
			datas: houses,
			predicate: func(h House) bool {
				if h != notFind {
					return false
				}
				return true
			},
			want: notFind,
		},
	}

	for _, test := range tests {
		if got := Finds(test.datas, test.predicate); got != test.want {
			t.Errorf("Name: %s got: %v and want %v", test.name, got, test.want)
		}
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		datas     []House
		predicate func(House) bool
		want      bool
	}{
		{"Testing all, false ", houses, func(h House) bool {
			if h.Name == "Center6" {
				return true
			}
			return false
		}, false},
		{"Testing all, true", houses, func(h House) bool {
			if h.Name == "Center1" || h.Name == "Center2" {
				return true
			} else if h.Name == "Center3" || h.Name == "Center4" {
				return true
			}
			return false
		}, true},
	}

	for _, tt := range tests {
		if got := All(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s got: %v and want %v", tt.name, got, tt.want)
		}
	}
}
