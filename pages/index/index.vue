<template>
	<view class="page">
		<view class="page-main" :style="'height:'+windowHeight+'px'">
			<view class="top-main">
				
			</view>
			<view class="mid-main">
				
				<view class="one" >
					<view class="titt">
						<text>通用工具</text>
					</view>
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
				<view class="two" style="margin-top: 10px;">
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
			</view>
		</view>

	</view>
</template>

<script>
	export default {
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
				currency: [{
						url:"../../static/logo.png",
						name: '风险疫情地区',
						notes:"查看各地疫情风险等级"
					}, {
						url:"../../static/logo.png",
						name: '最新高速路况',
						notes:"最新高速路况信息"
					}, {
						url:"../../static/logo.png",
						name: '货车定位',
						notes:"12吨及以上货车实时位置查询"
					}, {
						url:"../../static/logo.png",
						name: '货车轨迹',
						notes:"12吨货车轨迹查询"
					}
				],
				drive: [{
						url:"../../static/logo.png",
						name: '附近企业',
						notes:"查看附近的物流园区"
					}, {
						url:"../../static/logo.png",
						name: '附近停车场',
						notes:"查看附近停车场"
					}, {
						url:"../../static/logo.png",
						name: '附近加油站',
						notes:"查看附近加油站"
					}, {
						url:"../../static/logo.png",
						name: '附近加气站',
						notes:"查看附近加气站"
					}, {
						url:"../../static/logo.png",
						name: '车辆违章',
						notes:"查看自己车辆违章"
					}, {
						url:"../../static/logo.png",
						name: '今日油价',
						notes:"查看今日油价"
					}, {
						url:"../../static/logo.png",
						name: '危险品名录',
						notes:"查看违禁品货物"
					}, {
						url:"../../static/logo.png",
						name: '货车名片',
						notes:"查看附近的物流园区"
					}
				],
				owner: [{
						url:"../../static/logo.png",
						name: '查老赖',
						notes:"查看交易对方是否有失信行为"
					}, {
						url:"../../static/logo.png",
						name: '司机资格证核验',
						notes:"最新高速路况信息"
					}, {
						url:"../../static/logo.png",
						name: '快递查询',
						notes:"查询货物回单快递"
					}, {
						url:"../../static/logo.png",
						name: '车辆营运证核验',
						notes:"查看车辆是否正规取得"
					}, {
						url:"../../static/logo.png",
						name: '查看周边货车',
						notes:"查看附近货车数量"
					}, {
						url:"../../static/logo.png",
						name: '套牌车验证',
						notes:"验证货车是否为套牌"
					}
				],
			};
		},
		onShow() {},
		onUnload() {},
		onLoad() {
			
		},
		methods: {
			toPage(e){
				console.log(e.currentTarget.id)
				var type = e.currentTarget.id;
				if(type == "风险疫情地区"){
					uni.navigateTo({
						url:"./coronavirus/coronavirus"
					})
				}
			}
		}
	}
</script>

<style lang="scss">
	.page{
		.page-main{
			.top-main{
				height: 130px;
				background-color: #007AFF;
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
