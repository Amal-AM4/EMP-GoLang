import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AddEmployee from '../views/AddEmployee.vue'
import UpdateEmployee from '../views/UpdateEmployee.vue'
import SingleRecord from '../views/SIngleRecord.vue'

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
        {
            path: '/single/:id',
            name: 'SingleRecord',
            component: SingleRecord
        },
    ]
})

export default router