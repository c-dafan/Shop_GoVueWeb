<template>
<div>
  <div v-if="is_info">
    <el-form label-position="left" label-width="100px" :model="model_user">
      <el-form-item label="用户名">
        <el-input v-model="model_user.username"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button>修改</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div v-else>
    <el-form label-position="top"  label-width="80px" :model="password_change">
      <el-form-item label="旧密码">
        <el-input placeholder="请输入旧密码" v-model="password_change.old_pass"></el-input>
      </el-form-item>
      <el-form-item label="新密码">
        <el-input placeholder="请输入新密码" v-model="password_change.new_pass"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="pass_change">修改</el-button>
      </el-form-item>
    </el-form>
  </div>
</div>
</template>

<script>
  export default {
    name: 'info_user',
    data(){
      return{
        is_info:true,
        password_change:{
          old_pass:'',
          new_pass:''
        },
        model_user:{
          username:'',

        }
      }
    },
    methods:{
      change_tab(){
        this.is_info = this.$route.params["inf"] == 1
      },
      pass_change(){
        console.log("change password")
      }
    },
    created () {
      this.change_tab()
    },
    watch:{
      '$route':'change_tab'
    }
  }
</script>

<style scoped>

</style>
