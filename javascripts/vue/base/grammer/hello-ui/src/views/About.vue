<template>
  <div class="about">
    <h1>This is an about page</h1>
    <p> {{ msg }} </p>
    <p> nativeCount:{{ count }} </p>
    <p>
      If Element is successfully added to this project, you'll see an
      <code v-text="'<el-button>'"></code>
      below
    </p>
    <el-button @click="add()">el-button</el-button>
  </div>

</template>

<script>
import store from '@/store'

export default {
  name: 'About',
  data () {
    return {
      msg: store.state.count,
      count: 0
    }
  },
  methods: {
    add () {
      this.count++
    }
  }

}

</script>
