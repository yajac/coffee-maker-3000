package slack

import (
	"testing"
)

func Test_handleIOTEvent(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"channel", "{\"channel\":\"#richmondcoffee\",\"text\":\"FRESH COFFEE!! - \",\"icon_emoji\":\":coffee:\"}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleIOTEvent()
			if (err != nil) != tt.wantErr {
				t.Errorf("handleIOTEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("handleIOTEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleMadeCoffeeEvent(t *testing.T) {
	type args struct {
		channel  string
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Handle", args{"coffee", "imcewan"}, "{\"channel\":\"#coffee\",\"text\":\"Coffee made by imcewan\",\"icon_emoji\":\":star2:\"}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleMadeCoffeeEvent(tt.args.channel, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleMadeCoffeeEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleMadeCoffeeEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleLeaderBoard(t *testing.T) {
	type args struct {
		channel string
		leaders []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"TestChannel", []string{"Ian 9", "Test 8", "Bah 7"}},
			"{\"channel\":\"#TestChannel\",\"text\":\"*Coffee All Stars* \\n ```1. Ian 9```\\n```2. Test 8```\\n```3. Bah 7```\\n\",\"icon_emoji\":\":star2:\"}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleLeaderBoard(tt.args.channel, tt.args.leaders)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleLeaderBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleLeaderBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
