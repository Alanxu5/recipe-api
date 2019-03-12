import Vue from 'vue'
import Router from 'vue-router'

// @ is an alias for src directory
import Recipes from '@/components/Recipes.vue'
import Submit from '@/components/Submit.vue'
import Home from '@/components/Home.vue'

// tells vue to use vue router
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/recipes',
      name: 'recipes',
      component: Recipes,
      children: [
        {
          path: '/recipes/submit',
          name: 'submit',
          component: Submit
        }
      ]
    },
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
    }
  ]
})