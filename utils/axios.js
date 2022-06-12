import axios from "axios"
import Vue from 'vue'
import baseUrl from './baseurl.js'
Vue.prototype.$http = axios
// http 状态码response拦截器
axios.interceptors.response.use(
    response => {
        console.log('response拦截器响应成功')
        console.log(response)
        return response
    },

    error => {
        console.log('response拦截器响应失败')
        console.log(error)

        if (error.request) {
            console.log("正在请求状态码")
            console.log(error.request)

        } else if (error.response) {
            console.log("访问状态码如下:")
            console.log(error.response.data);
            console.log(error.response.status);
            console.log(error.response.headers);
        }
        if (error && error.response) {
            switch (error.response.status) {
                case 400: error.message = '请求错误(400)'; break;
                case 401: error.message = '无法获取token检验码，请重新登录(401)'; break;//代表用户未携带token发起请求或所携带token过期
                case 403: error.message = '拒绝访问(403)'; break;
                case 404: error.message = '请求出错(404)'; break;
                case 408: error.message = '请求超时(408)'; break
                case 412: error.message = '暂无数据接口访问权限(412)'; break;//代表用户虽然token有效，但是用户并未拥有访问该接口的权限
                case 500: error.message = '服务器错误(500)'; break;
                case 501: error.message = '服务未实现(501)'; break;
                case 502: error.message = '网络错误(502)'; break;
                case 503: error.message = '服务不可用(503)'; break;
                case 504: error.message = '网络超时(504)'; break;
                case 505: error.message = 'HTTP版本不受支持(505)'; break;
                default: error.message = `连接出错(${error.response.status})!`;
            }
        } else {
            error.message = '连接服务器失败!'
            console.log(error.message)
        }
        return Promise.reject(error)
    }
)
export function vget(url, params) {
    return new Promise(function (resolve, reject) {
		let atoken = localStorage.getItem("atoken")
        let turl = baseUrl + url
       axios({
       method:"get",
       url:turl,
       data:data,
       headers:{
        	"Content-Type":'application/json',
       		"Authorization":"Bearer "+atoken
        }
		}).
		then((res) => {
            resolve(res);
			// withCredentials: true
        }).catch((err) => {
            reject(err)
        })
    })
}
export function nullGet(url) {
    return new Promise(function (resolve, reject) {
        let turl = baseUrl + url
		let atoken = localStorage.getItem("atoken")
       axios({
       method:"get",
       url:turl,
       headers:{
        	"Content-Type":'application/json',
       		"Authorization":"Bearer "+atoken
        }
		}).
		then((res) => {
            resolve(res);
			// withCredentials: true
        }).catch((err) => {
            reject(err)
        })
    })
}

export function get(turl) {
    return new Promise(function (resolve, reject) {
       axios({
       method:"get",
       url:turl,
       headers:{
        	"Content-Type":'application/json'
        }
		}).
		then((res) => {
            resolve(res);
			// withCredentials: true
        }).catch((err) => {
            reject(err)
        })
    })
}


export function post(turl, data) {
    return new Promise(function (resolve, reject) {
        axios({
		method:"post",
        url:turl,
		data:data,
		headers:{
			"Content-Type":'application/json',
		}
        }).then((res) => {
            resolve(res);
        }).catch((err) => {
            reject(err)
        })
    })
}

export function npost(url, data) {
    return new Promise(function (resolve, reject) {
        let turl = baseUrl + url
        axios({
		method:"post",
        url:turl,
		data:data,
		headers:{
			"Content-Type":'application/json',
		}
        }).then((res) => {
            resolve(res);
        }).catch((err) => {
            reject(err)
        })
    })
}
export function vpost(url, data) {
    return new Promise(function (resolve, reject) {
        let turl = baseUrl + url
		let atoken = localStorage.getItem("atoken")
      axios({
      method:"post",
      url:turl,
      data:data,
      headers:{
       	"Content-Type":'application/json',
		"Authorization":"Bearer "+atoken
       }
       }).then((res) => {
            resolve(res);
			// withCredentials: true
        }).catch((err) => {
            reject(err)
        })
    })
}
export function put(url, params) {
    return new Promise(function (resolve, reject) {
        let turl = baseUrl + url
        axios({
        method:"put",
        url:turl,
        data:data,
        headers:{
         	"Content-Type":'application/json',
        		"Authorization":"Bearer "+atoken
         }
        }).
		then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
}
export function del(url, params) {
    return new Promise(function (resolve, reject) {
        let turl = baseUrl + url
	   axios({
	   method:"get",
	   url:turl,
	   data:data,
	   headers:{
	    	"Content-Type":'application/json',
	   		"Authorization":"Bearer "+atoken
	    }
	   	})
	   .then((res) => {
            resolve(res)
        }).catch((err) => {
            reject(err)
        })
    })
}

const vtoken=(config) => {
	const token = localStorage.getItem("atoken");
    config.headers = {
        "Content-Type":'application/json',
	    "Token":token,
    }
    return config
}

const ntoken=(config) => {
    config.headers = {
        "Content-Type":'application/json',
    }
    return config   
}

export const settoken=(token)=>{
	return localStorage.setItem('atoken',token)
}
export const gettoken=()=>{
	return localStorage.getItem('atoken')
}
export const rmtoken=()=>{
   return localStorage.removeItem('atoken')
}
