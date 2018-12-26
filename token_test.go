package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			name: "TestNewCase1",
			args: args{
				key: []byte("key"),
			},
			want: &Token{key: []byte("key")},
		},
		{
			name: "TestNewCase2",
			args: args{
				key: []byte(""),
			},
			want: &Token{key: []byte("")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.key)
			assert.NotNil(t, got)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTokenSign(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSignCase1",
			args: args{
				data: []byte("TestSignCase1"),
			},
		},
		{
			name: "TestSignCase2",
			args: args{
				data: []byte("TestSignCase2"),
			},
		},
		{
			name: "TestSignCase2",
			args: args{
				data: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.Sign(tt.args.data)
			assert.NotNil(t, got)
			assert.NoError(t, err)

			err = token.Verify(got)
			assert.NoError(t, err)
		})
	}
}

func TestTokenVerify(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	type args struct {
		sign []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "TestVerifyCase1",
			args: args{
				sign: []byte("TestVerifyCase1"),
			},
			wantErr: nil,
		},
		{
			name: "TestVerifyCase2",
			args: args{
				sign: []byte(nil),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sign, err := token.Sign(tt.args.sign)
			assert.NotNil(t, sign)
			assert.NoError(t, err)
			err = token.Verify(sign)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestTokenAuth(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	type args struct {
		payload []byte
	}
	type want struct {
		res []byte
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "TestAuthCase1",
			args: args{
				payload: []byte("TestAuthCase1"),
			},
			want: want{
				res: []byte("TestAuthCase1"),
				err: nil,
			},
		},
		{
			name: "TestAuthCase2",
			args: args{
				payload: []byte(nil),
			},
			want: want{
				res: []byte(nil),
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sign, err := token.Sign(tt.args.payload)
			assert.NotNil(t, sign)
			assert.NoError(t, err)
			res, err := token.Auth(sign)
			assert.Equal(t, res, tt.want.res)
			assert.Equal(t, err, tt.want.err)
		})
	}
}

func Test_messageMarshalBinary(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	tests := []struct {
		name     string
		payload  []byte
		wantData string
	}{
		{
			name:     "TestMarshalBinary",
			payload:  []byte("TestMarshalBinary"),
			wantData: "TestMarshalBinary-1545205200-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				version:  CurrentVesion,
				createAt: 1545205200,
				payload:  tt.payload,
			}
			gotData, err := m.MarshalBinary()
			assert.NotNil(t, gotData)
			assert.NoError(t, err)

			assert.Equal(t, string(gotData), tt.wantData)

		})
	}

}
func Test_messageUnMarshalBinary(t *testing.T) {
	token := New([]byte("key"))
	assert.NotNil(t, token)
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			name: "TestMarshalBinary",
			data: []byte("TestMarshalBinary-1545205200-1"),
			want: []byte("TestMarshalBinary"),
		},
		{
			name: "TestMarshalBinary_nil",
			data: []byte("-1545205200-1"),
			want: []byte(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{}
			err := m.UnmarshalBinary(tt.data)
			assert.Equal(t, m.payload, tt.want)
			assert.NoError(t, err)

		})
	}

}
