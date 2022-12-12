import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {path:'/',redirect:'/index/instances'},
    {
      path: '/index',
      name: 'index',
      component: () => import('../views/index.vue'),
      children:[
        {
          path: 'instances',
          name: 'instances',
          component: () => import('../views/instances.vue')
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
