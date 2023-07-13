<template>
  <div id="header" @dblclick="handleFullScreen">
    <el-row :gutter="0" class=" ">
      <!-- 
        xs <768
        sm >=768
        md >=992
        lg >=1200
        xl >=1920
       -->
      <el-col :span="12" :xs="16" :sm="16" :md="20" :lg="20" :xl="20" class="" >
        <el-icon @click="changeNavState" :size="35" class="NavState">
          <Fold v-if="!store.state.HomeModule.navBool" />
          <Expand v-else /> </el-icon>
      </el-col>
      <el-col :span="4" :xs="8" :sm="8" :md="4" :lg="4" :xl="4" class="text-right" >
        <el-button @click="toggleDark()" :icon="GetButtonMode()" size="small" circle plain></el-button>
        <!-- <el-container @mouseenter="handleMouseEnter"> -->
          <el-button @click="handleSmallScreen" @mouseenter="handleMouseEnter" type="success" :icon="SemiSelect" size="small" circle />
          <el-button @click="handleFullScreen" @mouseenter="handleMouseEnter" type="warning" :icon="data.isFullScreen ? CirclePlusFilled : RemoveFilled" size="small" circle />
        <!-- </el-container> -->
        <el-button @click="handleQuit" @mouseleave="handleMouseLeave" @contextmenu="handleRightClick" @mouseenter="handleMouseEnter" type="danger" :icon="SwitchButton" color="#c45656" size="small" circle />
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { reactive, computed } from 'vue'
import { useDark, useToggle } from '@vueuse/core'
import { useStore } from '@/store'
import { } from '@element-plus/icons-vue'
import { 
  Quit, 
  WindowFullscreen,
   WindowIsFullscreen, 
   WindowUnfullscreen,
   WindowMinimise
} from '#/runtime'

import {
  Message,
  SwitchButton,
  CirclePlusFilled,
  RemoveFilled,
  SemiSelect,
  Moon, Sunny,
} from '@element-plus/icons-vue'
import aitePng from '@/assets/images/aite.png'
const avatar = aitePng

const store = useStore()

const isDark = useDark()
const toggleDark = useToggle(isDark);

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
  isFullScreen: false,
  resultText: "Please enter your name below 👇",
})

// const Quit = async() => {
//   const path = await SelectFile("");
//   console.log("选择文件的路径", path)
// }

const changeNavState =() => {
  store.commit("SET_NAV_BOOL")
}

const handleQuit =() => Quit();

const handleSmallScreen =() => WindowMinimise();

const handleFullScreen = async () => {
  data.isFullScreen = await WindowIsFullscreen();
  if( data.isFullScreen )
    WindowUnfullscreen();
  else
    WindowFullscreen();
}


function GetButtonMode() {
  return isDark.value ? 'Moon' : 'Sunny';
}

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

.NavState {
  @apply float-left animate-duration-2s animate-delay-500ms hover:bg-purple-500;
}

</style>