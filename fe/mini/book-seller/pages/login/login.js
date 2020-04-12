const app = getApp();
let api = require("../../http/api.js")
let router = require("../../utils/router.js")
Page({
      /**
       * 页面的初始数据
       */
      data: {
            ids: -1,
            phone: '',
            wxnum: '',
            qqnum: '',
            email: '',
            campus: JSON.parse(app.init.data).campus,
      },
      choose(e) {
            let that = this;
            that.setData({
                  ids: e.detail.value
            })
            //下面这种办法无法修改页面数据
            /* this.data.ids = e.detail.value;*/
      },
      //获取用户手机号
      getPhoneNumber: function (e) {
            let that = this;
            //判断用户是否授权确认
            if (!e.detail.errMsg || e.detail.errMsg != "getPhoneNumber:ok") {
                  wx.showToast({
                        title: '获取手机号失败',
                        icon: 'none'
                  })
                  return;
            }
            wx.showLoading({
                  title: '获取手机号中...',
            })
            wx.login({
                  success(re) {
                        // 发送请求获取信息， getSession
                        console.log(re.code)
                        wx.hideLoading();
                        that.setData({
                              phone: "18512605170"
                        })
                  },
                  fail: err => {
                        console.error(err);
                        wx.hideLoading()
                        wx.showToast({
                              title: '获取失败,请重新获取',
                              icon: 'none',
                              duration: 2000
                        })
                  }
            })
      },
      wxInput(e) {
            this.data.wxnum = e.detail.value;
      },
      qqInput(e) {
            this.data.qqnum = e.detail.value;
      },
      emInput(e) {
            this.data.email = e.detail.value;
      },
      getUserInfo(e) {
            let that = this;
            console.log(e);
            let test = e.detail.errMsg.indexOf("ok");
            if (test == '-1') {
                  wx.showToast({
                        title: '请授权后方可使用',
                        icon: 'none',
                        duration: 2000
                  });
            } else {
                  that.setData({
                        userInfo: e.detail.userInfo
                  })
                  that.check();
            }
      },
      //校检
      check() {
            let that = this;
            //校检手机
            let phone = that.data.phone;
            if (phone == '') {
                  wx.showToast({
                        title: '请先获取您的电话',
                        icon: 'none',
                        duration: 2000
                  });
                  return false
            }
            //校检校区
            let ids = that.data.ids;
            let campus = that.data.campus;
            if (ids == -1) {
                  wx.showToast({
                        title: '请先获取您的校区',
                        icon: 'none',
                        duration: 2000
                  });
            }
            //校检邮箱
            let email = that.data.email;
            if (!(/^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/.test(email))) {
                  wx.showToast({
                        title: '请输入正确的邮箱',
                        icon: 'none',
                        duration: 2000
                  });
                  return false;
            }
            //校检QQ号
            let qqnum = that.data.qqnum;
            if (qqnum !== '') {
                  if (!(/^\s*[.0-9]{5,11}\s*$/.test(qqnum))) {
                        wx.showToast({
                              title: '请输入正确QQ号',
                              icon: 'none',
                              duration: 2000
                        });
                        return false;
                  }
            }
            //校检微信号
            let wxnum = that.data.wxnum;
            if (wxnum !== '') {
                  if (!(/^[a-zA-Z]([-_a-zA-Z0-9]{5,19})+$/.test(wxnum))) {
                        wx.showToast({
                              title: '请输入正确微信号',
                              icon: 'none',
                              duration: 2000
                        });
                        return false;
                  }
            }
            wx.showLoading({
                  title: '正在提交',
                  mask: true
            })
            // 发送post注册请求
            let data = {
                  nick_name: this.data.userInfo.nickName,
                  email: this.data.email,
                  open_id: "fgdsfgafffffgserg",
                  mobile: this.data.phone,
                  avatar: this.data.userInfo.avatarUrl,
                  gender: "男",
                  school: "大连理工大学",
                  major: "电信"
            }
            app.post(api.register, JSON.stringify(data)).then(res => {
                  wx.hideLoading()
                  let msg = ""
                  if (res.res_code == -1) {
                        msg = "该用户已经注册，请直接登录"
                  } else if (res.res == 100001) {

                  } else {
                        msg = "注册成功"
                        // TODO 用户信息设置到变量中
                  }
                  wx.showModal({
                        confirmColor: "FFFFFF",
                        title: msg,
                        complete: function () {
                              wx.switchTab({
                                    url: '/pages/account/account',
                              })
                        }
                  })
                  console.log("res:", res)
            }).catch(err => {
                  wx.hideLoading()
                  console.log("err: ", err)
            })
      },
      onLoad: function () {}
})