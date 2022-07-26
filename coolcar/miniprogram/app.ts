// app.ts
App<IAppOption>({
  globalData: {},
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        wx.request({
          url: "http://localhost:10000/v1/auth/login",
          method: "post",
          data: {
            code: res.code,
          },
          success: res => {
            console.log(res)
          },
          fail: console.log,
        })
      },
    })
  },
})