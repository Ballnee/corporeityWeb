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
        <le-pagination 
        @current-change="currentChange" 
        @size-change = 'sizeChange'
        layput="prev,pager,next"
        :total="selectData.count"/>
    </div>
</template>

<script lang="ts">
import { computed,reactive, defineComponent, toRefs } from 'vue'
import { getStudent } from '@/request/api';
import { InitStu,type ListInitStudent } from '../type/students'
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
        });
        const currentChange=(page:number)=>{
            data.selectData.page = page
        };
        const sizeChange =(pagesize:number)=>{

        }
        // 需要改变表格展示的数据，所以需要修改list 改成计算属性
        const dataList=reactive({
            // 使用计算属性当data发生改变 计算属性也会改变
            computedList:computed(()=>{
                // 切割展示的数组
                // 页码 1 展示数组【0,9】
                // 页码 2 展示数组【10,19】
                // ...slice左闭右开
                return data.list.slice((data.selectData.page-1)*data.selectData.pageSize,(data.selectData.page)*data.selectData.pageSize)
            })
        });
        const onSubmit=()=>{
            let arr:ListInitStudent[] = []
            if(data.selectData.stuId)
        }
        return {
            ...toRefs(data),currentChange,sizeChange
        }
    }
})
</script>

<style scoped>

</style>