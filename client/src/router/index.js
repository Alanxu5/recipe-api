import Vue from 'vue'
import Router from 'vue-router'

// @ is an alias for src directory
import Recipes from '@/components/Recipes.vue'
import Submit from '@/components/Submit.vue'
import Home from '@/components/Home.vue'

// TODO: necessary?
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      children: [
        {
          path: '/submit',
          name: 'submit',
          component: Submit
        }
      ]
    },
    {
      path: '/recipes',
      name: 'recipes',
      component: Recipes,
    }
  ]
})