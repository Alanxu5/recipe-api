import Vue from 'vue'
import Vuex from 'vuex'
import router from '@/router'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    recipe: []
  },
  mutations: {
    SUBMIT_RECIPE: (state, { recipe }) => {
      state.recipe.push(recipe);
      router.push({name: 'recipes', params: { recipeData: recipe }});
    },
    GET_RECIPE: (state, { recipe }) => {
      router.push({name: 'recipe', params: { recipeDate: recipe }});
    }
  },
  actions: {
    SUBMIT_NEW_RECIPE: async function ({ commit }, recipe) {
      try {
        // TODO: make it dynamic
        const response = await fetch('http://localhost:8000/recipes', {
          method: "POST",
          headers: { 
            'Content-Type': 'application/json' 
          },
          body: JSON.stringify(recipe)
        });

        const responseJson = await response.json();

        if (response.ok) {
          commit('SUBMIT_RECIPE', {recipe: recipe})
        } else {
          console.log(response)
        }
      } catch(err) {
        console.log(err);
      }
    }, 
    GET_NEWEST_RECIPE: async function ({ commit }, recipe) {
      try {
        // TODO: make it dynamic
        const response = await fetch('http://localhost:8000/recipes', {
          method: "GET",
          headers: { 
            'Content-Type': 'application/json' 
          },
        });

        const responseJson = await response.json();

        if (response.ok) {
          commit('GET_RECIPE', {recipe: recipe})
        } else {
          console.log(response)
        }
      } catch(err) {
        console.log(err);
      }      
    }
  }
})
