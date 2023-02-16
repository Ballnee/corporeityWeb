<template>
    <div>
        

        <el-form :model="formData" label-width="120px" ref="form" class="ruleForm">
            <el-form-item label="学号" prop="studentId">
                <el-input v-model="formData.studentId"></el-input>
            </el-form-item>
            <el-form-item label="班级编号" prop="classId">
                <el-input v-model="formData.classId"></el-input>
            </el-form-item>
            <el-form-item label="姓名" prop="name">
                <el-input v-model="formData.name"></el-input>
            </el-form-item>
            <el-form-item label="性别" prop="gender">
                <el-radio-group v-model="formData.gender">
                    <el-radio label="male">男</el-radio>
                    <el-radio label="female">女</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="年龄" prop="age">
                <el-input-number v-model="formData.age" :min="0" :max="150"></el-input-number>
            </el-form-item>
            <el-form-item label="生日" prop="birthday">
                <el-input type="date" :="formData.birthday" value-format="yyyy-MM-dd" ></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="handleSubmit">提交</el-button>
            </el-form-item>
        </el-form>
    </div>

</template>
  
<script lang="ts">
import { defineComponent } from 'vue';
import { ElForm, ElFormItem, ElInput, ElRadio, ElRadioGroup, ElInputNumber, ElDatePicker, ElButton } from 'element-plus';

interface FormData {
    studentId: string;
    classId: string;
    name: string;
    gender: string;
    age: number;
    birthday: string;
}

export default defineComponent({
    name: 'FormDemo',
    components: {
        ElForm,
        ElFormItem,
        ElInput,
        ElRadio,
        ElRadioGroup,
        ElInputNumber,
        ElDatePicker,
        ElButton,
    },
    data() {
        return {
            formData: {
                studentId: '',
                classId: '',
                name: '',
                gender: '',
                age: 0,
                birthday: '',
                
            } as FormData,
        };
    },
    methods: {
        handleSubmit() {
            (this.$refs.form as any).validate((valid: boolean) => {
                if (valid) {
                    console.log(this.formData);
                } else {
                    console.log('表单校验失败');
                }
            });
        },
    },
});
</script>
  
<style lang="scss">
.ruleForm {
    // text-align: center;
    width: 500px;

    margin: 200px auto;
    padding: 40px;
    // background-color: aliceblue;
    border-radius: 20px;

}
</style>