package dynamodb

import "testing"

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
