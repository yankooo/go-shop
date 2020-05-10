let Api = require("./http/api.js")
let req = require("./http/request.js")
let config = require("./env/demain.js")
let router = require("./utils/router.js")
let init = require("config.js");

let env = "Dev";
App.config = config[env];
App.version = "v0.0.0";

App({
      config: config[env],
      Api,
      router,
      get: req.fetch,
      init: init,
      userinfo:{
            code:'',
            phone: '',
      },
      post: (url, data, option) => {
            if (!option) {
                  option = {}
            }
            option.method = 'post';
            return req.fetch(url, data, option);
      },
      onLaunch: function () {

      }
})