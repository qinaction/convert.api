package utils

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const (
	//消息类型
	MSG_TEXT         = 1 //文字
	MSG_PICTURE      = 2 //图片
	MSG_FILE         = 3 //文件
	MSG_EMOJI        = 4 //emoji
	MSG_AUDIO        = 5 //音频
	MSG_RED_ENVELOPE = 6 //红包
	MSG_LOCATION     = 7 //位置

	//使用状态
	USE_ON  = 1 //使用
	USE_OFF = 2 //关闭

	//报案方式
	IS_CLAIM_ONLINE     = 1 //线上
	NOT_IS_CLAIM_ONLINE = 2 //线下

	//报案类型
	DAN_LS_CLAIM = 1 //单连锁报案
	DUO_LS_CLAIM = 2 //多连锁报案
	GD_CLAIM     = 3 //个单报案
	JP_CLAIM     = 4 //见票报案

	//是否仅限工作日
	IS_WEEK     = 1 //限制
	NOT_IS_WEEK = 2 //不限制
)

var StatusText = map[int]string{
	USE_ON:  "生效中",
	USE_OFF: "关闭中",
}

var IsClaimOnlineText = map[int]string{
	IS_CLAIM_ONLINE:     "线上",
	NOT_IS_CLAIM_ONLINE: "线下",
}

var ClaimTypeText = map[int]string{
	DAN_LS_CLAIM: "单连锁报案",
	DUO_LS_CLAIM: "多连锁报案",
	GD_CLAIM:     "个单报案",
	JP_CLAIM:     "见票报案",
}

var IsWeekText = map[int]string{
	IS_WEEK:     "限制",
	NOT_IS_WEEK: "不限制",
}

//http 请求方法
func Request(url string, method string, data string) (result string, err error) {
	client := &http.Client{}
	//jsonData, _ := json.Marshal(data)
	//fmt.Println(string(jsonData))
	reqest, _ := http.NewRequest(method, url, strings.NewReader(data))
	reqest.Header.Set("Content-Type", "application/json")
	response, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		bodystr := string(body)
		//fmt.Println(string(bodystr))
		return bodystr, nil
	}
	return "", err
}

//http xml请求方法
func XMLRequest(url string, method string, data string) (result string, err error) {
	//fmt.Println(data)
	client := &http.Client{}
	reqest, _ := http.NewRequest(method, url, strings.NewReader(data))
	reqest.Header.Set("Content-Type", "text/xml")
	response, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		bodystr := string(body)
		//fmt.Println(string(bodystr))
		return bodystr, nil
	}
	return "", err
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数

func Paginator(page, prepage int, nums int) map[string]interface{} {

	var firstpage int //前一页地址

	var lastpage int //后一页地址

	//根据nums总数，和prepage每页数量 生成分页总数

	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数

	if page > totalpages {

		page = totalpages

	}

	if page <= 0 {

		page = 1

	}

	var pages []int

	switch {

	case page >= totalpages-5 && totalpages > 5: //最后5页

		start := totalpages - 5 + 1

		firstpage = page - 1

		lastpage = int(math.Min(float64(totalpages), float64(page+1)))

		pages = make([]int, 5)

		for i, _ := range pages {

			pages[i] = start + i

		}

	case page >= 3 && totalpages > 5:

		start := page - 3 + 1

		pages = make([]int, 5)

		firstpage = page - 3

		for i, _ := range pages {

			pages[i] = start + i

		}

		firstpage = page - 1

		lastpage = page + 1

	default:

		pages = make([]int, int(math.Min(5, float64(totalpages))))

		for i, _ := range pages {

			pages[i] = i + 1

		}

		firstpage = int(math.Max(float64(1), float64(page-1)))

		lastpage = page + 1

		//fmt.Println(pages)

	}

	paginatorMap := make(map[string]interface{})

	paginatorMap["first"] = int(1)

	paginatorMap["before"] = firstpage

	paginatorMap["current"] = page

	paginatorMap["last"] = totalpages

	paginatorMap["next"] = lastpage

	paginatorMap["totalPages"] = totalpages

	paginatorMap["totalItems"] = nums

	paginatorMap["limit"] = prepage

	paginatorMap["pages"] = pages

	return paginatorMap

}

func CalculateFloat64(float float64) (res float64, err error) {

	strfloat := strconv.FormatFloat(float, 'f', 2, 64)
	res, err = strconv.ParseFloat(strfloat, 64)
	return res, err
}

//ToJson 转为json格式
func ToJson(data interface{}) string {
	jsonData, _ := jsoniter.Marshal(data)
	return string(jsonData)
}