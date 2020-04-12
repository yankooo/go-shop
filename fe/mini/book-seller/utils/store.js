/**
 * 存储数据
 * Storage信息存储
 */

 module.exports = {
     /**
      * 设置值
      * @param {键} key 
      * @param {值} value 
      * @param {信息模块} module_name 
      */
     setItem(key, value, module_name) {
         if (module_name) {
             let module_name_info = this.getItem(module_name);
             module_name_info[key] = value
             wx.setStorageSync(key, module_name_info)
         } else {
             wx.setStorageSync(key, data)
         }
     },

     /**
      * 获取值
      * @param {键} key 
      * @param {信息模块} module_name 
      */
     getItem(key, module_name) {
         if (module_name) {
             let val = this.getItem(module_name);
             if (val) return val[key]
             return ""
         } else {
             return wx.getStorageSync(key);
         }
     },

     /**
      * 
      * @param {键} key 
      */
     clear(key) {
         key ? wx.removeStorageSync(key) : wx.clearStorageSync();
     },

     getSystemInfo() {
         return wx.getSystemInfoSync();
     }

 }