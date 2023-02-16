<template>
    <div class="register-container">
        <el-form :model="ruleForm" :rules="rules" ref="formel" label-width="100px" class="ruleForm">
            <el-form-item label="用户名" prop="username">
                <el-input v-model="ruleForm.username" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input type="password" v-model="ruleForm.password" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="confirmPassword">
                <el-input type="password" v-model="ruleForm.confirmPassword" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="register">注册</el-button>
            </el-form-item>
        </el-form>
    </div>

</template>
  
<script lang="ts">
import { useRouter } from 'vue-router';
import { defineComponent,toRefs, ref,reactive} from 'vue'
import { type ElForm, ElFormItem, ElInput, ElButton, } from 'element-plus'
import {signup} from '../request/api'
import {signupData} from "../type/signup"

interface RulesType {
    username: { required: boolean, message: string, trigger: string }[]
    password: { required: boolean, message: string, trigger: string }[]
    confirmPassword: { required: boolean, message: string, trigger: string, validator: (rule: any, value: any, callback: any) => void }[]
}

export default defineComponent({
    // components: {
    //     ElForm,
    //     ElFormItem,
    //     ElInput,
    //     ElButton,
    // },
    setup() {
        const form = reactive(new signupData())
        const router = useRouter();
        const formel = ref<ElForm>();
        const rules = ref({
            username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
            password: [
                { required: true, message: '请输入密码', trigger: 'blur' },
                { min: 6, max: 20, message: '密码长度在6-20个字符之间', trigger: 'blur' },
            ],
            confirmPassword: [
                { required: true, message: '请确认密码', trigger: 'blur' },
                {
                    //可以正常验证
                    validator: (rule, value, callback) => {
                        if (value !== form.ruleForm.password) {
                            callback(new Error('两次输入的密码不一致'))
                        } else {
                            callback()
                        }
                    },
                    trigger: 'blur',
                },
            ],
            email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }],
        });
        
        const register = ()=>{
            if (!formel.value) return;
            // 对表单的内容进行验证
            formel.value.validate((valid:boolean)=>{
                if(valid) {
                    console.log("submit!")
                    signup(form.ruleForm).then(()=>{
                        // console.log(res)
                        alert("注册成功,请登陆")
                        router.push("/login")
                    })
                }else{
                    console.log("err submit!")
                    return false;
                }
            });
        };

        return {
            ...toRefs(form),
            rules,
            register,
            formel
        }
    },
})
</script>
  
<style lang="scss">
.register-container {
    width: 100%;
    height: 100%;
    padding: 1px;
    background-image: url(../assets/登录背景.jpg);
    background-size: cover;
    .ruleForm {
            text-align: center;
            width: 500px;
            
            margin :200px auto;
            padding: 40px;
            // background-color: aliceblue;
            border-radius: 20px;
            
        }
}
</style>