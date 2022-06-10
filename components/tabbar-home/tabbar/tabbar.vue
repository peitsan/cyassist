<template>
	<view>
		<view class="content">
			<Home v-if="currentPage == 0"/>
			<Notice v-if="currentPage == 1"/>
			<Tools v-if="currentPage == 2"/>
			<Option v-if="currentPage == 3"/>
			<User v-if="currentPage == 4" />
		</view>
		<view class="tabbar" :style="{'padding-bottom': paddingBottomHeight + 'rpx'}">
		<!-- <view class="tabbar"> -->
			<view class="tabbar-item" v-for="(item, index) in list" :key="index" @click="tabbarChange(index)">
				<i :class="item.icon_a" v-if="currentPage == index"></i>
				<i :class="item.icon" v-else></i>
				<view class="item-name" :class="{'tabbarActive': currentPage == index}" v-if="item.text">{{item.text}}
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import Home from "../../../pages/index/index.vue";
	import Notice from "../../../pages/notice/notice.vue";
	import Tools from "../../../pages/tools/tools.vue";
	import Option from "../../../pages/option/option.vue";
	import User from "../../../pages/my/my.vue";
	export default {
		components: {
			Home,
			User,
			Tools,
			Option,
			Notice
		},
		data() {
			return {
				currentPage: 0,
				paddingBottomHeight: 0, //苹果X以上手机底部适配高度
				list: [{
					text: '首页',
					icon: "el-icon-house", //未选中图标
					icon_a: "el-icon-s-home", //选中图片
					path: "../../../pages/index/index", //页面路径
				}, {
					text: '消息',
					icon: "el-icon-bell", //未选中图标
					icon_a: "el-icon-message-solid", //选中图片
					path: "../../../pages/notice/notice", //页面路径
				}, {
					text: '工具',
					icon: "el-icon-more-outline",
					icon_a: "el-icon-more",
					path: "../../../pages/tools/tools",
				}, {
					text: '设置',
					icon:"el-icon-setting",
					icon_a: "el-icon-s-tools",
					path: "../../../pages/option/option",
				}, {
					text: '我的',
					icon: "el-icon-user",
					icon_a: "el-icon-user-solid",
					path: "../../../pages/my/my",
				}, ]
			};
		},
		created() {
		    let that = this;
		    uni.getSystemInfo({
		        success: function (res) {
		            let model = ['X', 'XR', 'XS', '11', '12', '13', '14', '15'];
		            model.forEach(item => {
		                //适配iphoneX以上的底部，给tabbar一定高度的padding-bottom
		                if(res.model.indexOf(item) != -1 && res.model.indexOf('iPhone') != -1) {
		                    that.paddingBottomHeight = 40;
		                }
		            })
		        }
		    });
		},
		watch: {
 
		},
		methods: {
			tabbarChange(index) {
				this.currentPage = index;
			}
		}
	};
</script>

<style lang="scss" scoped>
	.content {
		width: 100%;
		height: 94%;
	}
 
	.tabbar {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		display: flex;
		justify-content: space-around;
		align-items: center;
		width: 100%;
		height: 6%;
		background-color: #ffffff;
 
		.tabbar-item {
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			width: 56rpx;
			height: 85%;
 
			.item-img {
				width: 54rpx;
				height: 54rpx;
			}
 
			.item-name {
				text-align: center;
				font-size: 22rpx;
				font-weight: 400;
				font-family: pingfangSC;
				color: #D8DCE7;
			}
 
			.tabbarActive {
				background-image: linear-gradient(to right top, #1CFDF1, #B330FF);
				font-size: 22rpx;
				-webkit-background-clip: text;
				-moz-background-clip: text;
				background-clip: text;
				box-decoration-break: clone;
				-webkit-box-decoration-break: clone;
				-moz-box-decoration-break: clone;
				color: transparent;
				position: relative;
				animation: mymove 2s infinite;
			}
 
			@keyframes mymove {
				0% {
					transform: scale(1);
					/*开始为原始大小*/
				}
 
				25% {
					transform: scale(1.2);
					/*放大1.1倍*/
				}
 
				50% {
					transform: scale(1);
				}
 
				75% {
					transform: scale(1.2);
				}
 
			}
		}
	}
</style>
