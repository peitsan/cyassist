<template>
	<view class="page">
		<view class="page-main" :style="'height:'+windowHeight+'px'">
			<view class="top-main">
					<view class="top-main-title">
						<text>重邮小助手</text>
					</view>
			</view>
			<view class="mid-main">
			</view>
		</view>
		<tarbar></tarbar>
	</view>
</template>

<script>

	import tarbar from '../../components/tabbar-home/tabbar/tabbar.vue'
	export default {
		components: {
			tarbar
		},
		data() {
			return {
				windowWidth: getApp().globalData.windowWidth,
				windowHeight: getApp().globalData.windowHeight,
				screenWidth: getApp().globalData.screenWidth, //屏幕宽度
				screenHeight: getApp().globalData.screenHeight, //屏幕高度
				statusBarHeight: getApp().globalData.statusBarHeight, //状态栏高度
				contentHeight: getApp().globalData.windowHeight - getApp().globalData.menuButtonInfo.top - getApp()
				.globalData.menuButtonInfo.height - 20, //内容高度
				menuButtonInfo: getApp().globalData.menuButtonInfo,
			}
		},
		methods:{
		
			}
		
	}
</script>

<style lang="scss">
	.page{
		.page-main{
			.top-main{
				height: 30px;
				background-color: #007AFF;
				.top-main-title{
					text-align: center;
					color:#FFFFFF;
					font-size:$uni-font-size-title
				}
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
