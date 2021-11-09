<script>
	export default {
		onLaunch: function() {

			let _this = this;
				uni.getSystemInfo({
					success: function(res) {

						 let menuButtonInfo = uni.getMenuButtonBoundingClientRect()
						_this.$options.globalData.screenWidth = res.screenWidth; //屏幕宽度
						_this.$options.globalData.screenHeight = res.screenHeight; //屏幕高度
						_this.$options.globalData.windowWidth = res.windowWidth; //可使用窗口宽度
						_this.$options.globalData.windowHeight = res.safeArea.height - res.safeArea.top; //可使用窗口高度
						_this.$options.globalData.statusBarHeight = res.statusBarHeight; //可使用窗口高度
						_this.$options.globalData.menuButtonInfo = menuButtonInfo; //可使用窗口高度

					}
				})
				var that = this;
				uni.getStorage({
					key: 'third_Session',
					success(res) {
						console.log("已获取用户凭证")
						uni.getSetting({
							success(res) {
								console.log(res)
								if (res.authSetting['scope.userInfo']) {

								}else{

								}
							}
						})
					},
					fail(res) {
						console.log("没有获取用户凭证")

					}
				})

		},
		onShow: function() {

		},
		onHide: function() {

		},
		globalData: {
			screenWidth: 0,
			screenHeight: 0,
			windowWidth: 0,
			windowHeight: 0,
			statusBarHeight: 0,
			menuButtonInfo:[],
			domain: 'https://loc.qmtruck.com/', //线上地址
		}
	}
</script>

<style>
	/*每个页面公共css */
	page {
		background-color: #F8F8F8;
	}
	.row {
		display: flex;
		flex-direction: row;
	}

	.center {
		justify-content: center;
		align-items: center;
	}
	.items {
		align-items: center;
	}
	.justify {
		justify-content: center;
	}
	.around {
		justify-content: space-around;
	}
	.between {
		justify-content: space-between;
	}
	.juend{
		justify-content: flex-end;
	}
	.start{
		align-items: flex-start;
	}
	.end{
		align-items: flex-end;
	}
	.column {
		display: flex;
		flex-direction: column;
	}
	.line{
		 overflow: hidden;
		 -webkit-line-clamp: 1;
		text-overflow: ellipsis;
		display: -webkit-box;
		 -webkit-box-orient: vertical;
	}
	@font-face {
		font-family: xmt;
		src: url('https://network-freight-1258702148.cos.ap-chengdu.myqcloud.com/fonts/xmt.ttf');
	}
	@font-face {
		font-family: rzzy;
		src: url('https://network-freight-1258702148.cos.ap-chengdu.myqcloud.com/fonts/rzzy.ttf');
	}

</style>
