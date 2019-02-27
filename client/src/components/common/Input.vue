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
  name: 'Input',
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
</style>
