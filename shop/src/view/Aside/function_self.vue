<template>
<el-row :gutter="20">
  <el-col :span="4" :offset="2">
    <el-menu :default-active="active_index" @select="to_menu">
      <el-submenu index="1">
        <template slot="title">
          <span>订单</span>
        </template>
        <el-menu-item-group>
          <el-menu-item index="1-1">全部订单</el-menu-item>
          <el-menu-item index="1-2">待付款</el-menu-item>
          <el-menu-item index="1-3">待发货</el-menu-item>
          <el-menu-item index="1-4">待收货</el-menu-item>
          <el-menu-item index="1-5">待评价</el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-menu-item index="2">收货地址</el-menu-item>
      <el-submenu index="3">
        <template slot="title">
          个人信息
        </template>
        <el-menu-item index="3-1">修改信息</el-menu-item>
        <el-menu-item index="3-2">修改密码</el-menu-item>
      </el-submenu>
      <el-menu-item index="4">我的关注</el-menu-item>
      <el-menu-item index="5">我的足迹</el-menu-item>
    </el-menu>
  </el-col>
  <el-col :span="16">
    <router-view></router-view>
  </el-col>
</el-row>
</template>

<script>
  export default {
    name: 'function_self',
    data(){
      return{
        active_index:'1-1',
      }
    },
    created () {
        this.active_index = this.$route.params['act_index']
    },
    methods:{
      to_menu(index){
        let index_s = index.split('-')
        if(index_s.length === 1){
          if (index_s[0] === '2') {
            this.$router.push("/info/address")
          }else if (index_s[0] == 4){
            this.$router.push("/info/focus")
          } else if (index_s[0] == 5){
            this.$router.push("/info/history")
          }
        }else{
          if(index_s[0] === "1"){
            this.$router.push("/info/order/"+index_s[1])
          }else if (index_s[0] === "3"){
            this.$router.push("/info/info/"+index_s[1])
          }
        }
      },
    }
  }
</script>

<style scoped>

</style>
