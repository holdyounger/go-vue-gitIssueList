<script lang="ts" setup>
import { useDark, useToggle  } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)

</script>

<template>
  <button @click="toggleDark()">
    <i inline-block align-middle i="dark:carbon-moon carbon-sun" />

    <span class="ml-2" >{{  isDark ? 'Dark' : 'Light' }} </span>
  </button>
</template>

<style lang="scss" scoped>
.ml-2{
		display:inline-block;
	}

</style>