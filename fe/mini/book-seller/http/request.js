/**
 * 网络请求的公共方法
 * 1. 基本请求
 * 2. 为了后续获取数据方方便，promise处理
 * 3. 对获取数据的状态处理: loadding toast
 * 4. 对请求头对处理，添加一些机型，大小，系统，屏幕等信息到请求头
 */
let store = require("../utils/store.js");
let system = store.getSystemInfo();
const clientInfo = {
  "clientType": "mp",
  "appnm": "school-eco",
  //
  "model": system.model,
  "os": system.system,
  "screen": system.screenWidth + "*" + system.screenHeight,
  "version": App.version
}

module.exports = {
  fetch: (url, data = {}, option = {}) => {
    let {
        method = 'get'
    } = option;
    return new Promise((resolve, reject) => {
      let env = App.config.baseApi;
      wx.request({
        url: env + url,
        data: data,
        method,
        header: {
          "clientInfo": JSON.stringify(clientInfo)
        },
        success: function (result) {
          let res = result.data;
          resolve(res.data);
        },
        fail: function (e) {
          reject(e);
        }
      })
    })
  }
}