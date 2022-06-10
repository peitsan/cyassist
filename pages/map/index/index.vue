<template>
	<view class="chooselocation">
        <view class="location">{{ location? JSON.stringify(location):'' }}</view>
        <view class="choose" @click="onChooseLocation">点击打开地图选择位置</view>
	</view>
</template>

<script>
export default {
    data() {
        return {
            location: null
        }
    },
    onShow () {
        uni.$once('location', (data) => {
            this.location = data;
        });
    },
    methods: {
        onChooseLocation () {
            uni.navigateTo({
                url: '/pages/chooselocation/index'
            });
        }
    }
}
</script>

<style scoped>
    .chooselocation {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 50rpx 0;
    }
    .location {
        padding: 20rpx 0;
        font-size: 28rpx;
        line-height: 1.3;
    }
    .choose {
        display: inline;
        padding: 20rpx 30rpx;
        border: 1rpx solid #e5e5e5;
        font-size: 28rpx;
        line-height: 1;
        border-radius: 10rpx;
    }
</style>
