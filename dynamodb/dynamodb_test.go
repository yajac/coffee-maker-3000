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
		{"Update Last Coffee", args{"testuser"}, false},
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
		{"Get User", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
