import Vue from "vue"
import VueRouter from "vue-router"
import Home from "../views/Home.vue"
import Room from "../views/Room.vue"
import About from "../views/About.vue"
Vue.use(VueRouter)

const routes = [{
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: "/room/:hash",
    name: "Room",
    component: Room,
  }
]

// eslint-disable-next-line no-new
const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
})

export default router