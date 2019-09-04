import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import  firebase from 'firebase';
import VueRouter from 'vue-router'

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes
});

const config = {
  apiKey: "AIzaSyC_u27mI5neEM6BVCM1yqrAN-WHTk8-QKE",
  authDomain: "course-online-b26e1.firebaseapp.com",
  databaseURL: "https://course-online-b26e1.firebaseio.com",
  projectId: "course-online-b26e1",
  storageBucket: "course-online-b26e1.appspot.com",
  messagingSenderId: "179670983174",
  appId: "1:179670983174:web:64c4ea52d04ec19d"
};

import routes from './routes';

Vue.config.productionTip = false

firebase.initializeApp(config);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
