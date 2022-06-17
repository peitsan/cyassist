package joshuaRequest

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Header map[string][]string = map[string][]string{
	"User-Agent": {"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"},
}

// MethodGet /* 本包支持的一些方法 */
var MethodGet string = "GET"

// MethodPost /* 本包支持的一些方法 */
var MethodPost string = "POST"

/*支持的参数*/
type RequestOption struct {
	timeout     time.Duration
	params      map[string]string
	header      map[string][]string
	contentType string
	data        string
}

type FuncRequestOption struct {
	f func(option *RequestOption)
}

func (fc *FuncRequestOption) apply(options *RequestOption) {
	fc.f(options)
}

func newFuncRequestOption(f func(option *RequestOption)) FuncRequestOption {
	return FuncRequestOption{f: f}
}

// SetTimeout /*设置链接的超时时长，从发出head到收到全部body的时间，默认为5s*/
func SetTimeout(time time.Duration) FuncRequestOption {
	return newFuncRequestOption(func(option *RequestOption) {
		option.timeout = time
	})
}

// SetParams /*设置GET需要发送的信息*/
func SetParams(data map[string]string) FuncRequestOption {
	return newFuncRequestOption(func(option *RequestOption) {
		option.params = data
	})
}

// SetData /*设置POST需要发送的信息*/
func SetData(data string) FuncRequestOption {
	return newFuncRequestOption(func(option *RequestOption) {
		option.data = data
	})
}

// SetHeader /*设置文件头的格式，默认的header设置了微信浏览器的useragent*/
func SetHeader(header map[string][]string) FuncRequestOption {
	return newFuncRequestOption(func(option *RequestOption) {
		option.header = header
	})
}

// SetContentType /*设置请求body的类型*/
func SetContentType(contentType string) FuncRequestOption {
	return newFuncRequestOption(func(option *RequestOption) {
		option.contentType = contentType

	})
}

//Request
/*opts中支持的参数有 header(默认有微信的useragent)、timeout(默认5秒)、data、contentType(覆盖header中设置)
   如需使用，需要先创建一个[]funcRequestOption数组，成员分别是想要修改的参数,使用Setxxx(修改值)
   e.g.      []FuncRequestOption{
				SetTimeout(2*time.Second),
				SetHeader(map[string][]string{"useragent": "goland"}),
				SetData("something to send"),
				SetContentType
           }
*/
func Request(method string, url string, opts ...FuncRequestOption) (*http.Response, error) {
	op := RequestOption{
		header:  Header,
		timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt.apply(&op)
	}

	if method == MethodGet {
		return get(url, &op)
	} else if method == MethodPost {
		return post(url, &op)
	} else {
		return nil, errors.New("Unknown or Unsupported Method:" + method)
	}

}

func post(murl string, r *RequestOption) (*http.Response, error) {
	//设置请求参数

	req, err := http.NewRequest(http.MethodPost, murl, strings.NewReader(r.data))
	if err != nil {
		panic(err)
	}

	//添加header
	req.Header = r.header

	//添加类型
	req.Header.Add("Content-Type", r.contentType)

	//设置超时
	myClient := &http.Client{Timeout: r.timeout}

	//请求访问
	resp, err := myClient.Do(req)

	//网络错误等
	if err != nil {
		return resp, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return resp, err
	}
	utf8RespBody, err := all2Utf8(respBody)

	if err != nil {
		return resp, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewReader(utf8RespBody))

	println(string(utf8RespBody))

	return resp, err
}

func get(url string, r *RequestOption) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	//添加header
	req.Header = r.header

	//设置超时
	myClient := &http.Client{Timeout: r.timeout}

	//设置请求参数
	q := req.URL.Query()
	for key, values := range r.params {
		q.Add(key, values)
	}
	req.URL.RawQuery = q.Encode()

	//请求访问
	resp, err := myClient.Do(req)

	//网络错误等
	if err != nil {
		return resp, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	err = resp.Body.Close()
	if err != nil {
		return resp, err
	}

	utf8RespBody, err := all2Utf8(respBody)

	if err != nil {
		return resp, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewReader(utf8RespBody))

	println(string(utf8RespBody))

	return resp, err

}

/*判断编码方式是否是GBK*/
func isGBK(data []byte) bool {
	length := len(data)
	for i := 0; i < length; {
		if data[i] <= 0x79 {
			i++
			continue
		} else if data[i] >= 0x81 && data[i] <= 0xfe && data[i+1] >= 40 && data[i+1] <= 0xfe && data[i+1] != 0xf7 {
			i += 2
			continue
		} else {
			return false
		}
	}
	return true
}

/*判断编码是否是utf8*/
func isUtf8(data []byte) bool {
	//检测该字节在第一个0bit之前有几个1bit，由utf-8编码得，几个1表示该字符用几个字节
	//算法未考虑0x11101111这种捣乱的
	preNum := func(detectByte byte) int {
		var mask byte = 0x80
		num := 0
		for i := 0; i < 8; i++ {
			if detectByte&mask == mask {
				num++
				mask = mask >> 1
			} else {
				break
			}
		}
		return num
	}
	length := len(data)
	for i := 0; i < length; {
		//注意这里并不是小于0x7e，因为utf-8编码字符只占一字节时只需满足首位是0
		if data[i]&0x80 == 0x0 {
			i++
			continue
		}
		preLength := preNum(data[i])
		if preLength > 2 {
			//检查紧跟着得几位是否满足编码要求的以10开头
			for j := 1; j < preLength; j++ {
				if data[i+j]&0xc0 == 0x80 {
					continue
				} else {
					return false
				}
			}
			i += preLength
		}
	}
	return true
}

/*
返回0表示utf8编码，1表示gbk编码方式，2表示未知编码方式
*/
func detectEncoding(data []byte) int {
	//必须先判断是否是utf8编码，因为utf8编码方式的编码都是落在gbk范围里面的，不同之处只有utf8多了个字节数判定的字节
	if isUtf8(data) {
		return 0
	} else if isGBK(data) {
		return 1
	} else {
		return 2
	}
}

func all2Utf8(data []byte) ([]byte, error) {
	encoding := detectEncoding(data)

	if encoding == 0 {
		return data, nil
	} else if encoding == 1 {
		return simplifiedchinese.GBK.NewDecoder().Bytes(data)
	} else {
		return data, errors.New("unknown encoding")
	}
}
