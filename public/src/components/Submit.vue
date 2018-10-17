<template>
  <div>
    <h1>Submit your Recipe!</h1>
    <form>
      <div>
        <label for="name">Enter recipe name</label>
        <input id="name" v-model="name">
      </div>
      <div>
        <label for="description">Enter description</label>
        <input id="description" v-model="description">
      </div>
      <div>
        <label for="ingredients">Enter ingredients</label>
        <input id="ingredients" v-model="ingredients">
      </div>
      <div>
        <label for="directions">Enter directions</label>
        <input id="directions" v-model="directions">
      </div>
      <button type="button" @click="recipeData">Submit</button>
    </form>
  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
export default {
  name: 'Submit',
  data: function () {
    return {
      name: "",
      description: "",
      ingredients: "",
      directions: "",
    }
  },
  methods: {
    ...mapMutations(['submitRecipe']),
    recipeData(e) {
      const recipe = {
        name: this.name,
        description: this.description,
        ingredients: this.ingredients,
        directions: this.directions,
      };

      this.submitRecipe(recipe);
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
</style>
