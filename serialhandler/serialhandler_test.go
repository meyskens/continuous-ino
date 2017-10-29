package serialhandler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialHandler_Write(t *testing.T) {
	type fields struct {
		errors     []string
		output     string
		cancelFunc context.CancelFunc
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:   "write test",
			fields: fields{},
			args: args{
				p: []byte("unit test"),
			},
			wantN:   9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SerialHandler{
				errors:     tt.fields.errors,
				output:     tt.fields.output,
				cancelFunc: tt.fields.cancelFunc,
			}
			gotN, err := s.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerialHandler.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("SerialHandler.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestSerialHandler_Output(t *testing.T) {
	handler := New(nil)
	handler.Write([]byte("Hello World"))

	assert.Equal(t, "Hello World", handler.Output())
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	handler := New(cancel)
	handler.Write([]byte("TEST OK"))
	handler.Write([]byte{0x07})

	done := ctx.Done()
	<-done
}
