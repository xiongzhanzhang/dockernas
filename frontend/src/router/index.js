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
          component: () => import('../views/setting.vue'),
          children:[
            {path:'',redirect:'/index/setting/host'},
            {
              path: 'host',
              name: 'host',
              component: () => import('../components/setting/host.vue')
            },
            {
              path: 'storage',
              name: 'storage',
              component: () => import('../components/setting/storage.vue')
            },
            {
              path: 'network',
              name: 'network',
              component: () => import('../components/setting/network.vue')
            },
            {
              path: 'subscribe',
              name: 'subscribe',
              component: () => import('../components/setting/subscribe.vue')
            },
            {
              path: 'image',
              name: 'image',
              component: () => import('../components/setting/image.vue')
            }
          ]
        },
        {
          path: 'instances/:name',
          name: 'instanceInfo',
          component: () => import('../views/instanceInfo.vue'),
          children:[
            {
              path: 'basicInfo',
              name: 'basicInfo',
              props: true,
              component: () => import('../components/instance/instanceBasicInfo.vue')
            },
            {
              path: 'log',
              name: 'log',
              props: true,
              component: () => import('../components/instance/instanceLog.vue')
            },
            {
              path: 'event',
              name: 'event',
              props: true,
              component: () => import('../components/instance/instanceEvent.vue')
            },
            {
              path: 'monitor',
              name: 'monitor',
              props: true,
              component: () => import('../components/instance/instanceMonitor.vue')
            },
            {
              path: 'terminal',
              name: 'terminal',
              props: true,
              component: () => import('../components/instance/instanceTerminal.vue')
            }
          ]
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login.vue')
    },
    {
      path: '/basepath',
      name: 'basePath',
      component: () => import('../views/basePath.vue')
    }
  ]
})

export default router
