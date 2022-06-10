<template>
	<view class="login">
		<view class="content-wrapper">
			<view class="title">
				<h1>欢迎使用重邮小助手</h1>
			</view>
			<view
					class="login-form"
			        ref="loginFormRef"
			        :model="loginForm"
			        :rules="loginFormRules"
			        label-width="0px"
			      >
			        <!-- 用户名 -->
			        <view prop="username"
							class="login-form-items">
					<view class="login-form-items-title">账号</view>
			          <input
						class = 'login-input'
			            v-model="loginForm.username"
			            prefix-icon="iconfont icon-user"
			            placeholder="请输入注册邮箱"
			          ></input>
			        </view>
			        <!-- 密码 -->
			        <view prop="password"
							class="login-form-items">
					<view class="login-form-items-title">密码</view>
			          <input
			            type="text"
						class = 'login-input'
			            placeholder="请输入密码密码"
			            v-model="loginForm.password"
			            prefix-icon="iconfont icon-3702mima"
			          ></input>
			        </view>
					
					<!-- 验证码 -->
					<view v-if='LoginVerify=false'>
						<view class="login-form-items">
							<view class="login-form-items-title">验证码</view>
							<input type="text" class="login-input" placeholder="请输入验证码" />
							<view class="captcha-wrapper">
								<img src="../../static/co.png"></img>
							</view>
						</view>
					</view>
			        <!-- 按钮区域 -->
						<view >
								<view class='list-align-form'>
							
										<router-link class ='list-align-form-detail' type = 'primary' :to="{path:'help/repasswd'}">忘记</router-link>
										<router-link  class ='list-align-form-detail' type = 'primary' :to="{path:'register'}">注册</router-link>
	
								</view>
						</view>
			        <div class="submit-wrapper">
			          <button type="primary" class="login-btn" @click="loginToken()">登录</button>
			          <!-- <el-button type="info" @click="resetLoginForm()">重置</el-button> -->
			        </div>
			      </view>
			</view>
		</view>
</template>

<script>
	
	import {npost} from "../../utils/axios.js"
	export default {
	  name: "login",
	  data() {
	    return {
		  LoginVerify:true,
	      // 登录表单的数据绑定对象
	      loginForm:{
	        username:"admin@163.com",
	        password:"123456",
	      },
	      // 表单的验证规则对象
	      loginFormRules: {
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
	        // 验证密码是否正确
	        password: [
	          { required: true, message: "请输入登录密码", trigger: "blur" },
	          {
	            min: 6,
	            max: 15,
	            message: "长度在 6 到 15 个字符",
	            trigger: "blur"
	          }
	        ]
	      }
	    };
	  },
	  
	  methods: {
		  async	loginToken(){
			  let res = await npost('/login',
			  {
				mail:this.loginForm.username,
				password:this.loginForm.password
			  },
			  ).then(res=>{
				  if(res.data.code == 200){
				    localStorage.setItem("atoken",res.data.token) 
					this.$router.push("/home"); 
	             console.log("Login in");
	            }else
	            console.log("Login out");
	          })
	          .catch(err=>{
	          console.log(err)
	        })
	        }
		  },
		  
		  // 作废Cookie验证
		  //  	var strCookie = document.cookie;
		  //  	//以;为分隔符将所有的cookie进行分割。将获得的所有cookie切割成数组
		  //  	var arrCookie = strCookie.split("; ");
		  //  	//通过for循环进行遍历arrCookie数组。
		  //  	for(var i = 0; i < arrCookie.length; i++){
		  //  	       //通过=进行分割，将本次循环的cookie分割为名字（等于号前），值（等于号后面）
		  //  	        var arr = arrCookie[i].split("=");
		  //  	        //将本次循环的cookie名字与需要查找的cookie进行比较
		  //  	        if(arr == "userInfo"){
		  //  	            //返回指定cookie的值
		  						// 			console.log(arr[1]) ;
		  //  	        }
		  						// 	}
		  // Cookie.set("userInfo",arr[1]);
	    // 获取token写入localStorage和角色权限
	    // 重置登录表单
	    resetLoginForm(){
	      this.$refs.loginFormRef.resetFields();
	    },
	    // 验证登录表单是否合法 
	    checkLogin(){
	      this.$refs.loginFormRef.validate(valid => {
	        if (!valid) return;
	        else
	        this.loginToken();          // 跳转到登录后的首页
	      });
	    }
	  };

</script>

<style lang="scss">
	page {
		background: #F4F5F6;
	}

	.login {
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

			.login-form {
				margin: 20px 10px 20px 15px;
				background: #FFFFFF;

				.login-form-items {
					padding: 15px 10px;
					border-bottom: 1px solid #F3F4F5;
					position: relative;
					display: -webkit-flex;
					display: flex;

					.login-form-items-title {
						width: 30%;
						line-height: 22px;
						height: 22px;
						flex-shrink: 0;
					}

					.login-input {
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