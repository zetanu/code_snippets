<template>
  <div>
    Hello Info Component!
    <button type="button" @click="add()">添加</button>
  </div>
</template>

<script>
import store from '@/store'

export default {
  name: 'Info',
  methods: {
    add () {
      console.log('add event from info')
      // 调试方法
      // console.error('add event from info')
      // alert('alert from info')
      // debugger // 添加断点
      store.commit('increase')
    }
  }
}

</script>
<style scoped>

</style>
