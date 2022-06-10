<template>
	<view class="page">
		<view class="page-main" :style="'height:'+windowHeight+'px'">
			<view class="top-main">
					<view class="top-main-title">
						<text>微重邮打卡小助手</text>
					</view>
			</view>
			<view class="mid-main">
				
				<view class="one" >
					<view class="one-cen row" style="flex-wrap: wrap;" >
						<view class="cen row items" v-for="(item,idx) in currency" :key="i" :id="item.name" style="width: calc(50% - 10px);" @click="toPage">
							<image :src="item.url" mode="aspectFill" style="width: 30px;height: 30px;margin-right: 10px;"></image>
							<view class="tex column " >
								<text style="font-size: 13px;">{{item.name}}</text>
								<text class="line" style="font-size: 12px;color: #888888;">{{item.notes}}</text>
							</view>
							
						</view>
					</view>
				</view>
			<!-- 	<view class="two" style="margin-top: 10px;">
					<view class="titt">
						<text>司机工具</text>
					</view>
					<view class="one-cen row" style="flex-wrap: wrap;" >
						<view class="cen row items" v-for="(item,idx) in drive" :key="i" style="width: calc(50% - 10px);" :id="item.name" @click="toPage">
							<image :src="item.url" mode="aspectFill" style="width: 30px;height: 30px;margin-right: 10px;"></image>
							<view class="tex column " >
								<text style="font-size: 13px;">{{item.name}}</text>
								<text class="line" style="font-size: 12px;color: #888888;">{{item.notes}}</text>
							</view>
							
						</view>
					</view>
				</view>
				<view class="three" style="margin-top: 10px;">
					<view class="titt">
						<text>货主工具</text>
					</view>
					<view class="one-cen row" style="flex-wrap: wrap;" >
						<view class="cen row items" v-for="(item,idx) in owner" :key="i" style="width: calc(50% - 10px);" :id="item.name">
							<image :src="item.url" mode="aspectFill" style="width: 30px;height: 30px;margin-right: 10px;"></image>
							<view class="tex column " >
								<text style="font-size: 13px;">{{item.name}}</text>
								<text class="line" style="font-size: 12px;color: #888888;">{{item.notes}}</text>
							</view>
							
						</view>
					</view>
				</view>
				 -->
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
