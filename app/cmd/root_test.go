package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	noArgHelpMenu, err := ioutil.ReadFile("./testdata/noArgRes.txt")
	if err != nil {
		panic("error reading test file")
	}
	tests := []struct {
		name    string
		wantErr bool
		want    string
	}{
		{
			name:    "displays help",
			wantErr: false,
			want:    string(noArgHelpMenu),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			if err := Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)

			got := buf.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, wantErr %v", got, tt.want)
			}

		})
	}
}
