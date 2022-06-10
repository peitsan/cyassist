
<template>
	<!-- <uni-grid :column="2" style="margin: 35px 0px;" :showBorder="false">
		<block v-for="(item,index) in lists">
			<uni-grid-item>
				<div class="flex_column main_axis_space_around second_axis_center" style="background: #ffffff;border-radius: 5px;padding: 40px;" @click="optionClick(index)" :class="index%2==0?'div_left':'div_right'">
					<image :src="item.imageUrl" mode="aspectFit"
					style="width: 50px;height: 50px;"></image>
					<view class="text_investment">
						{{item.title}}
					</view>
				</div>
			</uni-grid-item>
		</block>
	</uni-grid> -->
    <daily-ck></daily-ck>
</template>
<script>
	// import uni-grid from "@/uni_modules/uni-grid/components/uni-grid/uni-grid"
	import dailyCk from "@/pages/assist/dailyCheck/dailyCheck"
	export default {
		components: {
			// uni-grid,
			dailyCk
		},
		mounted(){

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
