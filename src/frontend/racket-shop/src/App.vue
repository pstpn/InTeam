<template>
  <AppMenuBar />
  <RouterView />
</template>

<script>
import AppMenuBar from './components/AppMenuBar.vue'

export default {
  components: {
    AppMenuBar, 
  }
}
</script>