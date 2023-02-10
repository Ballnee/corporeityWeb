import service from ".";
interface loginData{
    username:string,
    password:string
}
export function login(data:loginData){
    return service({
        url:"/login",
        method:"post",
        data
    })
}

//学生列表接口
export function getStudent(){
    return service({
        url:"/student",
        method:"get",
        
    })
}