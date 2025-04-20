<template>
  <div class="container">
    <el-button type="primary" @click="fetchData">获取数据</el-button>
    <el-alert
      v-if="message"
      :title="message"
      type="success"
      show-icon
    />
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>服务器响应数据</span>
        </div>
      </template>
      <pre>{{ responseData }}</pre>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const message = ref('')
const responseData = ref(null)

const fetchData = async () => {
  try {
    const response = await this.$axios.get('/hello')
    responseData.value = response.data
    message.value = '请求成功'
  } catch (error) {
    ElMessage.error('请求失败: ' + error.message)
    message.value = ''
  }
}
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}
.box-card {
  margin-top: 20px;
}
</style>