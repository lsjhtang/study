package Mytets

import (
	"cc.la/Utils"
	"testing"
)

func Test_str(t *testing.T)  {
	/*str := Utils.Join("aa","bb")
	t.Log(str)*/

	type args struct {
		strs []string
	}
	tests := []struct{
		name string
		args args
		want string
	}{
		{ "abc1",args{[]string{"test","abc"}}, "testabc"},
		{ "abc2",args{[]string{"test","edf"}}, "testedf"},
		{ "abc3",args{[]string{"test","gmh"}}, "testgmh"},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			if got := Utils.Join(tt.args.strs...); got != tt.want{
				t.Errorf("预期%v,结果为%v", tt.want, got)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "abc1",args{[]string{"aa","bb"}}, "aabb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utils.Join(tt.args.strs...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}