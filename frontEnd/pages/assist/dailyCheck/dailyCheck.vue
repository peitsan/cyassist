<template>
	<div>
		<view>
			<view class="top-main">
					<view class="top-main-title">
						<text>重邮小助手</text>
					</view>
			</view>
		  <div>
			  <el-dialog
			    title="请订阅推送微信服务号"
			    :visible.sync="dialogVisible"
				style="text-align:center;width:95%"
			    :before-close="handleClose">
					<view>
						请使用微信扫码订阅推送！
					</view>
						<view> 
						<img :src="this.QRcodeUrl"
							 style="height:250px;
								width:250px
								"> </img>
					</view>
					
				<span slot="footer" class="dialog-footer">
			      <el-button type="primary" @click="dialogVisible = false">结束</el-button>
			    </span>
			  </el-dialog>
		    <el-dialog
		      @close="clearDialog"
		      :close-on-click-modal="false"
		      :title="text"
		      style="text-align:center;width:95%"
		      :visible.sync="popup"
		      width="10%"
		    >
		      <div class="form-layer">
		        <el-form label-width="100px" size="mini">
		          <el-form-item label="获取定位">
		            <el-button type="primary" @click="fixedPos">重新定位</el-button>
		          </el-form-item>
		          <el-form-item label="当前纬度">
		            <p>{{latitude}}</p>
		          </el-form-item>
		          <el-form-item label="当前经度">
		            <p>{{longitude}}</p>
		          </el-form-item>
		          <el-form-item>
		            <div class="f-a-c">
		              <el-input v-model="keyWords" placeholder="请输入地区" style="width: 230px;margin-right: 6px;"></el-input>
		              <el-button type="primary" @click="setPlace" :disabled="!keyWords">查询</el-button>
		            </div>
		          </el-form-item>
		        </el-form>
		        <div id="map"></div>
		      </div>
		      <div slot="footer" class="dialog-footer">
		        <el-button
		          size="small"
		          type="primary"
		          @click="btnSubmit()"
		          >确 认</el-button
		        >
		        <el-button size="small" @click="popup = false">取 消</el-button>
		      </div>
		    </el-dialog>
			
		  </div>
		  
		<el-button type='warning' 
					class="location"
					style="margin: 2%  0 0 5%"
					@click="openDialog('获取当前定位')">
			获取定位
		</el-button>
			
			<view class="check-form-items">
			<view class="check-form-items-title">当前经度</view>
			  <input
				class ='check-input'
				v-model="latitude"
				disabled
			  ></input>
			</view>
			<view class="check-form-items">
				<view class="check-form-items-title">当前纬度</view>
				  <input
					class ='check-input'
					v-model="longitude"
				  ></input>
				</view>
	
		<view class="check-form-items">
			<view class="check-form-items-title">打卡地址</view>
			  <input
				class ='check-input'
				v-model="location"
			    placeholder="请输入详细打卡位置"
				maxlength="50"
			  ></input>
		</view>
		<view class="check-form-items">
			<view class="check-form-items-title">打卡时间偏移</view>
			 <input
				class ='check-input'
				v-model="offset"
			    placeholder="范围在0-3600"
				minlength="1"
				maxlength="4"
			  ></input>
			
		</view>
		
		<view class="check-form-items">
			<view class="check-form-items-title">风险等级</view>
			  <el-select
				class ='check-select'
				v-model="ywjcqzbl"
			    placeholder="请选择新冠风险等级"
			  >
			  <el-option
			        v-for="list in riskLevel"
			        :key="'riskLevel'+list.sort"
			        :label="list.label"
			        :value="list.label">
			      </el-option>
			  </el-select>
		</view>
		<view class="check-form-items">
			<view class="check-form-items-title">渝康码状态</view>
			  <el-select
				class ='check-select'
				v-model="jkmresult"
				placeholder="请选择渝康码颜色"
			  >
			  <el-option
			        v-for="list in codeLevel"
			        :key="'codeLevel'+list.sort"
			        :label="list.label"
			        :value="list.label">
			      </el-option>
			  </el-select>
		</view>
		<view class="check-form-items">
			<view class="check-form-items-title">14天高风险旅居史</view>
			  <el-radio-group v-model="ywjchblj">
			  	<el-radio label= 1 border>是</el-radio>
			  	<el-radio label= 0 border>否</el-radio>
			  </el-radio-group>
		</view>
	
		<view class="check-form-items">
			<view class="check-form-items-title">14天接触高风险者</view>
			 
			 <el-radio-group v-model="xjzdywqzbl"
							 value>
			  	<el-radio label= 1 border>是</el-radio>
			  	<el-radio label= 0 border>否</el-radio>
			 </el-radio-group>
		</view>
		<view class="check-form-items">
			<view class="check-form-items-title">体温是否正常</view>
			  
			  <el-radio-group v-model="twsfzc">
				<el-radio label= 1 border>是</el-radio>
				<el-radio label= 0 border>否</el-radio>
			  </el-radio-group>
			  
		</view>
		<view class="check-form-items">
		<view class="check-form-items-title">相关感染症状</view>
			<el-radio-group v-model="ywytdzz">
				<el-radio label= 0 border>无</el-radio>
				<el-radio label= 1 border>有</el-radio>
			</el-radio-group>
		</view>	
		<view class="check-form-items">
			<view class="check-form-items-title">备注</view>
			  <input
				class ='check-input'
				v-model="beizhu"
			    placeholder="请输入文本备注(选填)"
				maxlength="50"
			  ></input>
		</view>
			<div class="submit-wrapper">
			  <button type="primary" class="check-btn" @click="checkInputForm()">提交自动打卡</button>
			  <!-- <el-button type="info" @click="resetLoginForm()">重置</el-button> -->
			</div>
	</view>
	</div>
</template>
<script>
	//局部注册   百度地图
	import slider from '@/components/e-slider-data/e-slider-data'
	import {vpost,post,get,gettoken} from '../../../utils/axios.js'
	import BaiduMap from 'vue-baidu-map'
	  export default {
		components:{
			slider,
		},
	    name: "mapView",
	    data() {
	      return {
			num: 0,
			timer: null,
			popup:false,
	        map: null,
	        local: null,
	        mk: null,
			QRcodeUrl:undefined,
			bindLinkUrl:undefined,
			subscribeUrl:undefined,
			wxtipUid:'',
			text:' ',
	        latitude: '',
	        longitude: '',
	        keyWords: '',
			location:'',
			ywjcqzbl:'',
			ywjchblj:'0',
			xjzdywqzbl:'0',
			twsfzc:'1',
			ywytdzz:'0',
			beizhu:" ",
			jkmresult:' ',
			offset:undefined,
			dialogVisible:false,
			riskLevel:[{
				sort:"0",
				label:"低风险",
			},{
				sort:"1",
				label:"中风险",
			},{
				sort:"2",
				label:"高风险",
			}
			],
			codeLevel:[{
				sort:"0",
				label:"红色",
			},{
				sort:"1",
				label:"黄色",
			},{
				sort:"2",
				label:"绿色",
			},{
				sort:"3",
				label:"其他",
			}
			]
	      };
	    },
		mouted(){
			// 保留小数
		},
		destroyed() {
	//离开页面是销毁
	    clearInterval(this.timer);
	    this.timer = null;
	  },
		created() {
	  // 实现轮询
	 this.timer = window.setInterval(() => {
	        setTimeout(this.getProjectList(), 0);
	      }, 20000);
	 },
		methods:{
		async responseQRcode(){
					  let code = localStorage.getItem("vcode")
					  let res = await get('https://wxpusher.zjiecode.com/api/fun/scan-qrcode-uid?code='+code
							).then(res=>{
							  if(res.data.code == 1000){
								  console.log("subscribe")
								  this.wxtipUid = res.data.data;
								  this.stop();
								  this.checkUID();
							  }
							  else if(res.data.code==1001){
									console.log(res.data.msg)
							  }
						})
		},
						  
		  async getQRcode(){
			  let res = await post('http://wxpusher.zjiecode.com/api/fun/create/qrcode',
			  {				
				  "appToken":"AT_sbsBbC45ylh7QPpH20kC4UEKIF8j3PfX",   //必填，appToken,前面有说明，应用的标志
				  "extra":gettoken(),      //必填，二维码携带的参数，最长64位
				  "validTime":1800    //可选，二维码的有效期，默认30分钟，最长30天，单位是秒
			  }).then(res=>{
				  localStorage.removeItem("QRcode")
				  localStorage.removeItem("vcode")
				  localStorage.setItem("QRcode",res.data.data.shortUrl)
				  localStorage.setItem("vcode",res.data.data.code)
				  if(res.data.code == 1000)
				  {                    
					  // let surl = 'https://wxpusher.zjiecode.com/api/qrcode/nrw9aX1mhYS05GHoSkiOqjOQND1Hc1r2reV5MBdoYapjAHTxvqmbTCgB5QBAE5dm.jpg'
					// let url= 'https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx9cb41575cad5c731&redirect_uri=http%3A%2F%2Fwxpusher.zjiecode.com%2Fapi%2Fwxuser%2Fauth%2Fcallback%2Fcode%3Fpath%3D%25252Fwxuser%25252F%252523%25252Ffollow&response_type=code&scope=snsapi_base&state=&connect_redirect=1#wechat_redirect'
				  						this.$alert('请进行推送公众号微信绑定!', '推送绑定', {
				  								  				     confirmButtonText: '确定',
				  								  				     callback: action => {
				  								  				       this.$message({
				  								  				         type: 'info',
				  								  				         message: `action: ${ action }`
				  								  				       })
				  													   this.$nextTick(function(){
				  													   	 this.showQRcode();
				  													   })
				  								  				     }
				  								  				   })
				  }
			    }).catch(err=>{
					   console.log(err)
			  })
			 },
		  async	checkUID(){
		  			  let res = await vpost('/user/checkin/submitUID',
		  			  {
		  				UID:this.wxtipUid
		  			  }
		  			  ).then(res=>{  
		        if(res.data.Code == 200){
				 this.dialogVisible = false
		         this.$alert('后台已校验订阅推送', '订阅成功', {
		         			  				     confirmButtonText: '确定',
		         			  				     callback: action => {
		         			  				       this.$message({
		         			  				         type: 'info',
		         			  				         message: `action: ${ action }`
		         			  				       })
		  												   this.$nextTick(function(){
		  												   		this.getQRcode();
		  												   })	
		         			  				     }
		         			  				   })
		  				 
		        }else{
		  					if(res.data.Code == 400)
		  						this.$alert('数据上传格式有误!', '订阅失败', {
		  								  				     confirmButtonText: '确定',
		  								  				     callback: action => {
		  								  				       this.$message({
		  								  				         type: 'info',
		  								  				         message: `action: ${ action }`
		  								  				       });
		  								  				     }
		  								  				   })
		  					else if(res.data.Code == 401)
		  						this.$alert('身份验证失败,请重新登录!', '订阅失败', {
		  								  				     confirmButtonText: '确定',
		  								  				     callback: action => {
		  								  				       this.$message({
		  								  				         type: 'info',
		  								  				         message: `action: ${ action }`
		  								  				       });
		  								  				     }
		  								  				   })
		  					else if(res.data.Code == 500)
		  						this.$alert('服务器匆忙,请稍后重试!', '订阅失败', {
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
		  },
		  async	checkApply(){
		  			  let res = await vpost('/user/checkin/submitdatas',
		  			  {
		  				longitude:String(this.longitude),
		  				latitude:String(this.latitude),
						xxdz:this.location,
						ywjcqzbl:this.ywjcqzbl,
						ywjchblj:(this.ywjchblj=='1'?true:false),
						xjzdywqzbl:(this.xjzdywqzbl=='1'?true:false),
						twsfzc:(this.twsfzc=='1'?true:false),
						ywytdzz:(this.ywytdzz=='1'?true:false),
						jkmresult:this.jkmresult,
						beizhu:this.beizhu,
						offset:Number(this.offset),
						time:this.getTime(),
		  			  }
		  			  ).then(res=>{  
	       	      // if(res.status == 200){
		        if(res.data.Code == 200){
				// console.log("okay")
		         this.$alert('已进入自动打卡模式,设置推送提醒！', '启动成功', {
		         			  				     confirmButtonText: '确定',
		         			  				     callback: action => {
		         			  				       this.$message({
		         			  				         type: 'info',
		         			  				         message: `action: ${ action }`
		         			  				       })
												   this.$nextTick(function(){
												   		this.getQRcode();
												   })	
		         			  				     }
		         			  				   })
				 
		        }else{
					if(res.data.Code == 400)
						this.$alert('数据上传格式有误!', '启动失败', {
								  				     confirmButtonText: '确定',
								  				     callback: action => {
								  				       this.$message({
								  				         type: 'info',
								  				         message: `action: ${ action }`
								  				       });
								  				     }
								  				   })
					else if(res.data.Code == 401)
						this.$alert('身份验证失败,请重新登录!', '启动失败', {
								  				     confirmButtonText: '确定',
								  				     callback: action => {
								  				       this.$message({
								  				         type: 'info',
								  				         message: `action: ${ action }`
								  				       });
								  				     }
								  				   })
					else if(res.data.Code == 500)
						this.$alert('服务器匆忙,请稍后重试!', '启动失败', {
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
		  },
	      // 打开弹窗，name为弹窗名称
	      async openDialog(name) {
	        this.text = name;
	        this.popup = true;
	        this.initMap();
	      },
		  stop() {
		        clearInterval(this.timer);
		        this.timer = null;
		      },
		    // 请求是否有新消息
		   getProjectList() {
		     console.log("请求" + this.num++ + "次");
			 this.responseQRcode();
			 if(this.num == 12){
			 this.stop() 
			}
		    },
		  // 打开对话框
		  showQRcode(url,surl){
			this.dialogVisible = true
			this.QRcodeUrl = localStorage.getItem("QRcode")
			this.$nextTick(function(){
				this.responseQRcode()
			})
				
		  },
		  // 关闭对话框
		  handleClose(done) {
		          this.$confirm('确认关闭？')
		            .then(_ => {
		              done();
		            })
		            .catch(_ => {});
		 },
		  
		  // 输入数据格式校验
		  getTime(){
			  let time = new Date();
			  time=JSON.stringify(time)
			  console.log(typeof(time));
			  var a= time.substring(0,20);
			  var b= time.substring(24)
			  console.log(a+b)
			  return (a+b).substring(1,21)
		  },
		  checkInputForm(){
			  if(this.offset < 0 || this.offset > 3600){
			  			  this.$alert('时间偏移参数不在范围(0,3600)s内!', '请修改参数', {
			  				     confirmButtonText: '确定',
			  				     callback: action => {
			  				       this.$message({
			  				         type: 'info',
			  				         message: `action: ${ action }`
			  				       });
			  				     }
			  				   })  				
			  }else
				this.checkApply()
			},
	      // 确认
	      btnSubmit() {
	        let key = {
	          latitude: Number(this.latitude),
	          longitude: String(this.longitude).replace(/^(.*\..{4}).*$/,"$1"),
	        }
	        this.popup = false;
	      },
		  
	      initMap() {
	        this.$nextTick(() => {
	          this.map = new BMap.Map("map");
	          let point = new BMap.Point(116.404, 39.915);
	          this.map.centerAndZoom(point, 12);
	          this.map.enableScrollWheelZoom(true); // 开启鼠标滚轮缩放
	          this.map.addControl(new BMap.NavigationControl());
	          this.fixedPos();
	        });
	      },
	      // 点击定位-定位到当前位置
	      fixedPos() {
	        const _this = this;
	        const geolocation = new BMap.Geolocation();
	        this.confirmLoading = true;
	        geolocation.getCurrentPosition(function (r) {
	          if (this.getStatus() == BMAP_STATUS_SUCCESS) {
	            _this.handleMarker(_this, r.point);
	            let myGeo = new BMap.Geocoder();
	            myGeo.getLocation(
	              new BMap.Point(r.point.lng, r.point.lat),
	              function (result) {
	                _this.confirmLoading = false;
	                if (result) {
	                  _this.latitude = String(result.point.lat).replace(/^(.*\..{4}).*$/,"$1");
	                  _this.longitude = String(result.point.lng).replace(/^(.*\..{4}).*$/,"$1");
				    }
	              }
	            );
	          } else {
	            _this.$message.error("failed" + this.getStatus());
	          }
	        });
	      },
		  // // 改变上传参数状态
		  // agreeChange(tmp){
			 //  that=String(tmp);
			 //  this.that=true;
		  // },
	      // 搜索地址
	      setPlace() {
	        this.local = new BMap.LocalSearch(this.map, {
	          onSearchComplete: this.searchPlace,
	        });
	        this.local.search(this.keyWords);
	      },
	      searchPlace() {
	        if (this.local.getResults() != undefined) {
	          this.map.clearOverlays(); //清除地图上所有覆盖物
	          if (this.local.getResults().getPoi(0)) {
	            let point = this.local.getResults().getPoi(0).point; //获取第一个智能搜索的结果
	            this.map.centerAndZoom(point, 18);
	            this.handleMarker(this, point);
	            console.log("经度：" + point.lng + "--" + "纬度" + point.lat);
	            this.latitude =String(point.lat).replace(/^(.*\..{4}).*$/,"$1");
	            this.longitude =String(point.lng).replace(/^(.*\..{4}).*$/,"$1");
	          } else {
	            this.$message.error("未匹配到地点!");
	          }
	        } else {
	          this.$message.error("未找到搜索结果!");
	        }
	      },
	      // 设置标注
	      handleMarker(obj, point) {
	        let that = this;
	        obj.mk = new BMap.Marker(point);
	        obj.map.addOverlay(obj.mk);
	        obj.mk.enableDragging(); // 可拖拽
	        obj.mk.addEventListener("dragend", function (e) {
	          // 监听标注的拖拽，获取拖拽后的经纬度
	          that.latitude = e.point.lat;
	          that.longitude = e.point.lng;
	        });
	        obj.map.panTo(point);
	      },
		} 
	  };
	
</script>


<style scoped>
	.top-main{
		height: 30px;
		background-color: #007AFF;
		
	}
	.top-main-title{
		text-align: center;
		color:#FFFFFF;
		font-size:$uni-font-size-title
	}
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
