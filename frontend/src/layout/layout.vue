<template>
  <div class="common-layout">
    <el-container>
        <el-aside class="el-aside" :class="{ 'dark-mode': $store.state.isDark }"><AsideBar/></el-aside>
      <el-container>
          <el-header class="el-header" height="45px" fixed><Header/></el-header>
          <el-main class="el-main"><Main/></el-main>
      </el-container>
    </el-container>
    <el-footer class="el-footer"><Footer/></el-footer>
  </div>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from "vue-router";
import Header from "@/layout/Header/Header.vue";
import AsideBar from "@/layout/Aside/Aside.vue";
import Footer from "@/layout/Footer/Footer.vue";
import Main from "@/layout/Main/Main.vue";

const router = useRouter();

const handleAbout = () => {
    router.push({ path: "/about" })
};

</script>

<style lang="scss" scoped>

.el-footer {
    position: absolute;
    top: 96%;
    width: 100%;
    height: 4%;      
    max-height: 40px;
}

.common-layout {
  margin: 10px 0px;
}

.el-main {
  @apply right-0 top-60 bottom-0 overflow-y-hidden;
  }

.el-aside {
  @apply w-auto  bg-gradient-to-br from-blue-500 to-purple-500 min-h-[calc(100%+1rem)];
}
</style>