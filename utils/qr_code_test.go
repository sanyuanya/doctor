package utils

import "testing"

func TestGenerateQRCodeBase64(t *testing.T) {
	type args struct {
		content string
		size    int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				content: "https://www.baidu.com",
				size:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateQRCodeBase64(tt.args.content, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateQRCodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("base64: %s", got)
			// if got != tt.want {
			// 	t.Errorf("GenerateQRCodeBase64() = %v, want %v", got, tt.want)
			// }
		})
	}
}
