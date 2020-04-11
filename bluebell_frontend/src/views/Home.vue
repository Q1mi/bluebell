<template>
  <div class="content">
    <div class="left">
      <!-- <h4 class="c-l-title">热门帖子</h4> -->
      <div class="c-l-header">
        <span class="hot btn-iconfont">
          <i class="iconfont icon-hot"></i>
          Hot
        </span>
        <div class="new btn-iconfont">
          <i class="iconfont icon-polygonred"></i>New
        </div>
        <div class="top btn-iconfont">
          <i class="iconfont icon-top"></i>Top
        </div>
      </div>
      <ul class="c-l-list">
        <li class="c-l-item">
          <div class="post">
            <a class="vote">
              <span class="iconfont icon-up"></span>
            </a>
            <span class="text">502</span>
            <a class="vote">
              <span class="iconfont icon-down"></span>
            </a>
          </div>
          <div class="l-container">
            <h4 class="con-title">Go语言真的那么流行吗？</h4>
            <div class="con-memo">
              <p>最近微信公众号、QQ群、新闻里面谈论的人有点多啊，怎么感觉大家都在说go语言，理性讨论，有一说一，是不是真的很火啊，应该要学习go语言么？</p>
            </div>
            <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>comment
              </span>
            </div>
          </div>
        </li>
        <li class="c-l-item">
          <div class="post">
            <a class>
              <span class="iconfont icon-up"></span>
            </a>
            <span class="text">107</span>
            <a class>
              <span class="iconfont icon-down"></span>
            </a>
          </div>
          <div class="l-container">
            <h4 class="con-title">What's the fastest way you've seen a streamer kill their stream?</h4>
            <div class="con-memo">
              <p>
                Obviously stuff like saying the 'n' word doesn't help, but what are some of the
                biggest mistakes you've seen streamers
                make over the years? It's usually a lot easier to see what not to do so I'd be
                interested to hear your thoughts.
              </p>
            </div>
            <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>comment
              </span>
            </div>
          </div>
        </li>
        <li class="c-l-item">
          <div class="post">
            <a class="vote">
              <span class="iconfont icon-up"></span>
            </a>
            <span class="text">50</span>
            <a class="vote">
              <span class="iconfont icon-down"></span>
            </a>
          </div>
          <div class="l-container">
            <h4 class="con-title">
              Miner's Mettle, our artillery-based strategy game, arrives on Steam April 22nd! Please
              check it out and
              consider wishlisting, it's very helpful for us!
            </h4>
            <div class="con-cover"></div>
            <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>comment
              </span>
            </div>
          </div>
        </li>
        <li class="c-l-item">
          <div class="post">
            <a class>
              <span class="iconfont icon-up"></span>
            </a>
            <span class="text">7</span>
            <a class>
              <span class="iconfont icon-down"></span>
            </a>
          </div>
          <div class="l-container">
            <h4 class="con-title">What's the fastest way you've seen a streamer kill their stream?</h4>
            <div class="con-memo">
              <p>
                Obviously stuff like saying the 'n' word doesn't help, but what are some of the
                biggest mistakes you've seen streamers
                make over the years? It's usually a lot easier to see what not to do so I'd be
                interested to hear your thoughts.
              </p>
            </div>
            <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>comment
              </span>
            </div>
          </div>
        </li>
        <li class="c-l-item">
          <div class="post">
            <a class>
              <span class="iconfont icon-up"></span>
            </a>
            <span class="text">1.7k</span>
            <a class>
              <span class="iconfont icon-down"></span>
            </a>
          </div>
          <div class="l-container">
            <h4 class="con-title">What's the fastest way you've seen a streamer kill their stream?</h4>
            <div class="con-memo">
              <p>
                Obviously stuff like saying the 'n' word doesn't help, but what are some of the
                biggest mistakes you've seen streamers
                make over the years? It's usually a lot easier to see what not to do so I'd be
                interested to hear your thoughts.
              </p>
            </div>
            <div class="user-btn">
              <span class="btn-item">
                <i class="iconfont icon-comment"></i>comment
              </span>
            </div>
          </div>
        </li>
      </ul>
    </div>
    <div class="right">
      <div class="communities">
        <h2 class="r-c-title">今日火热频道排行榜</h2>
        <ul class="r-c-content">
          <li class="r-c-item">
            <span class="index">1</span>
            <i class="icon"></i>
            b/coding
          </li>
          <li class="r-c-item">
            <span class="index">2</span>
            <i class="icon"></i>
            b/tree_hole
          </li>
          <li class="r-c-item">
            <span class="index">3</span>
            <i class="icon"></i>
            b/job
          </li>
        </ul>
        <button class="view-all">查看所有</button>
      </div>
      <div class="r-trending">
        <h2 class="r-t-title">持续热门频道</h2>
        <ul class="rank">
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/Book</span>
                <p class="info-num">7.1k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/coding</span>
                <p class="info-num">3.2k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
          <li class="r-t-cell">
            <div class="r-t-cell-info">
              <div class="avatar"></div>
              <div class="info">
                <span class="info-title">b/job</span>
                <p class="info-num">2.5k members</p>
              </div>
            </div>
            <button class="join-btn">JOIN</button>
          </li>
        </ul>
      </div>
    </div>
  </div>

</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {}
};
</script>

<style scoped src="../static/css/index.css"></style>
<style scoped src="../static/css/iconfont.css"></style>
