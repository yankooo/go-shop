/**
 * 处理路由的跳转
 */

// 创建映射关系
const routerPath = {
  "index": "/pages/index/index",
  "publisgh": "/pages/publish/publish",
  "account": "/pages/account/account"
}

module.exports = {
  // this.$router.push("/index", {path:"index",query:{}})
  redirectTo({}) {
    routerPath[path]
  },
  push(path, option = {}) {
    if (typeof path === 'string') {
      option.path = path;
    } else {
      option = path;
    }

    // 获取url
    let url = routerPath[option.path]
    // openType跳转类型
    let {
      query = {}, openType
    } = option;
    let params = this.parse(query);

    if (params) {
      url += '?' + params
    }
    this.to(openType, url)
  },

  to(openType, url) {
    let obj = {
      url
    };
    if (openType === 'redirect') {
      wx.redirectTo(obj);
    } else if (openType === 'reLaunch') {
      wx.reLaunch(obj);
    } else if (openType === 'back') {
      wx.navigateBack({
        delta: 1
      })
    } else if (openType === 'switch') {
      wx.switchTab(obj)
    } else {
      wx.navigateTo(obj)
    }
  },

  parse(data) {
    let arr = [];
    for (let key in data) {
      arr.push(key + '=' + data[key]);
    }
    return arr.join("&");
  }
}