package addressBook

import (
	"testing"
)

var correctString = []string{"John Daggett, 341 King Road, Plymouth MA", "Alice Ford, 22 East Broadway, Richmond MA", "Alice Ford, 22 East Broadway, Richmond VA"}
var correctStringAnswer = "Massachusetts\n..... Alice Ford 22 East Broadway Richmond Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia"
var emptyItem = []string{"John Daggett, 341 King Road, Plymouth MA", "Alice Ford, 22 East Broadway, Richmond MA", ""}
var emptyItemAnswer = "Massachusetts\n..... Alice Ford 22 East Broadway Richmond Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts"
var incorrectState = []string{"John Daggett, 341 King Road, Plymouth OO", "Alice Ford, 22 East Broadway, Richmond MA", "Alice Ford, 22 East Broadway, Richmond VA"}

func TestBook_Format(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"CorrectCase", args{list: correctString}, correctStringAnswer, false},
		{"EmptyItem", args{list: emptyItem}, emptyItemAnswer, false},
		{"IncorrectStateName", args{list: incorrectState}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{}
			got, err := b.Format(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
