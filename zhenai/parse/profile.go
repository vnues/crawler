package parse

import (
	"../../engine"
	"../../model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"strings"
)
var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		//fmt.Printf("json : %s\n",json)
		profile := parseJson(json)
		profile.Name = name
		//fmt.Println(profile)
		result.Items = append(result.Items, profile)
		//fmt.Println(result)
	}

	return result

}

//解析json数据
func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}

	infos, err := res.Get("objectInfo").Get("basicInfo").Array()//判断是否是数组
	//infos是一个切片，里面的类型是interface{}

	//fmt.Printf("infos:%v,  %T\n", infos, infos) //infos:[离异 47岁 射手座(11.22-12.21) 157cm 55kg 工作地:阿坝汶川 月收入:3-5千 教育/科研 大学本科],  []interface {}

	var profile model.Profile
	//所以我们遍历这个切片，里面使用断言来判断类型
	for k, v := range infos {
		//fmt.Printf("k:%v,%T\n", k, k)
		//fmt.Printf("v:%v,%T\n", v, v)

		/*
				 "basicInfo":[
		            "未婚",
		            "25岁",
		            "魔羯座(12.22-01.19)",
		            "152cm",
		            "42kg",
		            "工作地:阿坝茂县",
		            "月收入:3-5千",
		            "医生",
		            "大专"
		        ],
		*/
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}

	}

	infos2, err := res.Get("objectInfo").Get("detailInfo").Array()

	/*
		"detailInfo":
		["汉族",
		"籍贯:江苏宿迁",
		"体型:富线条美",
		"不吸烟",
		"不喝酒",
		"租房",
		"未买车",
		"没有小孩",
		"是否想要孩子:想要孩子",
		"何时结婚:认同闪婚"],
		汉族籍贯:安徽合肥体型:运动员型稍微抽一点烟社交场合会喝酒已购房已买车有孩子且住在一起是否想要孩子:视情况而定何时结婚:一年内
	*/
	for _, v := range infos2 {
		/*
				"detailInfo": ["汉族", "籍贯:江苏宿迁", "体型:富线条美", "不吸烟", "不喝酒", "租房", "未买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:认同闪婚"],
			   因为每个 每个用户的detailInfo数据不同，我们可以通过提取关键字来判断
		*/
		if e, ok := v.(string); ok {
			//fmt.Println(k, "--->", e)
			if strings.Contains(e, "族") {
				profile.Hukou = e
			} else if strings.Contains(e, "房") {
				profile.House = e
			} else if strings.Contains(e, "车") {
				profile.Car = e
			}
		}
	}

	//性别：

	gender, err := res.Get("objectInfo").Get("genderString").String()
	profile.Gender = gender

	//fmt.Printf("%+v\n", profile)
	return profile

}

/*

{"objectInfo":{"age":26,"avatarPhotoID":233747231,"avatarPraiseCount":1,"avatarPraised":false,"avatarURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F298678\u002F1194708821\u002F26273208306731268.jpg","basicInfo":["未婚","26岁","天蝎座(10.23-11.21)","158cm","59kg","工作地:阿坝若尔盖","月收入:3-5千","公务员","大学本科"],"detailInfo":["汉族","籍贯:四川成都","不吸烟","社交场合会喝酒","租房","没有小孩","是否想要孩子:以后再告诉你","何时结婚:时机成熟就结婚"],"educationString":"大学本科","emotionStatus":0,"gender":1,"genderString":"女士","hasIntroduce":true,"heightString":"158cm","hideVerifyModule":false,"introduceContent":"只有适合结婚的人，没有适合结婚的年龄！～～","introducePraiseCount":0,"isActive":false,"isFollowing":false,"isInBlackList":false,"isStar":false,"isZhenaiMail":false,"lastLoginTimeString":"2天前活跃","liveAudienceCount":0,"liveType":0,"marriageString":"未婚","memberID":1194708821,"momentCount":0,"nickname":"静听雨声","objectAgeString":"26-30岁","objectBodyString":"运动员型","objectChildrenString":"没有小孩","objectEducationString":"大学本科","objectHeightString":"168-178cm","objectInfo":["26-30岁","168-178cm","工作地:四川成都","月薪:5千以上","大学本科","未婚","体型:运动员型","不要吸烟","没有小孩"],"objectMarriageString":"未婚","objectSalaryString":"5000元以上","objectWantChildrenString":"未填写","objectWorkCityString":"四川成都","onlive":0,"photoCount":1,"photos":[{"createTime":"2019-04-08 18:56:27","isAvatar":true,"photoID":233747231,"photoType":1,"photoURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F298678\u002F1194708821\u002F26273208306731268.jpg","praiseCount":1,"praised":false,"verified":true}],"praisedIntroduce":false,"previewPhotoURL":"","salaryString":"3001-5000元","showValidateIDCardFlag":false,"totalPhotoCount":1,"validateEducation":false,"validateFace":false,"validateIDCard":false,"videoCount":0,"videoID":0,"workCity":10127174,"workCityString":"阿坝","workProvinceCityString":"阿坝若尔盖"}
*/