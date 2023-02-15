import axios from "axios";
//创建axios实例
const service = axios.create({
    baseURL:"/path/api/v1",
    timeout:5000,
    headers:{
        "Content-Type":"application/json;charset=utf-8"
    }
})

// 请求拦截

service.interceptors.request.use((config)=>{
    config.headers=config.headers || {}
    if(localStorage.getItem('token')){
        const token_type = "Bearer"
        config.headers.Authorization = token_type+ ' ' + localStorage.getItem('token')    
    }
    return config
})

// 响应拦截
service.interceptors.response.use((res)=>{
    // const code:String = res.data.Msg
    // if (code != "login success" || code != "query all success") {
    //     console.log("++++++++++")
    //     return Promise.reject(res.data)
    // }
    return res.data
},(err)=>{console.log(err)})

export default service

