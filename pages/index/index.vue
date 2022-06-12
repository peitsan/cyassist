
<template>
	<view class="page">
		<view class="page-main">
			<view class= "top-main">
				<grid :girdData="girdData" :gutter="'1.563vw'" :gutterBottom="'1.563vw'" :columnNum="'4'">
			  <template #content="acceptProps">
			    <view class="card">
					<transition>
					  <keep-alive>
							<router-link :to="{path:acceptProps.item.path}">
									<view class="vicon">
									<i 	
										:class="acceptProps.item.icon">
									
									</i>
								  </view>
								</router-link>
					  </keep-alive>
					</transition>
				 
				  </router-link>
			      <view class="chName">{{acceptProps.item.chName}}</view>
			      <view class="enName">{{acceptProps.item.enName}}</view>
			    </view>
			  </template>
			</grid>
			</view>
			<view class="mid-main">
				
			</view>
		</view>
</view>
</template>
<script>
	import grid from "@/components/tabbar-grid/tabbar-grid"
	import dailyCk from "@/pages/assist/dailyCheck/dailyCheck"
	export default {
		components: {
			grid,
			dailyCk
		},
		mounted(){
			
		},
		data() {
			return {
			girdData:[{
				chName:"健康打卡",
				enName:"Check Daily",
				icon:"el-icon-document-checked",
				path:"/dailyCheck"
			},{
				chName:"一键出校",
				enName:"Quick Quit",
				icon:"el-icon-position",
				path:"/quickQuit"
				},
				{
					chName:"附近租房",
					enName:"Near Rent",
					icon:"el-icon-office-building",
					path:"/home"
				},{
					chName:"代课发布",
					enName:"Lesson Help",
					icon:"el-icon-help",
					path:""
				},
				{
					chName:"网课急救",
					enName:"Online Help",
					icon:"el-icon-magic-stick",
					path:""
				},{
					chName:"闲置发布",
					enName:"Second Hand",
					icon:"el-icon-finished",
					path:""
					},
					{
						chName:"竞赛发布",
						enName:"Match Host",
						icon:"el-icon-notebook-2",
						path:""
					},{
						chName:"心动表白",
						enName:"Express Love",
						icon:"el-icon-message",
						path:""
						},
						
				],
			// 	windowWidth: getApp().globalData.windowWidth,
			// 	windowHeight: getApp().globalData.windowHeight,
			// 	screenWidth: getApp().globalData.screenWidth, //屏幕宽度
			// 	screenHeight: getApp().globalData.screenHeight, //屏幕高度
			// 	statusBarHeight: getApp().globalData.statusBarHeight, //状态栏高度
			// 	contentHeight: getApp().globalData.windowHeight - getApp().globalData.menuButtonInfo.top - getApp()
			// 		.globalData.menuButtonInfo.height - 20, //内容高度
			// 	menuButtonInfo: getApp().globalData.menuButtonInfo,
			
		onShow() {},
		onUnload() {},
		onLoad() {
			
		},
		
	 }
	},
	mounted() {
	
   },
	methods:{
		onPress(e){
		        console.log(e)
		    }
	},
}
</script>

<style lang="scss"> 
	
	.card{
		border-radius: 4px;
		border: solid;
		background-color: #EcfffE;
		border-color: #EFEFEF;
		.vicon {
			font-size:50px;	
		}
		.chName{
			font-size:15px;
		}
		.enName{
			font-size:12px;
		}
	}
	
	.vicon {
		font-size:50px;	
	}
	.chName{
		font-size:15px;
	}
	.enName{
		font-size:12px;
	}
	.page{
		.page-main{
			.top-main{
				height: 220px;
				background-color: #edefff;
			}
			.mid-main{
				
				.one,.two,.three{
					box-sizing: border-box;
					padding: 10px;
					background-color: #fff;
					.titt{
						margin-left: 5px;
					}
					.one-cen{
						
						.cen{
							box-sizing: border-box;
							padding: 10px;
							margin: 5px;
							border: #f8f8f8 solid 1px;
							border-radius: 10px;
							box-shadow: 2px 2px 0px 0px #f8f8f8;
							.tex{
								height: 100%;
								
								font-weight: 500;
							}
						}
					}
				}
				
			}
			
		}
	}
</style>
