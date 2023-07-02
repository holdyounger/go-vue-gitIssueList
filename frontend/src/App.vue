<script lang="ts" setup>
// import HelloWorld from './components/HelloWorld.vue'
import { useDark, useToggle  } from '@vueuse/core'
import Home from "@/layout/layout.vue"
const isDark = useDark()
const toggleDark = useToggle(isDark)

</script>

<template>
  <!-- <HelloWorld/> -->
  <div style="--wails-draggable:drag">
    <Home/> 
  </div>
</template>

<style>

</style>
