<template>
  <div class="input-container">
    <label>
      {{ labelName }}
    </label>
    <input
      v-model="inputVal">
  </div>
</template>

<script>
export default {
  name: 'TextInput',
  props: {
    labelName: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      inputVal: this.value
    }
  },
  watch: {
    inputVal(val) {
      this.$emit('input', val);
    }
  }  
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .input-container {
    display: grid;
    grid-row-gap: 5px;
    grid-template-rows: 20px 30px;
  }
</style>
