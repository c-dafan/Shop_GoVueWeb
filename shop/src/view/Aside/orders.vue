<template>
  <div>
    <el-card v-for="(o,i) in orders_" :key="i" style="margin-top: 2px">
      <el-row style="background-color: cornflowerblue;line-height: 2">
        <el-col :span="4" :offset="1">时间:{{o.order.Time.substring(0,10) }}</el-col>
        <el-col :span="8">订单号：{{o.order.Id}}</el-col>
        <el-col :span="6">店名：{{o.items[0].Gid.Sid.Name}}</el-col>
        <el-col :span="3">
          <el-popover
            placement="top-start"
            title="收货地址"
            width="200"
            trigger="hover"> <!--:content="o.order.Aid"-->
            <!--<el-table :data="o.order.Aid">-->
              <!--<el-table-column width="150" property="date" label="日期"></el-table-column>-->
              <!--<el-table-column width="100" property="name" label="姓名"></el-table-column>-->
              <!--<el-table-column width="300" property="address" label="地址"></el-table-column>-->
            <!--</el-table>-->
            <div>
              <p>{{o.order.Aid.AddressJd}}</p>
              <p>{{o.order.Aid.AddressDetails}}</p>
              <p>{{o.order.Aid.Mail}}</p>
              <p>{{o.order.Aid.Phone}}</p>
              <p>{{o.order.Aid.Name}}</p>
            </div>
            <el-button slot="reference" type="text" style="color: aquamarine;">查看收货地址</el-button>
          </el-popover>
        </el-col>
      </el-row>
      <el-row v-for="(item,ii) in o.items" :key="ii">
        <el-col :span="6"><img src="../../assets/logo.png" height="120"></el-col>
        <el-col :span="7">商品全称<br>{{item.Gid.Name}}</el-col>
        <el-col :span="3">单价:<br>￥：{{item.Gid.Price}}</el-col>
        <el-col :span="2">数量:<br>{{item.Num}}</el-col>
        <el-col :span="2">总价:<br>{{item.Gid.Price * item.Num}}</el-col>
        <el-col :span="3" :offset="1">
          <!--<el-button type="text">查看评论</el-button>-->
          <el-popover
            placement="top-start"
            title="评论"
            width="200"
            trigger="click"> <!--:content="o.order.Aid"-->
            <div>
              <div v-if="item.Cid == null">
                <el-form>
                  <el-form-item label="评论内容">
                    <el-input v-model="content_up"></el-input>
                  </el-form-item>
                  <el-form-item label="评分">
                    <el-rate
                      v-model="score_up"
                      :colors="['#99A9BF', '#F7BA2A', '#FF9900']">
                    </el-rate>
                  </el-form-item>
                </el-form>
                <el-button type="primary" size="mini" @click="submit_com(item.Id, item.Gid.Id)">确定</el-button>
              </div>
              <div v-else>
                <p>{{item.Cid.Content}}</p>
                <el-rate
                  v-model="item.Cid.Socre"
                  disabled
                  show-score
                  text-color="#ff9900">
                </el-rate>
                <p>{{item.Cid.Time.substring(0,10)}}</p>
              </div>
            </div>
            <el-button slot="reference" type="text" v-if="item.Cid == null">添加评论</el-button>
            <el-button slot="reference" type="text" v-else>查看评论</el-button>
          </el-popover>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
  import {get_order_uid, post_comment} from '../../api/api'

  export default {
    name: 'orders',
    data(){
      return{
        orders_:[
        ],
        score_up:0,
        content_up:''
      }
    },
    methods:{
      submit_com(otid,gid){
        // console.log(gid)
        post_comment({Content:this.content_up,Socre:this.score_up,Otid:{Id:otid},Gid:{Id:gid}}).then(
          (response)=>{
            if (response.data["Code"] == 200){
              this.content_up=''
              this.score_up=0
              this.load_comment()
            } else {
              console.log(response.data["Msg"])
            }
          }
        )
      },
      load_comment(){
        get_order_uid(11).then(
          (response)=>{
            if (response.data["Code"] == 200){
              this.orders_ = response.data["Data"]
            } else {
              console.log(response.data["Msg"])
            }
          }
        )
      }
    },
    created () {
      this.load_comment()
    }
  }
</script>

<style scoped>

</style>
