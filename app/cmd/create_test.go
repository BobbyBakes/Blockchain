package cmd

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/BobbyBakes/CopyrightedPhrases/app/service"
	"github.com/spf13/cobra"
)

func Test_createPhrase(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name                 string
		args                 args
		want                 string
		createPhraseResponse bool
	}{
		{
			name: "creates a phrase",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{},
			},
			want:                 "Copyright Created\n",
			createPhraseResponse: true,
		}, {
			name:                 "logs failed phrase creation",
			args:                 args{},
			want:                 "Copyright Creation Failed\n",
			createPhraseResponse: false,
		},
	}
	for _, tt := range tests {

		service.MsgService = &mockMsgService{
			returnValue: tt.createPhraseResponse,
		}

		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			createPhrase(tt.args.cmd, tt.args.args)

			w.Close()
			os.Stdout = old
			var buf bytes.Buffer
			io.Copy(&buf, r)

			got := buf.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPhrase = %v, want %v", got, tt.want)
			}
		})
	}
}
