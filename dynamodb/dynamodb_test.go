package dynamodb

import (
	"testing"
)

func TestUpdateLastCoffee(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Basic User", args{"testuser"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateLastCoffee(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLastCoffee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Get Users", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUsers(); (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
