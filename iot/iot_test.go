package made

import "testing"

func Test_handleIOTEvent(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"channel", "{\"channel\":\"#richmondcoffee\",\"text\":\"FRESH COFFEE!! - \",\"icon_emoji\":\":coffee:\",\"image_url\":\"\",\"attachments\":null}", false},
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
