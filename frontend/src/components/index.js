import GDarkMode from './GDarkMode/GDarkMode.vue'
import Breadcrumb from './Breadcrumb/index.vue'

const components = {
    GBreadcrumb:Breadcrumb,
    GDarkMode
}

//自动注册全局组件，
export default function setupComponent(app) {
    Object.keys(components).forEach(key => {
        app.component(key, components[key])
    })
}
