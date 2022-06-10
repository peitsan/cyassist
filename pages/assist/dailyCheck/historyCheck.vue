<template>
	<view class = "inquireCheckData">
		<el-button type="success"
					@click ="getHistoryCheck()"
		>	获取历史打卡
		</el-button>
	</view>
</template>
<script>
	//局部注册   百度地图
	import {nullGet} from '../../../utils/axios.js'

	  export default {
		components:{
			
		},
	    name: "mapView",
	    data() {
	      return {
			text:" ",
			popup:false,
	        map: null,
	        local: null,
	        mk: null,
	        latitude: '',
	        longitude: '',
	        keyWords: '',
			location:'',
			ywjcqzbl:'',
			ywjchblj:'',
			xjzdywqzbl:' ',
			twsfzc:' ',
			ywytdzz:' ',
			beizhu:' ',
			jkmresult:' ',
			offset:undefined,
			
	    }
	},
		mouted(){
			// 保留小数
			
		},
	    methods: {
		async getHistoryCheck(){
					  let res = await nullGet('/user/checkin'
					  ).then(res=>{
		      if(res.data.Code == 200){
						console.log("okay")
		       this.$alert('成功获取历史打卡信息!', '获取成功', {
		       			  				     confirmButtonText: '确定',
		       			  				     callback: action => {
		       			  				       this.$message({
		       			  				         type: 'info',
		       			  				         message: `action: ${ action }`
		       			  				       });
		       			  				     }
		       			  				   })  	
							
		      }else{
							if(res.data.Code == 400)
								this.$alert('请求失败!', '获取失败', {
										  				     confirmButtonText: '确定',
										  				     callback: action => {
										  				       this.$message({
										  				         type: 'info',
										  				         message: `action: ${ action }`
										  				       });
										  				     }
										  				   })
							else if(res.data.Code == 401)
								this.$alert('身份验证失败,请重新登录!', '获取失败', {
										  				     confirmButtonText: '确定',
										  				     callback: action => {
										  				       this.$message({
										  				         type: 'info',
										  				         message: `action: ${ action }`
										  				       });
										  				     }
										  				   })
							else if(res.data.Code == 500)
								this.$alert('服务器匆忙,请稍后重试!', '获取失败', {
										  				     confirmButtonText: '确定',
										  				     callback: action => {
										  				       this.$message({
										  				         type: 'info',
										  				         message: `action: ${ action }`
										  				       });
										  				     }
										  				   }) 
							else {
									console.log("Err601")
									}
						}
						 	
		    })
		    .catch(err=>{
		    console.log(err)
		  })
	}
 },
}
</script>


<style scoped>
	.check-form {
			margin: 20px 10px 20px 15px;
			background: #FFFFFF;
			}
	.check-form-items {
				padding: 15px 10px;
				border-bottom: 1px solid #F3F4F5;
				position: relative;
				display: -webkit-flex;
				display: flex;
	}
	.check-form-items-title {
					width: 40%;
					line-height: 22px;
					height: 22px;
					flex-shrink: 0;
				}
	
	.check-input {
					width: 100%
				}
	.check-select {
					width: 55%
				}
	.check-btn {
					width: 100%
				}
	.submit-wrapper {
		padding: 10px;
		padding-top: 10px;
				}
	
	  .form-layer {
	    width: 98%;
	  }
	  #map {
	    margin-top: 30px;
	    width: 100%;
	    height: 300px;
	    border: 1px solid gray;
	    box-sizing: border-box;
	    overflow: hidden;
	  }
	  /deep/ .el-dialog {
		margin-left:2%;
	    min-width:98%;
		align:center
	  }
	  /deep/ .el-dialog__body {
	    padding: 10px;
	  }
	  .dialog-footer
	.location{
		margin:50%,10px,50%,10px
	}
</style>
