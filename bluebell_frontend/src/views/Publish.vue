<template>
<div class="page">
         <header class="header">
            <span class="logo">reddit</span>
            <div class="search">
                <label class="s-logo"></label>
                <input type="text" class="s-input" placeholder="搜索">
            </div>
            <div class="btns">
                <a class="login-btn">LOG IN</a>
                <a class="login-btn">SIGN UP</a>
                <span class="user"></span>
            </div>
        </header>
        <div class="content">
            <div class="left">
                <div class="post-name">
                    我好想写点什么
                </div>
                <div class="post-type">
                    <input type="text" class="post-type-value" placeholder="选择一个频道">
                    <i class="p-icon"></i>
                </div>
                <div class="post-content">
                    <ul class="cat">
                        <li class="cat-item"><i class="iconfont icon-edit"></i>post</li>
                        <li class="cat-item"><i class="iconfont icon-image"></i>image/video</li>
                    </ul>
                    <div class="post-sub-container">
                        <div class="post-sub-header">
                            <textarea class="post-title" name="" id="" cols="30" rows="10" placeholder="标题"></textarea>
                            <span class="textarea-num">0/300</span>
                        </div>
                        <!---此处放置富文本--->
                        <div class="post-text-con"></div>
                    </div>
                    <div class="post-footer">
                        <!-- <ul class="sign">
                            <li class="sign-item">oc</li>
                            <li class="sign-item">spoiler</li>
                            <li class="sign-item">nsfw</li>
                            <li class="sign-item">flair</li>
                        </ul> -->
                        <div class="btns">
                            <button class="btn">取消</button>
                            <button class="btn">发表</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="right">
                <div class="post-rank">
                    <h5 class="p-r-title"><i class="p-r-icon"></i>发帖规范</h5>
                    <ul class="p-r-content">
                        <li class="p-r-item">
                            1.网络不是法外之地
                        </li>
                        <li class="p-r-item">
                            2.网络不是法外之地
                        </li>
                        <li class="p-r-item">
                            3.网络不是法外之地
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped src="../static/css/publish.css"></style>
<style scoped src="../static/css/iconfont.css"></style>