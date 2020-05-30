import Vue from 'vue'
import App from './App.vue'
import router from './router'

import store from './store'
import ElementUI from 'element-ui'
import Viewer from 'v-viewer'
import 'element-ui/lib/theme-chalk/index.css'
import VueSplit from 'vue-split-panel'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueSplit)


new Vue({
  Viewer,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
Viewer.setDefaults({
  Options: { "inline": true, "button": true, "navbar": true, "title": true, "toolbar": true, "tooltip": true, "movable": true, "zoomable": true, "rotatable": true, "scalable": true, "transition": true, "fullscreen": true, "keyboard": true, "url": "data-source" }
});
