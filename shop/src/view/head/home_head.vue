<template>
  <div>
    <el-row>
      <el-col :span="5" :offset="3">
        <span>欢迎：{{Username}}</span>
        <el-button type="text" @click="to_home()">首页</el-button>
      </el-col>
      <el-col :span="8" :offset="8">
        <el-dropdown  @command="handleCommand">
        <span>
          我的淘宝 <i class="el-icon-arrow-down"></i>
        </span>
          <el-dropdown-menu>
            <el-dropdown-item command="4">我的关注</el-dropdown-item>
            <el-dropdown-item command="5">我的足迹</el-dropdown-item>
            <el-dropdown-item command="1">个人中心</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <el-button type="text" size="medium" @click="cart()">
          <i class="el-icon-goods"></i>
          购物车
        </el-button>
        <el-button type="text" size="medium" @click="to_login()"> 请登录</el-button>
        <el-button type="text" size="medium" @click="to_login()"> 免费注册</el-button>
        <el-button type="text" size="medium" @click="exit_login()"> 退出登陆</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import cookie from '../../static/js/cookie'
  import store from '../../stores/store'

  export default {
    name: 'home_head',
    computed:{
      Username:()=>{
        return store.state.userInfo.Username
      }
    },
    methods:{
      exit_login(){
        cookie.delCookie("Id")
        cookie.delCookie("token")
        store.commit("set_info")
        // console.log(store.state.userInfo)
        store.state.userInfo={Id:'',token:''}
        this.$router.push("/")
      },
      cart(){
        this.$router.push("/cart")
      },
      to_login(){
        this.$router.push("/")
      },
      to_home(){
        this.$router.push('/home')
      },
      handleCommand(command){
        if(command == 5){
          this.$router.push('/info/history')
        }else if (command == 4){
          this.$router.push("/info/focus")
        }else if(command == 1){
          this.$router.push("/info")
        }
      }
    }
  }
</script>

<style scoped>
  .el-button {
    margin-left: 20px;
  }
</style>
