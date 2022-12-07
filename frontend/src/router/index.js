import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {path:'/',redirect:'/index/apps'},
    {
      path: '/index',
      name: 'index',
      component: () => import('../views/index.vue'),
      children:[
        {
          path: 'apps',
          name: 'apps',
          component: () => import('../views/apps.vue')
        },
        {
          path: 'store',
          name: 'store',
          component: () => import('../views/store.vue')
        },
        {
          path: 'setting',
          name: 'setting',
          component: () => import('../views/setting.vue')
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login.vue')
    }
  ]
})

export default router
