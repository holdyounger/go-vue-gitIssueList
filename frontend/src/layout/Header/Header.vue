<template>
  <div id="header">
    <el-row :gutter="0">
      <el-col :span="4" :xs="1" :sm="1" :lg="1" :xl="1" ></el-col>
      <el-col :span="4" :xs="1" :sm="7" class="text-right">
        <el-row>
          <el-col :span="3" :xs="7" :sm="7" :md="5" :lg="5" :xl="8">
            <el-button @click="toggleDark()" :icon="GetButtonMode()" size="small" circle plain></el-button>
          </el-col>
          <el-col :span="3" :xs="10" :sm="12" :md="5" :lg="10" :xl="8">
            <el-badge :value="12" class="item">
              <el-button type="success" :icon="Message" size="small" circle />
            </el-badge>
          </el-col>
          <el-col :span="3" :xs="1" :sm="1" :md="5" :lg="5" :xl="8">
            <el-button @click="Quit" @mouseleave="handleMouseLeave" @contextmenu="handleRightClick" @mouseenter="handleMouseEnter" type="danger" :icon="SwitchButton" color="#c45656" size="small" circle />
          </el-col>
        </el-row>
        <el-row  @mouseenter="handleMouseEnter">
          <el-col :span="3" :xs="0" :sm="7" :md="5" :lg="5" :xl="8">
          </el-col>
          <el-col :span="3" :xs="0" :sm="12" :md="5" :lg="10" :xl="8">
          </el-col>
          <el-col :span="3" :xs="1" :sm="1" :md="5" :lg="5" :xl="8">
            <el-space @mouseleave="handleMouseLeave" @mouseenter="handleMouseEnter" direction="vertical" v-show="data.isShowMenu" wrap>
              <el-button @mouseenter="handleMouseEnter" type="success" :icon="SemiSelect" size="small" circle />
              <el-button @mouseenter="handleMouseEnter" type="warning" :icon="FullScreen" size="small" circle />
            </el-space>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { reactive, computed } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import { } from '@element-plus/icons-vue'
// import {Close,SelectFile} from '../../../wailsjs/go/service/App'
import {Close,SelectFile} from '#/go/service/App'

import {
  Message,
  SwitchButton,
  FullScreen,
  SemiSelect,
  Moon, Sunny,
} from '@element-plus/icons-vue'
import aitePng from '@/assets/images/aite.png'
const avatar = aitePng

const isDark = useDark()
const toggleDark = useToggle(isDark)

const srcList = [
  'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
  'https://fuss10.elemecdn.com/1/34/19aa98b1fcb2781c4fba33d850549jpeg.jpeg',
  'https://fuss10.elemecdn.com/0/6f/e35ff375812e6b0020b6b4e8f9583jpeg.jpeg',
  'https://fuss10.elemecdn.com/9/bb/e27858e973f5d7d3904835f46abbdjpeg.jpeg',
  'https://fuss10.elemecdn.com/d/e6/c4d93a3805b3ce3f323f7974e6f78jpeg.jpeg',
  'https://fuss10.elemecdn.com/3/28/bbf893f792f03a54408b3b7a7ebf0jpeg.jpeg',
  'https://fuss10.elemecdn.com/2/11/6535bcfb26e4c79b48ddde44f4b6fjpeg.jpeg',
]

const data = reactive({
  isShowMenu: false,
  isHover: false,
  resultText: "Please enter your name below 👇",
})

function Quit() {
  SelectFile("","*.rdb").then((path:string) => {
    console.log(`选择的文件路径为：${path}`);
    data.resultText = path;
  })
}

function GetButtonMode() {
  return isDark.value ? 'Moon' : 'Sunny';
}

// 添加一个计算属性来返回颜色的值
const iconColor = computed(() => {
  return "red"; // 这里可以根据你的需求返回不同的颜色值
})

function handleRightClick(event: any) {
  // 阻止浏览器默认的右击事件
  event.preventDefault();

  // 判断鼠标右击事件
  if (event.button === 2) {
    // 执行右击事件的逻辑
    console.log("Right click event detected");
    data.isShowMenu = true
  }
}

function handleMouseEnter() {
  // 在这里编写处理鼠标进入事件的逻辑
  // 例如显示提示信息、改变样式等
  // console.log("Mouse enter event triggered");
  data.isHover = true
}

function handleMouseLeave() {
  // 在这里编写处理鼠标进入事件的逻辑
  // 例如显示提示信息、改变样式等
  data.isHover = false
  setTimeout(() => {
    if(!data.isHover)
    {
      data.isShowMenu = false
    }
  }, 100)
  // console.log("Mouse enter event triggered");

}
</script>

<style lang="scss" scoped>
#header {
  margin-top: 0px;
  ;
}

.el-row {
  margin-bottom: 5px;
}
</style>