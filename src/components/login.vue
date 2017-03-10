<template>
<transition name="fade">
  <div class="login">
    <!--logo开始-->
    <div class="logo">
    <img src="/static/images/login/andon-logo.png" />
    <div class="logo-text">SAAS OA管理系统</div>
    </div>
    <!--logo结束-->
    <!--input区域开始-->
     <group class="andon-input">
      <x-input title="用户" placeholder="" :min="2" :max="10"></x-input>
    </group>
     <group class="andon-input">
      <x-input title="密码" placeholder="" :min="2" :max="10"></x-input>
    </group>
    <!--input区域结束-->
    <div class="andon-forget"><router-link to="/first">忘记密码？</router-link></div>
    <br>
    <!--登陆按钮区开始-->
    <x-button class="andon-btn andon-btn-login" style="margin-top:10px" type="primary" action-type="button" @click.native="goFitst">登陆</x-button>
    <x-button class="andon-btn" style="margin-bottom:10px" @click.native="testAjax">注册</x-button>
    <!--登陆按钮区结束-->
    <!--第三方账号链接登陆开始-->
    <divider>第三方账号链接登陆</divider>
    <div class="additional">
      <img src="/static/images/login/login_01.png" />
      <img src="/static/images/login/login_02.png" />
      <img src="/static/images/login/login_03.png" />
    </div>
    <!--第三方账号链接登陆结束-->
  </div>
  </transition>
</template>

<script>
import { XInput, Group, XButton, Cell, Divider } from 'vux'
export default {
  components: {
    XInput,
    XButton,
    Group,
    Cell,
    Divider
  },
  data () {
    return {
      b: {}
    }
  },
  // 方法声明
  methods: {
    goFitst: function () {
      this.$router.push({ path: '/first' })
    },
    testAjax: function () {
      this.$http.post('http://andontest.lxtxvip.com/api/token/GetTestForVerification',
        {
          // UserName: '11',
          // Type: '7'
        },
        {
          emulateJSON: true
        }).then(function (response) {
          alert('成功')
          alert(response.body.Info)
          alert(response.body.StatusCode)
        }).catch(function (response) {
          alert('失败')
          alert(response.Data)
        })
    }
  }
}
</script>

<style>
body{
  height:100%;
  background: url("/static/images/login/login_background.jpg") no-repeat ;
  background-size: 300%;
}
.logo{
  width: 60%;
  margin: 80px auto;
  
}
.logo .logo-text{
  text-align: center;
  color: #666;
}
.logo img{
  width: 100%;
}
.additional{
  width: 120px;
  margin: 0 auto;
  margin-top:5px;
}
.additional img{
  width:30px;
  margin:0 2px;
}
.andon-btn{
  margin: 0 auto;
  width: 40% !important;
}
.andon-btn-login{
  background-color:rgb(53, 73, 94) !important;
}
.andon-input{
  margin: 0 auto;
  width: 80%;
}
.andon-forget{
  float: right;
  margin-right: 10%;
}
.andon-forget a{
  font-size: 14px;
  color: #666;
}
.red {
  color: red;
}
.green {
  color: green;
}
</style>
