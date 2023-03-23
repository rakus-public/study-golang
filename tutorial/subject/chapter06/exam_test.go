package chapter06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAhoNumber_Call(t *testing.T) {
	tests := []struct {
		name string
		n    AhoNumber
		want string
	}{
		{
			name: "throw 1",
			n:    1,
			want: "1",
		},
		{
			name: "throw 3",
			n:    3,
			want: "さぁん",
		},
		{
			name: "throw 13",
			n:    13,
			want: "じゅうさぁん",
		},
		{
			name: "throw 19",
			n:    19,
			want: "19",
		},
		{
			name: "throw 131",
			n:    131,
			want: "ひゃくさんじゅういぃち",
		},
		{
			name: "throw 202",
			n:    202,
			want: "202",
		},
		{
			name: "throw 498",
			n:    498,
			want: "よんひゃくきゅうじゅうはぁち",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.n.Call())
		})
	}
}

func TestNabeatsu(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "throw 3",
			args: args{3},
			want: []string{"1", "2", "さぁん"},
		},
		{
			name: "throw 10",
			args: args{10},
			want: []string{"1", "2", "さぁん", "4", "5", "ろぉく", "7", "8", "きゅう", "10"},
		},
		{
			name: "throw 40",
			args: args{40},
			want: []string{
				"1", "2", "さぁん", "4", "5", "ろぉく", "7", "8", "きゅう",
				"10", "11", "じゅうにぃ", "じゅうさぁん", "14", "じゅうごぉ", "16", "17", "じゅうはぁち", "19", "20",
				"にじゅういぃち", "22", "にじゅうさぁん", "にじゅうしぃ", "25", "26", "にじゅうしぃち", "28", "29", "さんじゅう",
				"さんじゅういぃち", "さんじゅうにぃ", "さんじゅうさぁん", "さんじゅうしぃ", "さんじゅうごぉ", "さんじゅうろぉく", "さんじゅうしぃち", "さんじゅうはぁち", "さんじゅうきゅう", "40",
			},
		},
		{
			name: "throw 140",
			args: args{140},
			want: []string{
				"1", "2", "さぁん", "4", "5", "ろぉく", "7", "8", "きゅう", "10",
				"11", "じゅうにぃ", "じゅうさぁん", "14", "じゅうごぉ", "16", "17", "じゅうはぁち", "19", "20",
				"にじゅういぃち", "22", "にじゅうさぁん", "にじゅうしぃ", "25", "26", "にじゅうしぃち", "28", "29", "さんじゅう",
				"さんじゅういぃち", "さんじゅうにぃ", "さんじゅうさぁん", "さんじゅうしぃ", "さんじゅうごぉ", "さんじゅうろぉく", "さんじゅうしぃち", "さんじゅうはぁち", "さんじゅうきゅう", "40",
				"41", "よんじゅうにぃ", "よんじゅうさぁん", "44", "よんじゅうごぉ", "46", "47", "よんじゅうはぁち", "49", "50",
				"ごじゅういぃち", "52", "ごじゅうさぁん", "ごじゅうしぃ", "55", "56", "ごじゅうしぃち", "58", "59", "ろくじゅう",
				"61", "62", "ろくじゅうさぁん", "64", "65", "ろくじゅうろぉく", "67", "68", "ろくじゅうきゅう",
				"70", "71", "ななじゅうにぃ", "ななじゅうさぁん", "74", "ななじゅうごぉ", "76", "77", "ななじゅうはぁち", "79", "80",
				"はちじゅういぃち", "82", "はちじゅうさぁん", "はちじゅうしぃ", "85", "86", "はちじゅうしぃち", "88", "89", "きゅうじゅう",
				"91", "92", "きゅうじゅうさぁん", "94", "95", "きゅうじゅうろぉく", "97", "98", "きゅうじゅうきゅう", "100",
				"101", "ひゃくにぃ", "ひゃくさぁん", "104", "ひゃくごぉ", "106", "107", "ひゃくはぁち", "109", "110",
				"ひゃくじゅういぃち", "112", "ひゃくじゅうさぁん", "ひゃくじゅうしぃ", "115", "116", "ひゃくじゅうしぃち", "118", "119", "ひゃくにじゅう",
				"121", "122", "ひゃくにじゅうさぁん", "124", "125", "ひゃくにじゅうろぉく", "127", "128", "ひゃくにじゅうきゅう", "ひゃくさんじゅう",
				"ひゃくさんじゅういぃち", "ひゃくさんじゅうにぃ", "ひゃくさんじゅうさぁん", "ひゃくさんじゅうしぃ", "ひゃくさんじゅうごぉ", "ひゃくさんじゅうろぉく", "ひゃくさんじゅうしぃち", "ひゃくさんじゅうはぁち", "ひゃくさんじゅうきゅう", "140",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Nabeatsu(tt.args.n))
		})
	}
}