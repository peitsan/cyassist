import Vue from 'vue';

import VueRouter from "vue-router";

Vue.use(VueRouter)
// 路由权限
 const router = new VueRouter ({
    mode:'history',
	scrollBehavior (to, from, savedPosition) {
	    return savedPosition || { x: 0, y: 0 }
	  },
    routes: [
        {
            path: '/',
            redirect:'/login',
			component: () => import('@/pages/login/login')
        },
        {
            path: '/login',
            component: () => import('@/pages/login/login'),
			meta: {title: '登录'},
        },
		{
		    path: '/home',
		    component: () => import('@/pages/index/home'),
		    meta: {title: '首页'},
				 
		},
		{
					path: '/dailyCheck',
					component: () => import('@/pages/assist/dailyCheck/dailyCheck'),
					meta: {title: '健康打卡'},
			},
		{
					path: '/quickQuit',
					component: () => import('@/pages/assist/quickQuit/quickQuit'),
					meta: {title: '一键出校'},
				 
		},
		{
		    path: '/register',
			component: () => import('@/pages/register/register'),
			meta: {title: '注册'}
		},
		{
		    path: '/notice',
			component: () => import('@/pages/notice/notice'),
		    meta: {title: '消息'},
		    
		},
		{
		    path: '/my',
		    component: () => import('@/pages/my/my'),
		     meta: {title: '我的'},
		    
		},
		{
		    path: '/option',
		    component: () => import('@/pages/option/option'),
		     meta: {title: '设置'},
		    
		},
		{
		    path: '/tools',
		   component: () => import('@/pages/tools/tools'),
		     meta: {title: '工具'},
		    
		},
		{
		    path: '/help',
		    // component: login,
			meta: {title: '帮助'},
			children:[
				{
				    path: '/repasswd',
				    component: () => import('@/pages/help/repassword/repassword'),
				    meta: {title: '重置密码'},
				}
			]
		}
		
    ]
})

// function addRe(naviData){
//     // 动态数据追加
//     naviData.forEach(v=>{    
//         router.push({
//             path:v.path,
//             // name:v.name,
//             // meta:{title:v-title},
//             component:()=>import('../pages/'+v.component)
//         })
//     })
//     return router
// }
export default router