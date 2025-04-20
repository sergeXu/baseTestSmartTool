<template>
  <div class="hello-container">
    <el-button type="primary" @click="fetchData">获取数据</el-button>
    
    <el-alert
      v-if="message"
      :title="message"
      type="success"
      show-icon
      class="alert-message"
    />
    
    <el-card class="response-card">
      <template #header>
        <div class="card-header">
          <span>服务器响应数据</span>
        </div>
      </template>
      <pre>{{ responseData || '暂无数据' }}</pre>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      message: '',
      responseData: null
    }
  },
  methods: {
    async fetchData() {
      try {
         // 正确访问全局axios实例
        const response = await this.$axios.get('api/v1/hello')
        this.responseData = response.data
        this.message = '请求成功'
      } catch (error) {
        this.$message.error('请求失败: ' + error.message)
        this.message = ''
      }
    }
  }
}
</script>

<style scoped>
.hello-container {
  padding: 20px;
}

.alert-message {
  margin: 15px 0;
}

.response-card {
  margin-top: 20px;
}

pre {
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>