<template>
    <button @click="handleAbout">
        <p>关于</p>
    </button>
    <div class="common-layout">
        common-layout
    <el-container>
      <el-header>Header</el-header>
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div>  
</template>

<script setup>
import { useRoute, useRouter } from "vue-router";

const router = useRouter();

const handleAbout = () => {
    router.push({ path: "/about" })
};

</script>