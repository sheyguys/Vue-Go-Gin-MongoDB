import Register from './components/Register.vue'
import Result from './components/Result.vue'
import Home from './components/Home.vue'


const routes = [
    { path: '/', component: Home },
    { path: '/register', component: Register },
    { path: '/result', component: Result },
];

export default routes;