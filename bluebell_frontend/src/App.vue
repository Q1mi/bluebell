<template>
  <div id="app">
    <div class="page">
      <HeadBar/>
      <router-view/>
    </div>
    
  </div>
</template>


<script>
import HeadBar from '@/components/HeadBar.vue'
export default {
  components:{
    HeadBar,
  }
}
</script>

<style>
body {
  margin: 0;
}
.page {
  width: 100%;
  height: auto;
  background: #dbe0e5;
}
</style>
