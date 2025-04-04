import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AddEmployee from '../views/AddEmployee.vue'
import UpdateEmployee from '../views/UpdateEmployee.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/addEmployee',
            name: 'AddEmployee',
            component: AddEmployee
        },
        {
            path: '/edit/:id',
            name: 'UpdateEmployee',
            component: UpdateEmployee
        },
    ]
})

export default router