import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    recipe: {}
  },
  mutations: {
    submitRecipe(state, recipe) {
      state.recipe = recipe;
    }
  },
  actions: {

  }
})
