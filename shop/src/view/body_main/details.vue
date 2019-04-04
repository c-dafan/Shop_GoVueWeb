<template>
  <div>
    <el-container>
      <el-header height="150px">
        <span style="margin-left: 200px;font-size: larger">
          {{goods.Sid.Name}}
        </span>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="6" :offset="2">
            <img src="../../assets/logo.png" :height="300">
          </el-col>
          <el-col :span="10">
            <h1 style="height:30px;">商品名字: {{goods.Name}}</h1>
            <div style="background-color: darkturquoise; line-height: 3">
              <h1>单 &nbsp;&nbsp; &nbsp; 价:￥： {{goods.Price}} 元</h1>
            </div>
            数&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;量
            <el-input-number style="margin-left: 15px" :min="1" :max="goods.Num" v-model="buy_num"></el-input-number>
            <br>
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;库存: {{goods.Num}}
            <br>
            <el-button-group style="margin-top: 15px">
              <el-button size="medium" type="primary" round @click="buy_button">立即购买</el-button>
              <el-button size="medium" type="primary" round @click="add_cart_button">加入购物车</el-button>
            </el-button-group>
          </el-col>
          <el-col :span="4">
            <el-card>
              <span>店铺名字: {{goods.Sid.Name}}</span>
              <p>心级: {{goods.Sid.Rank}}</p>
              <p>开店时间: {{goods.Sid.BeginTime.substring(0,10)}}</p>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  import {add_cart, buy_now_post, good_detail, post_buy_item} from '../../api/api'

  export default {
    name: 'detail_shop',
    data () {
      return {
        goods: {
          Sid: {
            Name: '',
            Rank: '',
            BeginTime: ''
          },
          Price: '',
          Num: 0,
          Name: ''
        },
        buy_num: 1,
        gid: 0
      }
    },
    created () {
      this.gid = this.$route.params['gid']
      good_detail(this.gid).then((response) => {
        this.goods = response.data
      })
    },
    methods: {
      add_cart_button () {
        add_cart({num: this.buy_num, gid: {id: parseInt(this.gid)}, uid: {id: 11}}).then(
          (response) => {
            if (response.data['Code'] == 200) {
              this.$alert('添加成功', '购物车', {
                confirmButtonText: '确定',
              })
            } else {
              console.log(response.data['Code'])
              console.log(response.data['Msg'])
            }
          }
        )
      },
      buy_button () {
        buy_now_post({Uid: {Id: 11}, Aid: {Id: 1}}).then(
          (response) => {
            if (response.data['Code'] == 200) {
              let oid = response.data['Data']
              post_buy_item({Num:this.buy_num,Oid:{Id:oid},Gid:{Id:parseInt(this.gid)}}).then(
                (response2)=>{
                  if (response2.data['Code'] == 200) {
                    this.$router.push("/info/order/1")
                  }else {
                    //del
                  }
              }
              )
            } else {
              console.log(response.data['Msg'])
            }
          }
        )
      }
    }
  }
</script>

<style scoped>
  .el-header {
    line-height: 150px;
    background-color: aquamarine;
  }
</style>
