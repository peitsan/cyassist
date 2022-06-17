import Vue from 'vue';

import VueRouter from "vue-router";

Vue.use(VueRouter)
// 登录
import login from "../pages/login/login"
// 主页
import home from "../pages/index/home"
// 通知
import historyCk from "../pages/assist/dailyCheck/historyCheck.vue"
 
// 用户管理
import my from "../pages/my/my"
// 设置
import option from "../pages/option/option"
// 工具箱
import tools from "../pages/tools/tools"

import register from "../pages/register/register"

import repasswd from "../pages/help/repassword/repassword"
// 助手

import assist from "../pages/assist/assist.vue"
import dailyCk from "../pages/assist/dailyCheck/dailyCheck.vue"
import quickQuit from "../pages/assist/quickQuit/quickQuit.vue"

// 路由权限
 const router = new VueRouter ({
    mode:'history',
	scrollBehavior (to, from, savedPosition) {
	    return savedPosition || { x: 0, y: 0 }
	  },
    routes: [
        {
            path: '/',
            redirect:'/login'
        },
        {
            path: '/login',
            component: login,
			meta: {title: '登录'},
        },
		{
		    path: '/home',
		    component: home,
		    meta: {title: '首页'},
				  
		},
		{
			path: '/assist',
		    component: assist,
			    meta: {title: '助手'},
			    children:[{
						path: '/dailyCheck',
						component: dailyCk,
						meta: {title: '健康打卡'},
							},
						{
						path: '/quickQuit',
						component: quickQuit,
						meta: {title: '一键出校'},
								}
						],
		},
		{
		    path: '/register',
		    component: register,
			meta: {title: '注册'},
		},
		{
		    path: '/notice',
		    component: home,
		    meta: {title: '消息'},
		    
		},
		{
		    path: '/my',
		    component: my,
		    meta: {title: '我的'},
		    
		},
		{
		    path: '/option',
		    component: tools,
		    meta: {title: '设置'},
		    
		},
		{
		    path: '/tools',
		    component: tools,
		    meta: {title: '工具'},
		    
		},
		{
		    path: '/help',
		    // component: login,
			meta: {title: '帮助'},
			children:[
				{
				    path: '/repasswd',
				    component: repasswd,
					meta: {title: '重置密码'},
				}
			]
		},
		
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