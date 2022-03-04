<template>
  <div id="app">
    <div id="page-head">
      <Navbar/>
    </div>
    <div id="page-body">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Navbar from './views/layout/Navbar'

export default {
  name: 'App',
  components: {
    Navbar
  }
}
</script>

<style>
html, body {
  height: 100%;
  margin: 0;
  padding: 0;
}

#page-head {
  height: 60px;
}

#page-body {
  height: calc(100% - 60px);
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /*text-align: center;*/
  color: #2c3e50;
  height: 100%;
}
</style>
