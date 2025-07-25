package wechat

import (
	"reflect"
	"testing"
)

func TestSubscribe(t *testing.T) {
	type args struct {
		send        *SubscribeSend
		accessToken string
	}
	tests := []struct {
		name    string
		args    args
		want    *SubscribeSendResponse
		wantErr bool
	}{
		{
			name: "TestSubscribeSuccess",
			args: args{
				send: &SubscribeSend{
					TemplateId: "YPDTWGzxiTt2nuBCZ0USM0ZD5wH6gvXD74tiRyWOXoc",
					Page:       "/pages/index/index",
					ToUser:     "o32UD7EaKKUvAc8cabEbuROv3x1I",
					Data: map[string]any{
						"phrase2": map[string]string{
							"value": "申请类型值",
						},
						"date7": map[string]string{
							"value": "2025-04-09",
						},
						"name5": map[string]string{
							"value": "用户姓名",
						},
					},
					MiniprogramState: "trial",
					Lang:             "zh_CN",
				},
				accessToken: "91_f-XQv6_6RQP0itUJGQRdg78IWmeTr3x7cf30PApufa8I4CQY3bbFklE7-sDDSlDICtDTA1GE6-dKsujSk0AUxNfdn87uUQ_UxVRLmeSp80qWejt-WIogXlAz1ZEYQJcAHANFY",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Subscribe(tt.args.send, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}
