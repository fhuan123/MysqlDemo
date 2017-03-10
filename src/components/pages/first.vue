<template>
<transition name="fade">
  <div class="first">
    <!--页头部分开始-->
    <!--<x-header :right-options="{showMore: true}" :left-options="{showBack: false}" @on-click-more="showMenus = true">首页</x-header>-->
    <x-header :left-options="{showBack: false}">主页<router-link slot="right" to="/">退出</router-link></x-header>
    <!--页头部分结束-->

    <!--内容部分开始-->
      <!--banner开始-->
    <swiper loop auto :list="demo06_list" :index="demo06_index"  @on-index-change="demo06_onIndexChange"></swiper>
      <!--banner结束-->
      
      <!--文字滚动区开始-->
    <swiper loop auto height="30px" direction="vertical" :interval=2000 class="text-scroll" :show-dots="false">
      <swiper-item><p>Andon &九安医疗（股票代码002432）</p></swiper-item>
      <swiper-item><p>技术设计心电监护领域</p></swiper-item>
      <swiper-item><p>采用单导联心电监护仪平台</p></swiper-item>
      <swiper-item><p>为血压计增加心率失常检测模块</p></swiper-item>
      <swiper-item><p>IHB测量提高心脑血管基本发现几率</p></swiper-item>
      <swiper-item><p>关心血压同时关爱心脏</p></swiper-item>
    </swiper>
    
      <!--文字滚动区结束-->
      <divider><router-link to="/assembly">新闻信息</router-link></divider>
      <!--有滚动条的竖向滚动开始-->
      <scroller lock-x scrollbar-y height="-352" ref="scroller">
      <!--消息区开始 header="新闻信息"-->
      <panel  :footer="footer" :list="list" :type="type"></panel>
      <!--消息区结束-->
    </scroller>
      <!--有滚动条的竖向滚动结束-->
      
    <!--内容部分结束-->

    <!--页脚部分开始-->
     <tabbar>
      <tabbar-item>
        <!--<i slot="icon" class="fa fa-home" aria-hidden="true"></i>
        <i class="fa fa-home" aria-hidden="true"></i>-->
        <img slot="icon" src="/static/images/icon_nav_button.png">
        <span slot="label">主页</span>
      </tabbar-item>
      <tabbar-item > 
        <!--show-dot-->
        <img slot="icon" src="/static/images/icon_nav_msg.png">
        <span slot="label">已办事项</span>
      </tabbar-item>
      <!--<tabbar-item selected link="/component/demo">
        <img slot="icon" src="../assets/demo/icon_nav_article.png">
        <span slot="label">Explore</span>
      </tabbar-item>-->
      <tabbar-item badge="2">
        <!--<img slot="icon" src="../assets/images/icon_nav_cell.png">-->
        <img slot="icon" src="/static/images/icon_nav_article.png">
        <span slot="label">消息</span>
      </tabbar-item>
    </tabbar>
    <!--页脚部分结束-->
  </div>
   </transition>
</template>

<script>
import { Group, Cell, XHeader, Tabbar, TabbarItem, Swiper, Divider, Panel, SwiperItem, Scroller } from 'vux'

const baseList = [{
  url: 'javascript:',
  img: '/static/images/banner/banner_1.jpg',
  title: '办公地点'
}, {
  url: 'javascript:',
  img: '/static/images/banner/banner_2.jpg',
  title: 'iHealth 智能血压计'
}, {
  url: 'javascript:',
  img: '/static/images/banner/banner_3.jpg',
  title: '移动健康'
}]

const imgList = [
  'http://placeholder.qiniudn.com/800x300/FF3B3B/ffffff',
  'http://placeholder.qiniudn.com/800x300/FFEF7D/ffffff',
  'http://placeholder.qiniudn.com/800x300/8AEEB1/ffffff'
]

const urlList = baseList.map((item, index) => ({
  url: 'https://www.baidu.com/',
  img: item.img,
  title: `${item.title}`
}))

const demoList = imgList.map((one, index) => ({
  url: 'javascript:',
  img: one
}))

const only2List = baseList.slice(0, 2)

export default {
  components: {
    Group,
    Cell,
    XHeader,
    Tabbar,
    TabbarItem,
    Swiper,
    Divider,
    Panel,
    SwiperItem,
    Scroller
  },
  ready () {

  },
  methods: {
    onSwiperItemIndexChange (index) {
      console.log('demo item change', index)
    },
    demo01_onIndexChange (index) {
      this.demo01_index = index
    },
    demo05_onIndexChange (index) {
      this.demo05_index = index
    },
    demo05_onLoad (id) {
      this.demo05_list = id === 1 ? baseList : demoList
    },
    demo06_onIndexChange (index) {
      this.demo06_index = index
    },
    demo07_onIndexChange (index) {
      this.demo07_index = index
    }
  },
  data () {
    return {
      // note: changing this line won't causes changes
      // with hot-reload because the reloaded component
      // preserves its current state and we are modifying
      // its initial state.
      msg: 'Hello World!',
      menus: {
        menu1: 'Take Photo',
        menu2: 'Choose from photos'
      },
      demo01_list: baseList,
      demo02_list: demoList,
      demo03_list: demoList,
      demo04_list: imgList,
      demo05_list: [],
      demo06_list: urlList,
      demo07_list: only2List,
      demo01_index: 0,
      demo02_index: 1,
      demo05_index: 0,
      demo06_index: 0,
      demo07_index: 0,
      swiperItemIndex: 1,
      type: '1',
      list: [{
        src: '/static/images/news/news_1.jpg',
        title: '第三届监事会第二十次会议决议公告',
        desc: '本公司及监事会全体成员保证公告内容真实、准确和完整， 并对公告中的虚假记载、误导性陈述或者重大遗漏承担连带责任。',
        url: '/component/cell'
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }, {
        src: '/static/images/news/news_2.jpg',
        title: 'SAAS OA',
        desc: '为企业提供软件一体化，应用网络化，管理数据化的软件产品和服务。',
        url: {
          path: '/component/radio',
          replace: false
        }
      }
      ],
      footer: {
        title: '查看更多',
        url: 'http://vux.li'
      }
    }
  }
}
</script>

<style >

 .first{
   background-color:#fff; 
 }
 .weui-panel{
  margin-top:0px !important;
}
  .vux-slider{
    margin:2px 0;

  }
.copyright {
  font-size: 12px;
  color: #ccc;
  text-align: center;
}
.text-scroll {
  border: 1px solid #ddd;
  border-left: none;
  border-right: none;
}
.text-scroll p{
  font-size: 12px;
  text-align: center;
  line-height: 30px;
}
.black {
  background-color: #000;
}
.title{
  line-height: 100px;
  text-align: center;
  color: #fff;
}
.animated {
  animation-duration: 1s;
  animation-fill-mode: both;
}
.vux-indicator.custom-bottom {
  bottom: 30px;
}
@-webkit-keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate3d(0, 100%, 0);
  }
  100% {
    opacity: 1;
    transform: none;
  }
}
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate3d(0, 100%, 0);
  }
  100% {
    opacity: 1;
    transform: none;
  }
}
.fadeInUp {
  animation-name: fadeInUp;
}
.swiper-demo-img img {
  width: 100%;
}
</style>
