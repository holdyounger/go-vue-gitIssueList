<template>
<div class="asidebar">
  <el-space>
    <el-image 
    style="width: 40px; height: 40px; border-radius: 10px;right: 10px;" 
    :src="avatar" 
    :zoom-rate="1.2"
    :preview-src-list="srcList" 
    :initial-index="4" 
    fit="cover" />
    <el-text class="mx-1" size="large" tag="b">GitIssue Tools</el-text>
  </el-space>
    <el-scrollbar>
        <el-menu :default-openeds="[]">
          <el-menu-item @click="handleRouter('Welcome')">
            <el-icon color="#409EFC" class="no-inherit"><HomeFilled />
            </el-icon>
            <el-text>
              首页
            </el-text>
          </el-menu-item>
          
          <el-sub-menu index="2">
            <template #title>
              <el-icon><Menu /></el-icon>API列表
            </template>

            <el-menu-item-group>
              <template #title>Issues</template>
              <el-menu-item index="2-1">Option 1</el-menu-item>
              <el-menu-item index="2-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="2-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="2-4">
              <template #title>Option 4</template>
              <el-menu-item index="2-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <el-icon><setting /></el-icon>设置
            </template>
            <el-menu-item-group>
              <template #title>基础设置</template>
              <el-menu-item index="3-1">Option 1</el-menu-item>
              <el-menu-item index="3-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
                <template #title>其它</template>
              <el-menu-item index="3-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="3-4">
              <template #title>其他</template>
              <el-menu-item index="3-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
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

</script>

<style lang="scss">

</style>