// store.ts
import HomeModule from './modules/HomeModules'
import TableList from './modules/TableList'
import { InjectionKey } from 'vue'
import { createStore, useStore as baseUseStore, Store } from 'vuex'

// 声明自己的 store state
interface State {
    isDark: boolean,
    HomeModule?: any,
    TableList?: any
}

declare module '@vue/runtime-core' {
 
    // 为 `this.$store` 提供类型声明
    interface ComponentCustomProperties {
      $store: Store<State>
    }
  }

// 定义 injection key
export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state:{
        isDark: false,
    },
    getters:{

    },
    mutations:{

    },
    actions:{

    },
    modules:{
        HomeModule,
        TableList
    }
})

// 定义自己的 `useStore` 组合式函数
export function useStore () {
    return baseUseStore(key)
  }