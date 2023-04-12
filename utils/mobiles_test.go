package utils_test

import (
	"testing"

	"github.com/ochom/orctic-library/utils"
)

func TestParseMobile(t *testing.T) {
	type args struct {
		mobile string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "712345678",
			},
			want: "254712345678",
		},
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "0712345678",
			},
			want: "254712345678",
		},
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "+254712345678",
			},
			want: "254712345678",
		},
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "112345678",
			},
			want: "254112345678",
		},
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "0112345678",
			},
			want: "254112345678",
		},
		{
			name: "should return a valid mobile number",
			args: args{
				mobile: "+254112345678",
			},
			want: "254112345678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ParseMobile(tt.args.mobile); got != tt.want {
				t.Errorf("ParseMobile() = %v, want %v", got, tt.want)
			}
		})
	}
}
