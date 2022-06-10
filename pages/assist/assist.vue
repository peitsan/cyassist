<template>
	<view class="register">
		<view class="content-wrapper">
			<view v-if="registerStep==false" class="title">
				<h1>登录验证</h1>
			</view>
			<view v-if="registerStep==true" class="title">
				<h1>修改统一认证码</h1>
			</view>
			<view
					class="register-form"
			        ref="registerFormRef"
			        :model="registerForm"
			        :rules="registerFormRules"
			        label-width="0px"
			      >
			        <!-- 用户名 -->
				<view v-if="registerStep==false">
			        <view prop="username"
							class="register-form-items">
					<view class="register-form-items-title">账号</view>
			          <input
						class = 'register-input'
			            v-model="registerForm.username"
			            prefix-icon="iconfont icon-user"
			            placeholder="请输入注册邮箱"
			          ></input>
			        </view>
			        <!-- 密码 -->
			        <view prop="password"
							class="register-form-items">
					<view class="register-form-items-title">密码</view>
			          <input
			            type="text"
						class = 'register-input'
			            placeholder="请输入密码"
			            v-model="registerForm.password"
			            prefix-icon="iconfont icon-3702mima"
			          ></input>
			        </view>
					
					<!-- 验证码 -->
					<view v-if='registerVerify=false'>
						<view class="register-form-items">
							<view class="register-form-items-title">验证码</view>
							<input type="text" class="register-input" placeholder="请输入邮箱验证码" />
							<view class="captcha-wrapper">
								<img src="../../static/co.png"></img>
							</view>
						</view>
					</view>
				</view>
			        <!-- 按钮区域 -->
					<!-- 	<view >
								<view class='list-align-form'>
							
										<router-link class ='list-align-form-detail' type = 'primary' to='help/repasswd'>忘记</router-link>
										<router-link  class ='list-align-form-detail' type = 'primary' to='register'>注册</router-link>
	
								</view>
						</view> -->
					<view v-if="registerStep==true">
						<view 
								class="register-form-items">
						<view class="register-form-items-title">账号</view>
						  <input
							class = 'register-input'
							v-model="userid"
						    prefix-icon="iconfont icon-user"
						    placeholder="请修改7位统一认证码"
						  ></input>
						</view>
						<!-- 密码 -->
						<view 
								class="register-form-items">
						<view class="register-form-items-title">密码</view>
						  <input
						    type="text"
							class = 'register-input'
						    placeholder="请修改认证密码"
						    v-model="password"
						    prefix-icon="iconfont icon-3702mima"
						  ></input>
						</view>
							
						<view class="register-form-items">
						<view class="register-form-items-title">openID</view>
						  <input
						    type="text"
							class = 'register-input'
						    placeholder="为微信小程序的登录验证码,选填"
						    v-model="openid"
						    prefix-icon="iconfont icon-3702mima"
						  ></input>
						</view>
						</view>
					</view>
			        <div class="submit-wrapper">
			          <button  v-if="registerStep==false" type="primary" class="register-btn" @click="registerNext()">验证登录</button>
					  <button v-if="registerStep==true" type="primary" class="register-btn" @click="registerToken()">修改信息</button>
			        </div>
			      </view>
			</view>
		</view>
</template>

<script>
	import {post} from "../../utils/axios"
	// import Vue from "vue"
	// Vue.prototype.$axios = axios
	export default {
	  name: "register",
	  data() {
	    return {
		  registerVerify:true,
		  registerStep:undefined,
	      tmpForm:{},
	      // 登录表单的数据绑定对象
	      registerForm:{
	        username:"admin@163.com",
	        password:"123456",
	      },
		    userid:undefined,
		    password:undefined,
			openid:null,
	      // 表单的验证规则对象
	      registerFormRules: {
	        // 验证用户名是否合法
	        username: [
	          { required: true, message: "请输入登录名", trigger: "blur" },
	          {
	            min: 5,
	            // max: 10,
	            message: "长度大于4个字符",
	            trigger: "blur"
	          }
	        ],
			// 验证统一认证码是否真确
			userid: [
			  { required: true, message: "请输入统一认证码", trigger: "blur" },
			  {
			    min: 7,
			    max: 7,
			    message: "长度为7个字符",
			    trigger: "blur"
			  }
			],
	        // 验证密码是否正确
	        password: [
	          { required: true, message: "请输入登录密码", trigger: "blur" },
	          {
	            min: 6,
	            max: 15,
	            message: "长度在 6 到 15 个字符",
	            trigger: "blur"
	          }
	        ],
			
	      }
	    };
	  },
	  mounted(){
		  this.registerStep= false
	  },
	  methods: {
		  async	registerToken(){
			  let res = await post('/register',
			  {
				mail:this.registerForm.username,
				password:this.registerForm.password,
				cqupt_username:this.userid,
				cqupt_password:this.password,
				openid:this.openid,
			  },
			  ).then(res=>{
	            if(res.data.code == 200){
	              // 此处使用mock 状态码随机不加以判断 后面记得修改
	              //localStorage储存在本地
	            localStorage.setItem("atoken",res.data.data.token);
	            localStorage.setItem("obj.role",res.data.data.role.id);
	             this.$router.push("/home"); 
	             console.log("register successfully");
	            }else
	            console.log("register fail");
	          })
	          .catch(err=>{
	          console.log(err)
	        })
	        },
			// 获取token写入localStorage和角色权限
			 // 重置登录表单
			registerNext(){
					  this.registerStep = true
					},
			 resetregisterForm(){
			   this.$refs.registerFormRef.resetFields();
			 },
			 // 验证登录表单是否合法 
			 checkregister(){
			   this.$refs.registerFormRef.validate(valid => {
			     if (!valid) return;
			     else
			     this.registerToken();          // 跳转到登录后的首页
			   })
			 }
			}
	  };
</script>

<style lang="scss">
	page {
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