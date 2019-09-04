<template>
  <div id="app">
        <appHeader> </appHeader>
        <router-view> </router-view>
        <appFooter> </appFooter>
  </div>
</template>

<script>

import Header from './components/share/Header'
import Footer from './components/share/Footer'
export default {
  components: {
    'appHeader': Header,
    'appFooter': Footer
  }
}
</script>


<style lang="scss">
  @import './assets/scss/style.scss';
</style>
