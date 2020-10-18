package service

import "testing"

func Test_msgService_CreatePhrase(t *testing.T) {
	tests := []struct {
		name string
		ms   *msgService
		want bool
	}{
		{
			name: "returns true when a phrase is created",
			ms:   &msgService{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &msgService{}
			if got := ms.CreatePhrase(); got != tt.want {
				t.Errorf("msgService.CreatePhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}
