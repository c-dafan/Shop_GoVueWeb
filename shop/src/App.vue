<template>
  <div id="app">
    <el-container>
      <el-header height="40px"><router-view name = "head"/></el-header>
      <el-main><router-view name = "main"/></el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
<style>
  .el-header{
    background-color: #B3C0D1;
    color: #333;
    line-height: 40px;
  }
  .el-main{
    height: auto;
    background-color: #E9EEF3;
    color: #333;
  }
</style>
