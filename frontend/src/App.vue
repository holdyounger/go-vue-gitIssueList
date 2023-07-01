<template>
  <div style="--wails-draggable:drag">
    <Titlebar />
  <hr>
      <router-view/>
  </div>
</template>

<script setup>
import Titlebar from './views/Titlebar/Titlebar.vue';
</script>

<style>
html {
  height: 100%;
  overflow: hidden;
  border-radius: 50px;
  margin: 10px;
}
</style>