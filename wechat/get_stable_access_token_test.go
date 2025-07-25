package wechat

import (
	"testing"
)

func TestGetStableAccessToken(t *testing.T) {
	type args struct {
		appid  string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    *getStableAccessTokenResp
		wantErr bool
	}{
		{
			name: "TestGetStableAccessToken",
			args: args{
				appid:  "wxdf5d96dc2486fefa",
				secret: "6204a7656de55838c41dc360254cda1d",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStableAccessToken(tt.args.appid, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStableAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("GetStableAccessToken() got = %v", got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetStableAccessToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
