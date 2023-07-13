<template>
<div class="asidebar">
    <el-scrollbar class="w-auto ">
        <el-menu :default-openeds="[]" :collapse="useStore().state.HomeModule.navBool" >
          <el-menu-item class="flex justify-center items-center">
            <el-space class="">
              <el-avatar :src="avatar" fit="cover" class="bg-white"/>
            </el-space>
            <el-text class="mx-1 m-auto" size="large">GitIssue Tools</el-text>
          </el-menu-item>
          <el-menu-item @click="handleRouter('Welcome')">
            <el-icon color="#409EFC" class="no-inherit"><HomeFilled />
            </el-icon>
            <el-text>
              首页
            </el-text>
          </el-menu-item>
          
          <el-sub-menu index="2">
            <template #title>
              <el-icon><Menu /></el-icon>
              <el-text>API列表</el-text>
            </template>
            <el-menu-item-group>
              <!-- 列出分配给经过身份验证的用户的问题 -->
              <el-menu-item index="2-1" @click="addAuthenticatedPage">Authenticated List</el-menu-item> 
              <!-- 列出分配给经过身份验证的用户的组织问题 -->
              <el-menu-item index="2-2">Authenticated & Organization</el-menu-item>
              <!-- 列出存储库问题 -->
              <el-menu-item index="2-2">Repository</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon><setting /></el-icon>
              <el-text>设置</el-text>
            </template>
              <el-menu-item index="3-1">基础设置</el-menu-item>
              <el-menu-item index="3-2">其它</el-menu-item>
          </el-sub-menu>
          
        <el-menu-item @click="handleRouter('About')">
            <el-icon color="#409EFC" class="no-inherit"><ChatDotRound />
            </el-icon>
            <el-text>
              关于
            </el-text>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>
</div>
</template>

<script lang="ts" setup>
import aitePng from '@/assets/images/aite.png'
import { reactive,ref } from 'vue';
import { useStore } from '@/store';
import { useRoute, useRouter } from "vue-router";

const avatar = aitePng
const srcList = [
  'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
  'https://fuss10.elemecdn.com/1/34/19aa98b1fcb2781c4fba33d850549jpeg.jpeg',
  'https://fuss10.elemecdn.com/0/6f/e35ff375812e6b0020b6b4e8f9583jpeg.jpeg',
  'https://fuss10.elemecdn.com/9/bb/e27858e973f5d7d3904835f46abbdjpeg.jpeg',
  'https://fuss10.elemecdn.com/d/e6/c4d93a3805b3ce3f323f7974e6f78jpeg.jpeg',
  'https://fuss10.elemecdn.com/3/28/bbf893f792f03a54408b3b7a7ebf0jpeg.jpeg',
  'https://fuss10.elemecdn.com/2/11/6535bcfb26e4c79b48ddde44f4b6fjpeg.jpeg',
]

const router = useRouter();

const handleRouter = (name:string) => {
    router.push({ name: name })
};

type IssueList = {
  
}

let APIList = ref<IssueList> ({

})

const addAuthenticatedPage = ()=>{

};

</script>

<style lang="scss">

.asidebar {
  @apply w-auto;
}

</style>