package service

import (
	"compress/gzip"
	"io"
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"golang_tes/depLearing/model"
	"log"
)

func DownloadString(remoteUrl string,queryValues url.Values) (body []byte,err error){
	client := &http.Client{};
	body = nil;
	uri,err := url.Parse(remoteUrl);
	if(err != nil){
		return ;
	}
	if(queryValues != nil){
		values := uri.Query();
		if(values != nil){
			for k,v := range values {
				queryValues[k] = v;
			}
		}
		uri.RawQuery = queryValues.Encode();
	}
	reqest, err := http.NewRequest("GET",uri.String(),nil);
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8");
	reqest.Header.Add("Accept-Encoding", "gzip, deflate");
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3");
	reqest.Header.Add("Connection", "keep-alive");
	reqest.Header.Add("Host", uri.Host);
	reqest.Header.Add("Referer", uri.String());
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	response, err := client.Do(reqest)
	defer response.Body.Close();
	if(err != nil){
		return ;
	}

	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					panic(err)
				}

				if n == 0 {
					break
				}
				body = append(body,buf...);
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

		}
	}
	return ;
}

func GetStuInfoByVq (appName string,vq string)(*model.YBUserInfo,error)   {

	resp, err := http.PostForm("http://yb.upc.edu.cn:8084/common/user/getStuIdByVq",
		url.Values{ "appName": {appName}, "vq": {vq}})

	if err != nil {
		// handle error
		log.Fatal("访问接口出错")
		panic(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	userInfo := new(model.YBUserInfo)

	if err := json.Unmarshal(body, &userInfo); err != nil {
		fmt.Println("fsdf")
		panic(err.Error())
	}
	return userInfo,err
}

