<template>
  <div>
    <el-row>
      <el-col :span="20" :offset="2">
        <!--<el-row v-for="(o,i) in 2" :key="o">-->
        <!--<el-checkbox-group v-model="cart_click"  @change="handleCheckedCitiesChange"> </el-checkbox-group>-->
        <el-row v-for="(o,i) in cart_list" :key="i">
          <i class="el-icon-menu"></i>
          {{o.Gid.Sid.Name}}
          <el-card>
            <el-row :gutter="10">
              <el-col :span="1">
                <el-checkbox v-model="cart_click" :label="o.Id" @change="handleCheckedCitiesChange">&nbsp;</el-checkbox>
                <!--:label="o.Id" :key="o.Id"-->
              </el-col>
              <el-col :span="3">
                <img src="../../assets/logo.png" height="100px">
              </el-col>
              <el-col :span="5">
                {{o.Gid.Name}}
              </el-col>
              <el-col :span="4">
                尺寸
              </el-col>
              <el-col :span="4">
                价格 ￥: {{o.Gid.Price | numFilter}}
              </el-col>
              <el-col :span="4">
                <el-input-number size="small" v-model="o.Num" :max="o.Gid.Num"
                                 @change="change_cart_num(i)"></el-input-number>
              </el-col>
              <el-col :span="2">
                价格:{{o.Gid.Price * o.Num | numFilter}}
              </el-col>
              <el-col :span="1">
                <el-button type="text" @click="delete_cart_id(o.Id, i)">删除</el-button>
              </el-col>
            </el-row>
          </el-card>
        </el-row>

        <el-row style="margin-top: 15px;background-color: lightgrey; line-height: 2">
          <el-checkbox style="margin-left: 10px" @change="handleCheckAllChange" v-model="checkAll"
                       :indeterminate="isIndeterminate">全选
          </el-checkbox>
          <el-button type="text" @click="delete_cart">删除</el-button>
          <div style="float: right;">
            已选商品{{cart_click.length}}件
            <span style="margin-left: 20px">合计:{{total_price | numFilter}}</span>
            <el-button style="margin-left: 20px" @click="buy_click()">结算</el-button>
          </div>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import {buy_now_post, delete_cart_cid, get_cart, patch_cart_num, post_buy_item} from '../../api/api'

  export default {
    name: 'cart_home',
    data () {
      return {
        cart_list: [],
        cart_click: [],
        cart_all: [],
        isIndeterminate: false,
        checkAll: false,
        total_price: 0
      }
    },
    watch: {
      cart_click: (val) => {
        // console.log(this.a.methods.price_com())
        // this.a.methods.price_com()
      }
    },
    filters: {
      numFilter (value) {
        // console.log(value)
        let realVal = parseFloat(value).toFixed(2)
        return parseFloat(realVal)
      }
    },
    methods: {
      buy_click () {
        buy_now_post({Uid:{Id:11},Aid:{Id:1}}).then(
          (response)=>{
            for(let v in this.cart_click){
              let ind = this.cart_all.indexOf(this.cart_click[v])
              let item = this.cart_list[ind]
              let oid = response.data['Data']
              post_buy_item({Num:item.Num,Oid:{Id:oid},Gid:{Id:parseInt(item.Gid.Id)}}).then(
                (response2)=>{
                  if (response2.data['Code'] === 200) {
                    delete_cart_cid(item.Id).then(
                      (response3) => {
                        this.cart_all.splice(ind, ind+1)
                        this.cart_list.splice(ind, ind+1)
                      }
                    )
                  }else {
                    //del
                  }
                }
              )
            }
            this.cart_click = []
            this.isIndeterminate = false
            this.checkAll = false
            this.price_com()
          }
        )

        console.log({"item":order_item})
      },
      price_com () {
        let price = 0
        for (let vv in this.cart_all) {
          if (this.cart_click.indexOf(this.cart_all[vv]) === -1) {
            // console.log(this.cart_all[vv])
          } else {
            price += this.cart_list[vv].Gid.Price * this.cart_list[vv].Num
          }
        }
        this.total_price = price
      },
      delete_cart () {
        // console.log(this.cart_click)
        for (let vv in this.cart_click) {
          delete_cart_cid(this.cart_click[vv]).then(
            (response) => {
              this.cart_list.splice(vv, vv + 1)
              this.cart_all.splice(vv, vv + 1)
            }
          )
        }
        this.cart_click = []
        this.isIndeterminate = false
        this.checkAll = false
        this.total_price = 0
      },
      handleCheckAllChange (val) {
        this.cart_click = val ? this.cart_all : []
        this.isIndeterminate = false
        this.price_com()
      },
      handleCheckedCitiesChange (val) {
        let checkedCount = this.cart_click.length
        this.checkAll = checkedCount === this.cart_all.length
        this.isIndeterminate = checkedCount > 0 && checkedCount < this.cart_all.length
        this.price_com()
      },
      delete_cart_id (cid, index) {
        delete_cart_cid(cid).then(
          (response) => {
            this.cart_list.splice(index, index + 1)
            // this.cart_click.remove(cid)
            this.cart_all.splice(index, index + 1)
            let index1 = this.cart_click.indexOf(cid)
            this.cart_click.splice(index1, index1 + 1)
            this.price_com()
          }
        )
      },
      change_cart_num (index) {
        patch_cart_num({Id: this.cart_list[index].Id, Num: this.cart_list[index].Num}).then(
          (response) => {
            this.price_com()
          }
        )
      }
    },
    created () {
      get_cart(11).then((response) => {
        if (response.data['Code'] == 200) {
          this.cart_list = response.data['Data']
          for (let i = 0; i < this.cart_list.length; i++) {
            this.cart_all[i] = this.cart_list[i].Id
          }
        } else {
          this.$alert(response.data['Msg'], '提示')
        }
      })
    }
  }
</script>

<style scoped>

</style>
