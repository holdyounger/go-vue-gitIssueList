import { DefineComponent } from 'vue'
import Breadcrumb from '@/components/Breadcrumb/index.vue'

const components: { [key: string]: DefineComponent<{}, {}, any> } = {
    GBreadcrumb:Breadcrumb,
}

//自动注册全局组件，
export default function setupComponent(app:any) {
    Object.keys(components).forEach(key => {
        app.component(key, components[key])
    })
}
