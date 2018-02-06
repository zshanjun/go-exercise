package test

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct{
		str string
		strlen int
	} {
		// normal cases
		{"abcabf", 4},
		{"eidldei", 4},

		// edge cases
		{"", 0},
		{"a", 1},
		{"aaaaaa", 1},

		// chines cases
		{"你好你好啊", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.str)
		if actual != tt.strlen {
			t.Errorf("%s expect %d, got %d", tt.str, tt.strlen, actual)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	str := "空空的口袋的大口大口嗯邓肯觉得"
	strlen := 7
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(str)
		if actual != strlen {
			b.Errorf("%s expect %d, got %d", str, strlen, actual)
		}
	}
}
