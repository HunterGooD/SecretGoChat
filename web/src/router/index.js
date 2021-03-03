import Vue from "vue"
import VueRouter from "vue-router"
import Home from "../views/Home.vue"
Vue.use(VueRouter)

const routes = [{
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    component: null,
  },
  {
    path: "/room/:hash",
    name: "Room",
    component: null,
  }
]

// eslint-disable-next-line no-new
const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
})

export default router