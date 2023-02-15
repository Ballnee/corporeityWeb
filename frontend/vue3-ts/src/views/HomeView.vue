<script setup lang="ts">
import TheWelcome from '../components/TheWelcome.vue'
</script>

<template>
  <div>
    <el-container>
        <el-header>
          <el-row >
            <el-col :span="4"><img src="../assets/logo.svg" class="logo" alt=""></el-col>
            <el-col :span="16"><h2>管理系统</h2></el-col>
            <el-col :span="4"><span class="logout"><i class="el-icon-user"></i>用户</span></el-col>
          </el-row>
        </el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
      background-color="#545c64"
      text-color="#fff"
      router
      active-text-color="#ffd04b">
      <!--<el-submenu index="1">
        <template slot="title">
          <i class="el-icon-location"></i>
          <span>导航一</span>
        </template>
         <el-menu-item-group>
          <template slot="title">分组一</template>
          <el-menu-item index="1-1">选项1</el-menu-item>
          <el-menu-item index="1-2">选项2</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="分组2">
          <el-menu-item index="1-3">选项3</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="1-4">
          <template slot="title">选项4</template>
          <el-menu-item index="1-4-1">选项1</el-menu-item>
        </el-submenu>
      </el-submenu> -->
      <el-menu-item :index="item.path" v-for="item in list" :key="item.path">
        <span slot="title">{{ item.meta.title }}</span>
      </el-menu-item>
      
    </el-menu>
        </el-aside>
        <el-main>
          <RouterView></RouterView>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router';


export default defineComponent({
 
  data () {
    const router = useRouter()
    console.log(router.getRoutes())
    // 先在子路由的isShow设置为true 把子路由添加到list
    const list = router.getRoutes().filter(v=>v.meta.isShow)
    console.log(list)

    return {
      list,
    }
  }
 
});
</script>

<style lang="scss" >
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
     height: calc(100vh - 80px);;
  }

.el-header {
  height:  80px;
  background-color: #b0acac;
  .logo{
    height: 80px;
  }
  h2,.logout{
    text-align: center;
    line-height: 80px;
    background-size: 80px;
  }
.el-aside{
  height: calc(100vh - 80px);
  .el-menu{
    height: calc(100vh - 80px);
  }
  
} 
}
</style>
