<script>
import { mapState, mapMutations } from 'vuex'

export default {
  name: 'Recipes',
  created() {
    this.name = this.$route.params.recipeData.name;
    this.description = this.$route.params.recipeData.description;
    this.ingredients = this.$route.params.recipeData.ingredients;
    this.directions = this.$route.params.recipeData.directions;
    this.prepTime = this.$route.params.recipeData.prepTime;
    this.cookTime = this.$route.params.recipeData.cookTime;
    this.feeds = this.$route.params.recipeData.feeds;
  }
}
</script>

<template>
  <div class="recipe-container">
    <h1>View Recipes</h1>
    <div class="container">
      <div class="summary">
        <div class="info">
          <div class="title">
            {{ name }}
          </div>
          <div class="time">
            <div>Prep: {{ prepTime }}</div>
            <div>Cook: {{ cookTime }}</div>
          </div>
          <div>Number of meals: {{ feeds }}</div>
          <div>Description: {{ description }}</div>
        </div>
        <div class="image-container">
          <img 
            class="image"
            src="https://cdn.fstoppers.com/styles/full/s3/media/2015/03/apple-iphone-6-world-gallery-fstoppers-alastair_b.png">
        </div>
      </div>
      <div class="instructions">
        <div class="ingredients">
          <div>{{ ingredients }}</div>
        </div>
        <div class="directions">
          <div>{{ directions }}</div>
        </div>
      </div>
    </div>
    <RouterView />
  </div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .recipe-container {
    margin: 30px 30px 30px 30px;
  }

  .container {
    display: grid;
    grid-template-rows: auto auto
  }

  .summary {
    display: grid;
    grid-template-columns: auto 500px;
  }

  .info {

  }

  .title {
    font-size: 16px;
  }

  .time {
    display: grid;
    grid-template-columns: 40px 40px;
    /* justify-content: center; */
  }

  .image-container {

  }

  .image {
    width: 300px;
    height: 300px;
  }

  .instructions {
    display: grid;
    grid-template-rows: auto auto;
  }

  .ingredients {

  }

  .directions {

  }
</style>
