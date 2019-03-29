<template>
  <div id="app">
    <div class="header">
      <div class="icon">
        <RouterLink to="/">
          <a class="router">
            Icon
          </a>
        </RouterLink>
      </div>
      <div class="link">
        <RouterLink to="/">
          <a class="router">
            Home
          </a>
        </RouterLink>
      </div>
      <div class="link">
        <RouterLink to="/recipes">
          <a class="router">
            Recipes
          </a>
        </RouterLink>
      </div>
      <div />
      <div class="link">
        <RouterLink 
          to="submit" 
          append>
          <a class="router">
            Submit
          </a>
        </RouterLink>
      </div>
    </div>
    <RouterView />
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.header {
  display: grid;
  grid-template-columns: 50px 150px 150px 1fr 150px;
  background-color: lightblue;
  height: 50px;
}

.icon {
  background-color: orange;
}

.link {
  display: grid;
  justify-items: center;
  align-items: center;
}

.link:hover {
  cursor: pointer;
  background-color: gray;
}

a:-webkit-any-link {
  color: black;
  text-decoration-line: none;
}
</style>
