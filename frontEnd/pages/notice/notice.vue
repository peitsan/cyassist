<template>
	<view class = "getNotice">
			<getHistoryCheck>
				
			</getHistoryCheck>
	</view>
</template>

<script>
	import getHistoryCheck from "../assist/dailyCheck/historyCheck.vue"
	export default {
	  name: "notice",
	  components:{
		  getHistoryCheck
	  },
	  data() {
	    return {
			}
		},
	  methods:{
		  
	  }
		
	}
</script>

<style lang="scss">
	.page {
		background: #F4F5F6;
	}

	.register {
		.content-wrapper {
			width: 100%;

			.title {
				margin-top: 35rpx;
				width: 100%;
				margin-bottom: 10px;

				h1 {
					border: 0px;
					width: 50%;
					margin: 0 auto;
					text-align: center;
					border-bottom: 1px solid #E3E3E3;
					height: 50px;
					line-height: 50px;
					font-size: 17px;
					overflow: hidden;
					font-weight: 400;
				}
			}

			.register-form {
				margin: 20px 10px 20px 15px;
				background: #FFFFFF;

				.register-form-items {
					padding: 15px 10px;
					border-bottom: 1px solid #F3F4F5;
					position: relative;
					display: -webkit-flex;
					display: flex;

					.register-form-items-title {
						width: 30%;
						line-height: 22px;
						height: 22px;
						flex-shrink: 0;
					}

					.register-input {
						width: 100%
					}

					img {
						width: auto;
						height: auto;
						max-width: 100%;
						max-height: 100%;
					}
				}
			}
		}

		.submit-wrapper {
			padding: 10px;
			padding-top: 10px;
		}
			
		.list-align-form{
			// margin-left:80%;
			text-align: right;
			.list-align-form-detail{
				// margin-right: ;
				padding-right: 10px;
				}
		}
		
	}
</style>