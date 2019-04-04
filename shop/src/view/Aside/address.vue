<template>
  <div>
    新增收货地址
    <el-row>
      <el-col :span="12">
        <el-form label-position="left" label-width="100px" :model="address_user">
          <el-form-item label="地址">
            <el-input v-model="address_user.AddressJd"></el-input>
          </el-form-item>
          <el-form-item label="详细地址">
            <el-input type="textarea" v-model="address_user.AddressDetails"></el-input>
          </el-form-item>
          <el-form-item label="邮编">
            <el-input v-model="address_user.Mail"></el-input>
          </el-form-item>
          <el-form-item label="收货人姓名">
            <el-input v-model="address_user.Name"></el-input>
          </el-form-item>
          <el-form-item label="手机号码">
            <el-input v-model="address_user.Phone"></el-input>
          </el-form-item>
          <el-form-item size="large">
            <el-button type="primary" @click="onSubmit">保存</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
    <el-table :data="tableData" border>
      <el-table-column prop="Name" label="姓名"></el-table-column>
      <el-table-column prop="AddressJd" label="所在地区"></el-table-column>
      <el-table-column prop="AddressDetails" label="详细地址"></el-table-column>
      <el-table-column prop="Mail" label="邮编"></el-table-column>
      <el-table-column prop="Phone" label="手机号码" width="150"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="text" @click="set_default_address(scope.$index)"
                     :disabled="tableData[scope.$index].Uid.Aid.Id == tableData[scope.$index].Id">
            设为默认地址
          </el-button>
          <el-button
            @click.native.prevent="delete_address_id(scope.$index)"
            type="text"
            :disabled="tableData[scope.$index].Uid.Aid.Id == tableData[scope.$index].Id">
            删除地址
          </el-button>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
  import {delete_address, get_address_uid, patch_user_address, post_address} from '../../api/api'

  export default {
    name: 'address_u',
    data () {
      return {
        address_user: {
          AddressJd: '',
          AddressDetails: '',
          Mail: '',
          Name: '',
          Phone: '',
          Uid: {
            Id: 11
          }
        },
        tableData: []
      }
    },
    methods: {
      set_default_address(index){
        let id = this.tableData[index].Id
        let uid = this.tableData[index].Uid.Id
        patch_user_address(uid,id).then(
          (response)=>{
            if (response.data["Code"] == 200){
              this.get_address_table()
                this.$alert("设置成功","默认地址")
            } else {
              console.log(response.data["Msg"])
            }
          }
        )
      },
      onSubmit () {
        post_address(this.address_user).then(
          (response) => {
            if (response.data['Code'] == 200) {
              this.get_address_table()
              this.address_user = {
                AddressJd: '',
                AddressDetails: '',
                Mail: '',
                Name: '',
                Phone: '',
                Uid: {
                  Id: 11
                }
              }
            } else {
              console.log(response.data['Msg'])
            }
          }
        )
      },
      get_address_table () {
        get_address_uid(11).then(
          (response) => {
            this.tableData = response.data['Data']
          }
        )
      },
      delete_address_id(index){
        delete_address(this.tableData[index].Id).then(
          (response) =>{
            if (response.data["Code"] == 200){
              this.tableData.splice(index, index+1)
            }else {
              console.log(response.data["Msg"])
            }
          }
        )
      }
    },
    created () {
      this.get_address_table()
    }
  }
</script>

<style scoped>

</style>
