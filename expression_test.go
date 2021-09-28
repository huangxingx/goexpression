package goexpression

import (
	"reflect"
	"testing"
)

func TestExpressExecute(t *testing.T) {
	type args struct {
		express string
		param   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// 加法
		{name: "test_add_1", args: args{express: "1+1"}, want: 2.0},
		{name: "test_add_2", args: args{express: "1.2+1"}, want: 2.2},
		{name: "test_add_3", args: args{express: "1.2+1.2"}, want: 2.4},
		// 括号
		{name: "test_bracket_1", args: args{express: "(1+1)*2"}, want: 4.0},
		// 与
		{name: "test_and_1", args: args{express: "true and true"}, want: true},
		{name: "test_and_2", args: args{express: "true and 0"}, want: false},
		// 或 or ||
		{name: "test_or_1", args: args{express: "true or 0"}, want: true},
		{name: "test_or_2", args: args{express: "1 or 0"}, want: true},
		{name: "test_or_3", args: args{express: "0 or false"}, want: false},
		{name: "test_or_4", args: args{express: "0 or f"}, want: false},
		{name: "test_or_5", args: args{express: "1 or 0 and 0"}, want: true}, // and > or

		// 减法
		{name: "test_sub_1", args: args{express: "100-1-2"}, want: 97.0},
		{name: "test_sub_2", args: args{express: "100 sub 1 sub 2"}, want: 97.0},
		// 除法
		{name: "test_div_1", args: args{express: "100/50"}, want: 2.0},
		{name: "test_div_2", args: args{express: "1/5"}, want: 0.2},
		{name: "test_div_3", args: args{express: "1/(5 * 2)"}, want: 0.1},
		{name: "test_div_4", args: args{express: "1/5 * 2"}, want: 0.4},
		// mix
		{name: "test_mix_1", args: args{express: "1+2*3-(4+5)+10/2+2"}, want: 5.0},
		{name: "test_mix_2", args: args{express: "f or (1 and true)"}, want: true},
		// keywork
		{name: "test_keywork_upper", args: args{express: "f or (1 AND TRUE)"}, want: true},

		// logic
		{name: "test_logic_gt", args: args{express: "2 > 1 and 3> 2"}, want: true},
		{name: "test_logic_ge_true", args: args{express: "2 >= 1 and 2>= 2"}, want: true},
		{name: "test_logic_ge_false", args: args{express: "2 <= 1 and 2>= 2"}, want: false},
		{name: "test_logic_not_true", args: args{express: "!0 and 2>= 2"}, want: true},
		{name: "test_logic_not_false", args: args{express: "!1 and 2>= 2"}, want: false},

		// --- 参数测试
		// param
		{name: "test_param_1", args: args{express: "a + 1", param: map[string]interface{}{"a": 1}}, want: 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			express := NewExpress(tt.args.express)
			if got := express.Execute(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExpress().Execute(nil) = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	express := NewExpress("1+1+2")
	for i := 0; i < b.N; i++ {
		express.Execute(nil)
	}
}

func BenchmarkDiv(b *testing.B) {
	express := NewExpress("10/5/2")
	for i := 0; i < b.N; i++ {
		express.Execute(nil)
	}
}
