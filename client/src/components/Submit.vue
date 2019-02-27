<template>
  <div>
    <div class="submit-modal">
      <div class="container">
        <div class="top-container">
          <div class="image-upload" />
          <div class="short-input">
            <Input 
              v-model="recipeModel.name"
              label-name="Name" />
            <Input 
              v-model="recipeModel.prepTime"
              label-name="Prep Time" />
            <Input 
              v-model="recipeModel.cookTime"
              label-name="Cook Time" />
            <Input 
              v-model="recipeModel.feeds"
              label-name="Servings" />                             
          </div>
        </div>
        <div class="bottom-container">
          <Input 
            v-model="recipeModel.description"
            label-name="Description" />
          <Input 
            v-model="recipeModel.ingredients"
            label-name="Ingredients" />
          <Input 
            v-model="recipeModel.directions"
            label-name="Directions" />                  
          <button
            type="button"
            @click="submitRecipe">
            Submit
          </button>                          
        </div>
      </div>
    </div>
    <div class="modal-background" />
  </div>
</template>

<script>
import Input from '@/components/common/Input'

export default {
  name: 'Submit',
  components: {
    Input
  },
  data: function () {
    return {
      recipeModel: {
        name: "",
        description: "",
        ingredients: "",
        directions: "",
        prepTime: 0,
        cookTime: 0,
        feeds: 0,
      }
    }
  },
  methods: {
    submitRecipe() {
      const recipe = {
        name: this.recipeModel.name,
        description: this.recipeModel.description,
        ingredients: this.recipeModel.ingredients,
        directions: this.recipeModel.directions,
        prepTime: this.recipeModel.prepTime,
        cookTime: this.recipeModel.cookTime,
        feeds: this.recipeModel.feeds,
      };

      this.$store.dispatch('SUBMIT_NEW_RECIPE', recipe)
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.submit-modal {
  background: white;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) !important;
  z-index: 15;
  width: 600px;
  height: 600px;
  border: 1px solid lightgrey;
  padding: 24px;
}

.container {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-rows: 45% 55%;
}

.top-container {
  display: grid;
  grid-template-columns: 60% 40%;
}

.image-upload {
  height: 270px;
  width: 300px;
  justify-self: left;
  background-color: lightgray;
  border: 1px solid #D5D1D1;
}

.short-input {
  display: grid;
  grid-row-gap: auto;
}

.bottom-container {
  display: grid;
}

label {
  float: left;
}

input {
  float: left;
}

button {
  width: 60px;
}

.modal-background {
  background: rgba(0, 0, 0, .30);
  height: 100%;
  width: 100%;
  z-index: 10;
  position: fixed;
  left: 0;
  top: 0;
}
</style>
