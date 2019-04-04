<template>
  <div>
    <el-row align="middle" type="flex">
      <el-col :span="5" :offset="2">
        <img src="../../assets/logo2.png" height="80px">
      </el-col>
      <el-col :span="12">
        <el-input v-model="search_content" clearable>
          <el-button slot="append" icon="el-icon-search">搜索</el-button>
        </el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="20" :offset="2">
        <el-row>
          <el-col :span="3" class="tag-height tag-back">
            品牌
          </el-col>
          <el-col :span="2" v-for="(o, index) in 10" :key="o" class="tag-height">
            <el-card shadow="hover" class="tag-height">
              品牌
            </el-card>
          </el-col>
          <el-col :span="1" class="tag-height">
            <el-card shadow="hover" class="tag-height">
              更多
            </el-card>
          </el-col>
        </el-row>
        <el-row style="margin-top: 5px">
          <el-col :span="3" class="tag-height tag-back">
            类别
          </el-col>
          <el-col :span="2" v-for="(o, index) in 10" :key="o" class="tag-height">
            <el-card shadow="hover" class="tag-height">
              类别
            </el-card>
          </el-col>
          <el-col :span="1" class="tag-height">
            <el-card shadow="hover" class="tag-height">
              更多
            </el-card>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-radio-group v-model="sort_by" size="small" @change="handle_sort_by">
              <el-radio-button label="1">
                综合
                <i class="el-icon-arrow-down"></i>
              </el-radio-button>
              <el-radio-button label="2">
                人气
                <i class="el-icon-arrow-down"></i>
              </el-radio-button>
              <el-radio-button label="3">
                销量
                <i class="el-icon-arrow-down"></i>
              </el-radio-button>
              <el-radio-button label="4">
                价格
                <i class="el-icon-arrow-down"></i>
              </el-radio-button>
              <el-radio-button label="5">
                价格
                <i class="el-icon-arrow-up"></i>
              </el-radio-button>
            </el-radio-group>
          </el-col>

          <el-col :span="3" :offset="5">
            <el-input
              placeholder="请输入最低价"
              v-model="params_.min_price"
              clearable size="small" style="display:inline-block">
            </el-input>
          </el-col>
          <el-col :span="3">
            <el-input
              placeholder="请输入最高价"
              v-model="params_.max_price"
              clearable size="small" style="display:inline-block">
            </el-input>
          </el-col>
          <el-col :span="1">
            <el-button size="small" @click="load_good(params_)">确定</el-button>
          </el-col>
        </el-row>
        <el-row :gutter="5" v-for="(o,i) in row_num" :key="o" style="margin-top: 10px">
          <el-col :span="4" v-for="(o1,i1) in 6" :key="o1">
            <el-card shadow="hover" style="text-align: center;">
              <img src="../../assets/logo.png" height="100px" @click="to_details(goods[i *6 + i1].Id)">
              <p>{{goods[i *6 + i1]['Name']}}</p>
              <p>￥:{{goods[i *6 + i1].Price}}</p>
            </el-card>
          </el-col>
        </el-row>
        <el-row :gutter="5" style="margin-top: 10px">
          <el-col :span="4" v-for="(o1,i1) in col_num" :key="o1">
            <el-card shadow="hover" style="text-align: center;">
              <img src="../../assets/logo.png" height="100px" @click="to_details(goods[row_num *6 + i1].Id)">
              <p>{{goods[row_num *6 + i1]['Name']}}</p>
              <p>￥:{{goods[row_num *6 + i1].Price}}</p>
            </el-card>
          </el-col>
        </el-row>
        <div style="text-align: center">
          <el-pagination
            background
            layout="prev, pager, next"
            @current-change="handleCurrentChange"
            :total="total_page" prev-text="上一页" next-text="下一页">
          </el-pagination>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import {good_list, good_list_page, good_list_v2} from '../../api/api'

  export default {
    name: 'home',
    data () {
      return {
        sort_by: '1',
        search_content: '',
        min_price: '',
        max_price: '',
        total_page: 200,
        goods: [],
        row_num: 0,
        col_num: 0,
        params_: {order: 'asc', sort: 'Id', page: 0, min_price: '', max_price: ''}
      }
    },
    methods: {
      handle_sort_by (val) {
        switch (val) {
          case '5':
            // {order:"asc",sort:"price"}
            this.params_['order'] = 'asc'
            this.params_['sort'] = 'price'
            this.load_good(this.params_)
            break
          case '4':
            this.params_['order'] = 'desc'
            this.params_['sort'] = 'price'
            this.load_good(this.params_)
            break
          default:
            this.params_['order'] = 'asc'
            this.params_['sort'] = 'Id'
            this.load_good(this.params_)
            break
        }
      },
      to_details (gid) {
        this.$router.push('/detail/' + gid)
      },
      load_good (params) {
        good_list_v2(params).then(
          (response) => {
            this.goods = response.data.goods
            // this.row_num = Math.ceil(this.goods.length / 6) -1
            this.row_num = parseInt(this.goods.length / 6)
            this.col_num = this.goods.length % 6
            this.total_page = response.data.num * 10
          }
        )
      },
      handleCurrentChange (val) {
        // good_list_page(val-1).then(
        this.params_['page'] = val - 1
        this.load_good(this.params_)
        // console.log(`当前页: ${val}`);
      }
    },
    created () {
      this.load_good(this.params_)
    },
  }
</script>

<style scoped>
  .tag-back {
    background-color: grey;
  }

  .tag-height {
    height: 100px;
  }

</style>
