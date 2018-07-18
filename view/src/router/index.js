import Vue from 'vue'
import VueRouter from 'vue-router'
import todolist from '../components/Todolist.vue'
import createTodolist from '../components/createTodolist.vue'

Vue.use(VueRouter)

const routes = [
	{ path: '/', todolist }
]

export default new VueRouter({
	routes, // short for `routes: routes`
	mode: 'history'
})
