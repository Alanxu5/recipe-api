import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    recipe: []
  },
  mutations: {
    SUBMIT_RECIPE: (state, { recipe }) => {
      state.recipe.push(recipe);
    }
  },
  actions: {
    SUBMIT_NEW_RECIPE: async function ({ commit }, recipe) {

      const response = await axios.post('http://localhost:8000/recipes', 
      {
        name: recipe.name,
        description: recipe.description,
        ingredients: recipe.ingredients,
        directions: recipe.directions,
        prepTime: recipe.prepTime,
        cookTime: recipe.cookTime,
        feeds: recipe.feeds,        
      },
      {
        headers: { 
          'Content-Type': 'application/json' 
        }
      })

      if (response.ok) {
        commit('SUBMIT_RECIPE', {recipe: response.data})
      } else {
        console.log(response)
      }
    }
  }
})
