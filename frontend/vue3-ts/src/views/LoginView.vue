<template>
    <div class="login-box">
    <el-form v-bind:model="ruleForm"  ref="ruleFormRef" :rules="rules" status-icon class="ruleForm">
        <h2 style="display: inline-block;">体质健康管理系统</h2>
        <el-form-item label="账号:" prop="username" class="item" >
            <el-input v-model="ruleForm.username" placeholder="Please input your account" ></el-input>
        </el-form-item>
   
        <el-form-item label="密码:" prop="password" class="item" >
            <el-input  v-model="ruleForm.password" placeholder="Please input your password" show-password ></el-input>
        </el-form-item>
        
        <el-form-item>
            <el-button class="loginBt"  type="primary" @click="submitForm(ruleFormRef)">提交</el-button>
            <el-button  class="loginBt" @click="resetForm">重置</el-button>
        </el-form-item>
    </el-form>
    </div>
</template>

<script lang="ts">

import { defineComponent,reactive,toRefs,ref } from 'vue'
import {LoginData} from "../type/login"
import type { FormInstance } from 'element-plus';
import { useRouter } from 'vue-router';
import { login } from '@/request/api';
export default defineComponent({
    setup () {
        const data=reactive(new LoginData());
        const router = useRouter();
        const rules = {
                username: [
                    {
                        required:true,
                        message:"Please input Activity username",
                        trigger:"blur"
                    },
                    {
                        min:3,
                        max:10,
                        message:"length should be 3 to 10"
                    },
                ],
                password:[
                    {
                        required:true,
                        message:"Please input password",
                        trigger:"blur"
                    },
                    {
                        min:6,
                        max:16,
                        message:"password length should be 6 to 16"
                    },
                ]

            }
        const ruleFormRef = ref<FormInstance>()
        const submitForm = (formEl:FormInstance | undefined)=>{
            if (!formEl) return;
            // 对表单的内容进行验证
            formEl.validate((valid)=>{
                if(valid) {
                    console.log("submit!")
                    login(data.ruleForm).then((res)=>{
                        // console.log(res)
                        localStorage.setItem('token',res.Data)
                        router.push("/")
                    })
                }else{
                    console.log("err submit!")
                    return false;
                }
            });
        };
        const resetForm = ()=>{
            data.ruleForm.username = "";
            data.ruleForm.password = "";
            return;
        }
         
        return {
            ...toRefs(data),submitForm,rules,ruleFormRef,resetForm
        }
    }
    
})
</script>

<style lang="scss">

    .login-box{
        width: 100%;
        height: 100%;
        padding: 1px;
        background: url("../assets/登录.jpg") no-repeat;
        background-size: 100% 100%;
        .ruleForm {
            text-align: center;
            width: 500px;
            
            margin :200px auto;
            padding: 40px;
            // background-color: aliceblue;
            border-radius: 20px;
            
        }
        .item.el-form-item__label {
            color: blue;
        }

        .loginBt{
            width: 40%;
            margin-left: 50px ;
        }
    }
</style>