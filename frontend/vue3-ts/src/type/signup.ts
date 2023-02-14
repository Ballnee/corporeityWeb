export interface FormType {
    username: string
    password: string
    confirmPassword: string
}

export class signupData {
    ruleForm:FormType={
        username:"",
        password:"",
        confirmPassword:""
    }
}