<template>
<div class="asidebar">
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
              <el-icon><Menu /></el-icon>IssueAPI列表
            </template>
            <el-menu-item-group>
              <template #title>Group 1</template>
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
import { useRoute, useRouter } from "vue-router";

const router = useRouter();

const handleRouter = (name:string) => {
    router.push({ name: name })
};

</script>

<style lang="scss">
    .asidebar{
        margin-top: 5px;
        top: 10px;
        overflow: scroll !important;
    }
</style>