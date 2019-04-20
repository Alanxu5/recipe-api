<script>
import TextInput from '@/components/common/TextInput'

export default {
  name: 'Submit',
  components: {
    TextInput
  },
  data() {
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

<template>
  <div>
    <div class="submit-modal">
      <div class="container">
        <div class="top-container">
          <div class="image-upload" />
          <div class="short-input">
            <TextInput 
              v-model="recipeModel.name"
              label-name="Name" />
            <div class="time-input">
              <TextInput 
                v-model="recipeModel.prepTime"
                label-name="Prep Time" />
              <TextInput 
                v-model="recipeModel.cookTime"
                label-name="Cook Time" />
            </div>
            <div class="time-input">
              <TextInput 
                v-model="recipeModel.feeds"
                label-name="Servings" />
              <div class="prep-container">
                <label>
                  Preperation 
                </label>
                <select class="prep-input">
                  <option value="bake">
                    Bake
                  </option>
                  <option value="pan">
                    Pan
                  </option>
                  <option value="pressureCooker">
                    Pressure Cooker
                  </option>                
                  <option value="slowCooker">
                    Slow Cooker
                  </option>
                </select>
              </div>
            </div>                      
          </div>
        </div>
        <div class="bottom-container">
          <div class="textarea-container">
            <label>Description</label>
            <textarea
              v-model="recipeModel.description" />
          </div>
          <div class="textarea-container">
            <label>Ingredients</label>
            <textarea
              v-model="recipeModel.ingredients" />
          </div>
          <div class="textarea-container">
            <label>Directions</label>
            <textarea
              v-model="recipeModel.directions" />
          </div>         
          <div class="submit-container">
            <button
              type="button"
              @click="submitRecipe">
              Submit
            </button>    
          </div>                                     
        </div>
      </div>
    </div>
    <div class="modal-background" />
  </div>
</template>

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
  grid-template-rows: 38% 62%;
  grid-row-gap: 1rem;
}

.top-container {
  display: grid;
  grid-template-columns: 50% 50%;
}

.image-upload {
  height: 100%;
  width: 90%;
  background-color: lightgray;
  border: 1px solid #D5D1D1;
}

.textarea-container {
  display: grid;
  grid-row-gap: 5px;
  grid-template-rows: 20px 1fr;
}

.short-input {
  display: grid;
  grid-template-rows: auto auto auto;
  grid-row-gap: auto;
}

.time-input {
  display: grid;
  grid-template-columns: 145px auto;
  column-gap: 10px;
}

.prep-container {
  display: grid;
  grid-template-rows: 20px 30px;
  row-gap: 5px;
}

.prep-input {
  height: 30px;
}

.bottom-container {
  display: grid;
  grid-row-gap: 10px;
}

label {
  float: left;
}

text-input {
  float: left;
}

button {
  width: 100px;
}

textarea {
  resize: none;
}

.submit-container {
  align-self: center;
  justify-self: end;
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
