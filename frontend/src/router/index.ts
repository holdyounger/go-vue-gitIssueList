import { createRouter, createWebHashHistory } from 'vue-router';


const routes = [
    { 
        path: '/', 
        name: 'Welcome',
        component: ()=>import('@/layout/Main/Welcome.vue'),
        meta: { title: "GitIssue", name:'Welcome' }
    },
    { 
        path: '/Home', 
        name: 'Home',
        children:[
            {
                path: "/home",
                component: () => import('@/views/Home/Home.vue'),
                meta: { title: "首页", name:'Home' }
            }
        ]
    },
    { 
        path: '/about', 
        name: 'About',
        component: ()=>import('@/views/About/About.vue'),
        meta: { title: "关于", name:'About' }
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
  });

export default router