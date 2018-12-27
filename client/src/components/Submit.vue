<template>
  <div>
    <h1>Submit your Recipe!</h1>
    <form>
      <div class="container">
        <div class="left-side">
          <div>
            <label for="prepTime">
              Enter prep time
            </label>
            <input
              id="prepTime"
              v-model="prepTime"
            >
          </div>
          <div>
            <label for="cookTime">
              Enter cook time
            </label>
            <input
              id="cookTime"
              v-model="cookTime"
            >
          </div>         
          <div>
            <label for="feeds">
              Enter number of people recipe feeds
            </label>
            <input
              id="feeds"
              v-model="feeds"
            >
          </div>                   
        </div>
        <div class="right-side">
          <div>
            <label for="name">
              Enter recipe name
            </label>
            <input
              id="name"
              v-model="name"
            >
          </div>
          <div>
            <label for="description">
              Enter description
            </label>
            <input
              id="description"
              v-model="description"
            >
          </div>
          <div>
            <label for="ingredients">
              Enter ingredients
            </label>
            <input
              id="ingredients"
              v-model="ingredients"
            >
          </div>
          <div>
            <label for="directions">
              Enter directions
            </label>
            <input
              id="directions"
              v-model="directions"
            >
          </div>
          <button
            type="button"
            @click="recipeData"
          >
            Submit
          </button>
        </div>
      </div>
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
      prepTime: 0,
      cookTime: 0,
      feeds: 0,
    }
  },
  methods: {
    ...mapMutations(['submitRecipe']),
    recipeData() {
      const recipe = {
        name: this.name,
        description: this.description,
        ingredients: this.ingredients,
        directions: this.directions,
        prepTime: this.prepTime,
        cookTime: this.cookTime,
        feeds: this.feeds,
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

.container {
  display: grid;
  grid-template-columns: 1fr 1fr
}

.left-side {
  display: grid;
}

.right-side {
  display: grid;
}
</style>
