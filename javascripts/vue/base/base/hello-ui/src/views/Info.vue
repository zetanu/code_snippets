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
      store.commit('increase')
    }
  }
}

</script>
<style scoped>

</style>
