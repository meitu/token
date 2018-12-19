package token

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want *Token
	}{
		{
			name: "TestNew",
			args: args{
				key: []byte("key"),
			},
			want: &Token{version: CurrentVesion, key: []byte("key")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*TODO
func createMessage(data []byte) *message {
	m := &message{version: CurrentVesion, createAt: 1545190747, payload: data}
	return m
}
func TestToken_Sign(t *testing.T) {
	token := New([]byte("key"))
	m := createMessage([]byte("bifrost"))
	assert.NotNil(t, m)
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "TestSign",
			args: args{
				data: []byte("bifrost"),
			},
			want: []byte{0x62, 0x69, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x2d, 0x31, 0x35, 0x34, 0x35, 0x31, 0x39, 0x30, 0x33, 0x35, 0x33, 0x2d, 0x31, 0x2d, 0x35, 0x31, 0x31, 0x63, 0x39, 0x32, 0x30, 0x30, 0x37, 0x32, 0x34, 0x63, 0x61, 0x32, 0x62, 0x61, 0x31, 0x66, 0x61, 0x38, 0x38, 0x34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(int64(time.Now().Unix()))
			got, err := token.Sign(tt.args.data)
			assert.NotNil(t, got)
			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}


func TestToken_Verify(t *testing.T) {
	type fields struct {
		version int32
		key     []byte
	}
	type args struct {
		sign []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Token{
				version: tt.fields.version,
				key:     tt.fields.key,
			}
			got, err := t.Verify(tt.args.sign)
			if (err != nil) != tt.wantErr {
				t.Errorf("Token.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Token.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_MarshalBinary(t *testing.T) {
	type fields struct {
		version  int64
		createAt int64
		payload  []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				version:  tt.fields.version,
				createAt: tt.fields.createAt,
				payload:  tt.fields.payload,
			}
			gotData, err := m.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("message.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("message.MarshalBinary() = %v, want %v", gotData, tt.wantData)
			}
		})
	}

}
*/
