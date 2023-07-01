import { createRouter, createWebHashHistory } from 'vue-router';


const routes = [
    { 
        path: '/', 
        name: 'Home',
        redirect: '/home',
        component: ()=>import('@/layout/layout.vue'),
        children:[
            {
                path: "/home",
                component: () => import('@/views/Home/Home.vue'),
                meta: { title: "首页" }
            }
        ]
    },
    { 
        path: '/about', 
        name: 'About',
        component: ()=>import('@/views/About/About.vue') 
    },
    { 
        path: '/title', 
        name: 'title',
        component: ()=>import('@/views/Titlebar/Titlebar.vue') 
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
  });

export default router