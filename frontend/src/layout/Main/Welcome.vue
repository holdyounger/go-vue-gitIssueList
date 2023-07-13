<template>
    <el-tabs
    v-model="editableTabsValue"
    type="card"
    editable
    class="demo-tabs bg-gradient-to-br"
    @edit="handleTabsEdit"
  >
    <el-tab-pane
      v-for="item in editableTabs"
      :key="item.name"
      :label="item.title"
      :name="item.name"
    >
      {{ item.content }}
    </el-tab-pane>
  </el-tabs>
</template>



<script lang="ts" setup>
import { ref } from 'vue'
import type { TabPaneName } from 'element-plus'

let tabIndex = 2
const editableTabsValue = ref('2')
const editableTabs = ref([
  {
    title: '首页',
    name: '1',
    content: 'Tab 1 content',
  },
  {
    title: 'Tab 2',
    name: '2',
    content: 'Tab 2 content',
  },
])

const handleTabsEdit = (
  targetName: TabPaneName | undefined,
  action: 'remove' | 'add'
) => {
  if (action === 'add') {
    const newTabName = `${++tabIndex}`
    editableTabs.value.push({
      title: 'New Tab',
      name: newTabName,
      content: 'New Tab content',
    })
    editableTabsValue.value = newTabName
  } else if (action === 'remove') {
    const tabs = editableTabs.value
    let activeName = editableTabsValue.value
    if (activeName === targetName) {
      tabs.forEach((tab, index) => {
        if (tab.name === targetName) {
          const nextTab = tabs[index + 1] || tabs[index - 1]
          if (nextTab) {
            activeName = nextTab.name
          }
        }
      })
    }

    editableTabsValue.value = activeName
    editableTabs.value = tabs.filter((tab) => tab.name !== targetName)
  }
}
  </script>


<style lang="scss" scoped>
.demo-tabs > .el-tabs__content {
  padding: 3px;
  color: #6b778c;
  font-size: 3px;
  font-weight: 600;
}

.bg-gradient-to-br {
  background-image: -o-linear-gradient(left, var(--tw-gradient-stops));
  background-image: -webkit-gradient(linear, left top, right top, from(var(--tw-gradient-stops)));
  background-image: linear-gradient(to right, var(--tw-gradient-stops));
}

</style>