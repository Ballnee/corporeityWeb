<template>
    <div>
        <div>
            <el-form :inline="true" :model="selectData" class="demo-form-inline">
                <el-form-item label="学号">
                    <el-input v-model="selectData.stuId" placeholder="请输入学号"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">查询</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table :data="list"  style="width: 100%">
            <el-table-column prop="StudentId" label="学号" width="180">
            </el-table-column>
            <el-table-column prop="StudentName" label="姓名" width="180">
            </el-table-column>
            <el-table-column prop="Gender" label="性别" width="180">
            </el-table-column>
            <el-table-column prop="Birthday" label="生日" width="180">
            </el-table-column>
            <el-table-column prop="Year" label="年级" width="180">
            </el-table-column>
            <el-table-column prop="ClassId" label="班级" width="180">
            </el-table-column>
        </el-table>
    </div>
</template>

<script lang="ts">
import { reactive, defineComponent, toRefs } from 'vue'
import { getStudent } from '@/request/api';
import { InitStu } from '../type/students'
export default defineComponent({
    setup() {
        const data = reactive(new InitStu)
        getStudent().then(res => {
            
            data.list = res.Data
            for (let index = 0; index < res.Data.length;index++) {
                
                data.list[index].Gender = res.Data[index].Gender > 0 ?"男":"女";
                data.list[index].Birthday = res.Data[index].Birthday.substr(0,10);
                data.list[index].Year = Math.floor((2023 - res.Data[index].StudentId/100000));
            }
        })

        return {
            ...toRefs(data)
        }
    }
})
</script>

<style scoped>

</style>