<template>
	<view class="page">

		<view class="page-main" :style="'height:'+contentHeight+'px'">
			<view class="cenTop">
				<view class="mainTop row center">
					<image src="../../../static/co.png" mode="scaleToFill" style="width: 100%;height: 100%;"></image>
				</view>
				<view class="mainMid column ">
					<view class="inv-h-w">
						<view :class="['inv-h',Inv==0?'inv-h-s':'']" :id="0" @click="InvDATA">
							<text style="z-index: 999;">出行政策</text>
						</view>
						<view :class="['inv-h',Inv==1?'inv-h-e':'']" :id="1" @click="InvDATA">
							<text style="z-index: 999;">疫情地区</text>
						</view>
					</view>
					<view class="mai" v-if="Inv == 0">
						<view class="list-row row around items"
							style="background-color: #f5f5f5;height: 50px;margin: 5px;border-radius: 10px;">
							<text @click="popup" id="start">{{cityIfon.startRegionText}}</text>
							<text>⇋</text>
							<text @click="popup" id="end">{{cityIfon.endRegionText}}</text>
						</view>
						<view class="butt"  style="height: 40px;padding: 0 80px;">
							<button type="primary" size="default" @click="determine">确定</button>
						</view>
						<wyb-popup ref="popup" type="bottom" height="400" width="500" radius="6" :showCloseIcon="true">
							<picker-view v-if="visible" :indicator-style="indicatorStyle" :value="value"
								@change="bindChange" class="picker-view">
								<picker-view-column>
									<view class="item" v-for="(item,i) in array" :key="index">{{item.province}}</view>
								</picker-view-column>
								<picker-view-column v-if="array[index[0]]">
									<view class="item" v-for="(item,index) in array[index[0]].citys" :key="index">
										{{item.city}}</view>
								</picker-view-column>
							</picker-view>
							
						</wyb-popup>



					</view>
					<view class="mai" v-if="Inv == 1">


					</view>
				</view>
				<view class="mainBot row center">
					<text
						style="flex: 1;font-weight: bolder;text-align: center;color: #333;font-size: 11px;padding: 0 10px;border-right: #333 solid 1px;line-height: 11px;"></text>
					<text
						style="flex: 1;font-weight: bolder;text-align: center;color: #333;font-size: 11px;padding: 0 10px;border-right: #333 solid 1px;line-height: 11px;"></text>
					<text
						style="flex: 1;font-weight: bolder;text-align: center;color: #333;font-size: 11px;padding: 0 10px;line-height: 11px;"></text>
				</view>
			</view>


		</view>

	</view>
</template>

<script>
	// import wybPopup from '../../../components/wyb-popup/wyb-popup.vue'
	export default {
		components: {
			// wybPopup
		},
		data() {
			return {
				Inv: 0,
				windowWidth: getApp().globalData.windowWidth,
				windowHeight: getApp().globalData.windowHeight,
				screenWidth: getApp().globalData.screenWidth, //屏幕宽度
				screenHeight: getApp().globalData.screenHeight, //屏幕高度
				statusBarHeight: getApp().globalData.statusBarHeight, //状态栏高度
				contentHeight: getApp().globalData.windowHeight, //内容高度
				menuButtonInfo: getApp().globalData.menuButtonInfo,
				index: [0, 0],
				array: [],
				province: [],
				visible: true,
				regionText:"选择城市",
				indicatorStyle: 'height: 50px;' + 'border-top: #19BE6B solid 2px;' + 'border-bottom: #19BE6B solid 2px;',
				cityIfon:{
					startRegionText:"选择城市",
					endRegionText:"选择城市",
					startCity_id:0,
					endCity_id:0,
				},
				
				type:""
			};
		},
		onShow() {},
		onUnload() {},
		onLoad() {
			var that = this;
			uni.request({
				url: 'https://apis.juhe.cn/springTravel/citys?key=' + 'd2ea05d0209b752531762b693f292f5a',
				method: 'get',
				success: (res) => {
					that.array = res.data.result;
					
				}
			});
		},
		methods: {
			determine(){
				var that = this;
				if(that.cityIfon.startCity_id == 0){
					uni.showToast({
						title:"起点未选择"
					})
				}else if(that.cityIfon.endCity_id == 0){
					uni.showToast({
						title:"目的地未选择"
					})
				}else{
					uni.request({
						url: 'http://apis.juhe.cn/springTravel/query?key=' + 'd2ea05d0209b752531762b693f292f5a',
						data:{
							from:that.cityIfon.startCity_id,
							to:that.cityIfon.endCity_id
						},
						method: 'get',
						success: (res) => {
							console.log(res)
							
						}
					});
				}
				
			},
			bindChange: function(e) {
				var that = this;
				const val = e.detail.value
				// var regionText = this.array[this.index[0]].province + this.array[this.index[0]]
				// 	.citys[this.index[1]].city
				if (this.index[0] !== val[0]) {
					// 如果滑动的是第一列数据，让二三列恢复到0
					this.index = [val[0], 0, 0]
					var regionText = this.array[this.index[0]].citys[this.index[1]].city
					var city_id = this.array[this.index[0]].citys[this.index[1]].city_id
					
					if(that.type == "start"){
						that.cityIfon.startCity_id = city_id
						that.cityIfon.startRegionText = regionText
					}
					if(that.type == "end"){
						that.cityIfon.endCity_id = city_id
						that.cityIfon.endRegionText = regionText
					}
					console.log(city_id)
					
				} else if (this.index[1] !== val[1]) {
					// 滑动的是第二列数据
					this.index = [val[0], val[1], 0]
					var regionText = this.array[this.index[0]].citys[this.index[1]].city
					var city_id = this.array[this.index[0]].citys[this.index[1]].city_id
					console.log(city_id)
					if(that.type == "start"){
						that.cityIfon.startCity_id = city_id
						that.cityIfon.startRegionText = regionText
					}
					if(that.type == "end"){
						that.cityIfon.endCity_id = city_id
						that.cityIfon.endRegionText = regionText
					}
				}

			},
			popup(res) {
				this.$refs.popup.show()
				
				var type = res.currentTarget.id;
				if(type == "start"){
					console.log("start")
					this.type = "start"
				}
				if(type == "end"){
					console.log("end")
					this.type = "end"
				}
				
			},
			hide() {
				this.$refs.popup.hide()
			},
			InvDATA(res) {
				console.log(res.currentTarget.id)
				var data = res.currentTarget.id;
				if (data == 0) {
					this.Inv = 0
				} else {
					// this.Inv = 1
					uni.showModal({
						title: '非常抱歉！',
						content: '零担业务正在筹备当中，给您带来的不便请谅解，我们会尽快开通零担业务。',
						showCancel: false,
						success: function(res) {
							if (res.confirm) {
								console.log('用户点击确定');
							} else if (res.cancel) {
								console.log('用户点击取消');
							}
						}
					});
				}
			},
			toPage(e) {
				console.log(e.currentTarget.id)
				var type = e.currentTarget.id;
				if (type == "风险疫情地区") {
					uni.redirectTo({
						url: "./coronavirus/coronavirus"
					})
				}
			}
		}
	}
</script>

<style lang="scss">
	.page {
		// background: linear-gradient(to bottom, #fce7d8, #fff 40%);

	}

	/* 状态栏样式 */
	.statusBarHeight {}

	.tittle {

		.page-main-top-left {
			box-sizing: border-box;
			font-weight: 900;
			font-size: 17px;
			font-weight: 900;
			// background-color: #f7fbfc;
			// border-radius: 20px;
			// opacity: .7;
		}
	}

	/* 主体框架 */
	.page-main {
		display: grid;
		// grid-template-rows: 90px 1fr 50px;
		grid-template-rows: 1fr;
		grid-template-columns: repeat(1, 1fr);
		grid-row-gap: 10px;
		position: relative;
		box-sizing: border-box;
		margin: 10px;
		// background-color: #fff;
		border-radius: 10px;

		.cenTop {

			overflow: hidden;
			border-radius: 10px;
			display: grid;
			grid-template-rows: 130px 1fr 50px;
			grid-template-columns: repeat(1, 1fr);
			// box-shadow: 0px 5px 8px -1px rgba(000, 000, 000, .1);

			.mainTop {
				padding: 0 0px;
				// background: linear-gradient(to left, #52cfdd, #1ebab3);
				border-radius: 10px;
				overflow: hidden;

				.mainTop-left {
					color: #fff;
					width: calc(100%);
					// padding:0 10px;
				}
			}

			.mainMid {
				background-color: #fff;

				.mai {
					box-sizing: border-box;
					
					width: 100%;
					height: 100%;

					.list-row {}

					.picker-view {
						width: 750rpx;
						height: 600rpx;
						margin-top: 20rpx;
					}

					.item {
						display: flex;
						align-items: center;
						justify-content: center;
						text-align: center;
						font-size: 16px;
					}
				}

				.inv-h-w {
					overflow: hidden;
					background-color: #f8f8f8;
					height: 100upx;
					width: 100%;
					display: flex;
				}

				.inv-h {
					font-size: 30upx;
					flex: 1;
					text-align: center;
					color: #C9C9C9;
					height: 100upx;
					line-height: 100upx;
				}

				.inv-h-s {
					position: relative;
					color: #007AFF;
					font-weight: 900;
					z-index: 99;
				}

				.inv-h-s::before {
					content: '';
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					bottom: 0;
					z-index: -1;
					background-color: #fff;
					transform: scaleY(1) perspective(.9em) rotateX(3deg);
					box-shadow: 0px 5px 10px -5px rgba(000, 000, 000, .2);
					transform-origin: left;
					border-top-left-radius: 10px;
					border-top-right-radius: 10px;
				}

				.inv-h-e {
					position: relative;
					color: #007AFF;
					font-weight: 900;
					z-index: 99;
				}

				.inv-h-e::before {
					content: '';
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					bottom: 0;
					z-index: -1;
					background-color: #fff;
					transform: scaleY(1) perspective(.9em) rotateX(3deg);
					box-shadow: 0px 5px 10px -5px rgba(000, 000, 000, .2);
					transform-origin: right;
					border-top-left-radius: 10px;
					border-top-right-radius: 10px;
				}
			}

			.mainBot {
				background-color: #f7fbfc;
			}
		}

		.cenBottom {
			background-color: transparent;
			// border: #000 solid 1px;
			overflow: hidden;
			// padding: 0px 0px;
			overflow-x: auto;
			display: flex;
			flex-flow: column wrap;

			.columnInfo {
				display: grid;
				grid-template-rows: 30px 1fr 30px;
				grid-template-columns: repeat(1, 1fr);
				overflow: hidden;
				height: 45%;
				margin: 10rpx;
				width: calc(100% / 2.1 - 20rpx);
				border-radius: 15px;
				// background: linear-gradient(to bottom, #52cfdd 20%, #fff 60%, #1ebab3 20%);
				box-shadow: 0px 5px 10px -5px rgba(000, 000, 000, .3);

				.columnInfo-top {
					color: #fff;
					font-size: 12px;
					font-weight: bolder;
					padding: 0 20px;
					background-color: #f7fbfc;
					background: linear-gradient(to left, #52cfdd, #1ebab3);

				}

				.columnInfo-mid {

					.midRows {
						height: 100%;
						display: grid;
						grid-template-rows: 1fr 1fr;
						grid-template-columns: repeat(1, 1fr);

						// border: #000 solid 1px;
						.midRows-info {

							// border: #000 solid 1px;
							.midRows-info-right {
								width: calc(100% - 28px);
							}
						}
					}
				}

				.columnInfo-bot {
					color: #000;
					font-size: 11px;
					font-weight: bolder;
					background-color: #f7fbfc;
					// background: linear-gradient(to left, #52cfdd , #1ebab3); 

				}
			}
		}
	}
</style>
