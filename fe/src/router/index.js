import Vue from 'vue'
import Router from 'vue-router'

function loadView(page) {
    return () => import(`../pages/${page}.vue`)
}

Vue.use(Router)

export default new Router ({
    linkActiveClass:'Active',
    routes:[
        {
            path:'/',
            component: loadView('home')
        },
        {
             path:'/article/:id',
             component:loadView('article')
        },
        {
            path:'/manage',
            component:loadView('manage/index')
        },
        {
            path:'/notfound',
            name:'NotFound',
            component: loadView("404")
        },
        {
            path:'*',
            redirect:'/notfound'
        }
    ]
})