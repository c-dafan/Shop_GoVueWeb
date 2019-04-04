<template>
  <div>
    <div v-if="login_is" style="background-image: url(https://img.alicdn.com/tfs/TB1apPmNSzqK1RjSZFLXXcn2XXa-2500-600.jpg);
    height:600px;background-repeat: no-repeat;background-position: center">
      <el-card style="float: right; width: 400px; margin-top: 100px; margin-right: 100px;height: 300px">
        密码登陆
        <el-input placeholder="用户id" style="margin-top: 15px" clearable v-model="user.Id">
          <i slot="prefix" class="el-input__icon el-icon-edit"></i>
        </el-input>
        <el-input show-password placeholder="密码" style="margin-top: 15px" clearable v-model="user.Password">
          <i slot="prefix" class="el-input__icon el-icon-edit"></i>
        </el-input>
        <el-button round size="middle" style="width: 100%;margin-top: 15px" @click="login_()">登陆</el-button>
        {{Msg}}
        <el-button-group style="float: right;margin-top: 20px">
          <el-button type="text" @click="login_is=false">免费注册</el-button>
          <el-button type="text" style="margin-left:15px">忘记密码</el-button>
        </el-button-group>
      </el-card>
    </div>
    <el-row v-if="!login_is">
      <el-col :span="20" :offset="2">
        <el-steps :active="step_reg" simple>
          <el-step title="验证手机号" icon="el-icon-phone"></el-step>
          <el-step title="用户信息" icon="el-icon-edit"></el-step>
          <el-step title="注册成功" icon="el-icon-circle-check-outline"></el-step>
        </el-steps>
        <el-row>
          <el-col :span="14" :offset="5">
            <div v-if="step_reg === 1">
              <el-form label-width="100px" label-position="top">
                <el-form-item label="手机号" >
                  <el-input v-model="phone_num"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button @click="step_reg_sub()">返回登陆</el-button>
                  <el-button @click="step_reg_add()">下一步</el-button>
                </el-form-item>
              </el-form>
            </div>
            <div v-if="step_reg === 2">
              用户名
              <el-input label="用户名" v-model="username_"></el-input>
              密码
              <el-input label="密码" clearable show-password v-model="password_"></el-input>
              重复密码
              <el-input clearable show-password></el-input>
              <el-button @click="step_reg_sub()">上一步</el-button>
              <el-button @click="step_reg_add()">下一步</el-button>
            </div>
            <div v-if="step_reg === 3">
              <h1>注册成功！</h1>
              <h1>用户名：{{user.Username}}</h1>
              <h1>账号为：{{user.Id}}</h1>
              <h1>手机号为：{{user.Phone}}</h1>
              <el-button @click="step_reg_add()">登陆</el-button>
            </div>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import {login, regist} from '../../api/api'
  import cookie from '../../static/js/cookie'
  import store from "../../stores/store"
  export default {
    name: 'login_home',
    data () {
      return {
        login_is: true,
        step_reg: 1,
        phone_num:'',
        username_:'',
        password_:'',
        user:{Id:'',Password:'',Phone:''},
        Msg:''
      }
    },
    methods: {
      step_reg_add () {
        this.step_reg += 1
        if (this.step_reg == 3){
          regist({username:this.username_,
            password:this.password_, phone:this.phone_num}).then((response)=>{
            this.user=response.data
          });
        }
        if (this.step_reg > 3) {
          this.step_reg = 1
          this.login_is = true
        }
      },
      login_ () {
        login({id:parseInt(this.user.Id), password:this.user.Password}).then(
          (response) =>{
            let data = response.data
            if ( data['Code'] == 200){
              cookie.setCookie("Id",data['Data'].Id,1)
              cookie.setCookie("username",data['Data'].Name,1)
              cookie.setCookie("token",data['Data'].Token,1)
              store.commit("set_info")
              this.$router.push('/home')
            }else{
              // console.log(data.Msg)
              this.Msg = data.Msg
            }
          }
        )
      },
      step_reg_sub(){
        if (this.step_reg == 1) {
            this.login_is = true
            return
        }
        this.step_reg -= 1
      }
    }
  }
</script>

<style scoped>

</style>
