package parser

import (
	"testing"
	"io/ioutil"
	"zshanjun/go-exercise/crawler/model"
)

func TestParseProfile(t *testing.T) {
	// 不依赖于网络获取的内容
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "惠儿")
	if len(result.Items) > 1 {
		t.Errorf("expect has 1 item but has %d", len(result.Items))
	}

	profile := model.Profile{
		Name:       "惠儿",
		Gender:     "女",
		Age:        50,
		Height:     156,
		Income:     "3000元以下",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售总监",
		Hokou:      "四川阿坝", //户口
		Xinzuo:     "魔羯座",  //星座
		House:      "租房",
		Car:        "未购车",
	}

	if result.Items[0].(model.Profile) != profile {
		t.Errorf("reusult shoud be %v but was %v", profile, result.Items[0])
	}
}
