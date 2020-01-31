package indexes

import (
	"testing"
)

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		datas     []int
		predicate func(int) bool
		want      bool
	}{
		{"Testing any is true", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 93 || i == 14 || i == 23 {
				return true
			}
			return false
		}, true},
		{"Testing any return false", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 11 {
				return true
			}
			return false
		}, false},
	}

	for _, tt := range tests {
		if got := Any(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s TestAny() got: %v but want %v", tt.name, got, tt.want)
		}
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		name      string
		datas     []int
		predicate func(int) bool
		want      bool
	}{
		{"TestNone must found item", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 86 {
				return true
			}
			return false
		}, false},
		{"TestNone don't found item", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 22 {
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
		datas     []int
		predicate func(int) bool
		want      int
	}{
		{"TestIndex must found 86 in index 2", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 86 {
				return true
			}
			return false
		}, 2},
		{"TestIndex must return -1", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 22 {
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

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		datas     []int
		predicate func(int) bool
		want      int
	}{
		{"TestFind must return 14", []int{23, 14, 86, 67, 93}, func(i int) bool {
			if i == 14 {
				return true
			}
			return false
		}, 14},
	}

	for _, tt := range tests {
		if got := Find(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s TestFind() got: %v and want %v", tt.name, got, tt.want)
		}
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		datas     []int
		predicate func(int) bool
		want      bool
	}{
		{"Testing all, false ", []int{21, 33, 45, 12, 93, 87}, func(i int) bool {
			if i == 11 {
				return true
			}
			return false
		}, false},
		{"Testing all, true", []int{93, 93, 24, 93, 93}, func(i int) bool {
			if i == 93 || i == 24 {
				return true
			}
			return false
		}, true},
	}

	for _, tt := range tests {
		if got := All(tt.datas, tt.predicate); got != tt.want {
			t.Errorf("Name: %s, TestAll() got: %v and want %v", tt.name, got, tt.want)
		}
	}
}
